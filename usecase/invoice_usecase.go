package usecase

import (
	"context"
	"time"

	"github.com/kohge2/upsdct-server/config"
	"github.com/kohge2/upsdct-server/domain/models"
	"github.com/kohge2/upsdct-server/domain/repository"
	"github.com/kohge2/upsdct-server/utils"
)

type InvoiceUseCase interface {
	CreateInvoice(paymentAmount int, partnerCompanyID, userID string, invoiceDueDate, now time.Time) error
	GetInvoices(userID string, startDate, endDate *time.Time) (models.InvoiceList, models.PartnerCompanyEmbedList, error)
}

type invoiceUseCase struct {
	invoiceRepository        repository.InvoiceRepository
	userRepository           repository.UserRepository
	companyRepository        repository.CompanyRepository
	partnerCompanyRepository repository.PartnerCompanyRepository
	transaction              Transaction
}

func NewInvoiceUseCase(invoiceRepositry repository.InvoiceRepository, userRepository repository.UserRepository, companyRepository repository.CompanyRepository, partnerCompanyRepository repository.PartnerCompanyRepository, transaction Transaction) InvoiceUseCase {
	return &invoiceUseCase{
		invoiceRepository:        invoiceRepositry,
		userRepository:           userRepository,
		companyRepository:        companyRepository,
		partnerCompanyRepository: partnerCompanyRepository,
		transaction:              transaction,
	}
}

func (u *invoiceUseCase) CreateInvoice(paymentAmount int, partnerCompanyID, userID string, invoiceDueDate, now time.Time) error {
	if err := u.transaction.RunTxn(context.Background(), func(ctx context.Context) error {
		// 存在しない user, company, partner_company の場合は404を返す
		user, err := u.userRepository.FindByUserID(ctx, userID)
		if err != nil {
			return err
		}
		if _, err := u.companyRepository.FindByCompanyID(ctx, user.CompanyID); err != nil {
			return err
		}
		if _, err := u.partnerCompanyRepository.FindByPartnerCompanyID(ctx, partnerCompanyID); err != nil {
			return err
		}

		invoice := &models.Invoice{
			ID:               utils.GenerateULID(),
			CompanyID:        user.CompanyID,
			PartnerCompanyID: partnerCompanyID,
			PaidAmount:       paymentAmount,
			PublishedDate:    now,
			PaidDueDate:      invoiceDueDate,
			CommissionRate:   utils.Float64Ptr(config.DefaultCommissionRate),
			TaxRate:          utils.Float64Ptr(config.DefaultTaxRate),
			CreatedBy:        userID,
			InvoiceStatus:    models.InvoiceStatusOpen,
		}

		commission, err := invoice.CalcCommission()
		if err != nil {
			return err
		}
		invoice.SetCommission(commission)

		billedAmount, tax, err := invoice.CalcBilledAmount(config.DefaultTaxRate)
		if err != nil {
			return err
		}
		invoice.SetBilledAmount(billedAmount)
		invoice.SetTax(tax)

		if err := u.invoiceRepository.CreateInvoice(ctx, invoice); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (u *invoiceUseCase) GetInvoices(userID string, startDate, endDate *time.Time) (models.InvoiceList, models.PartnerCompanyEmbedList, error) {
	ctx := context.Background()
	// 存在しない user, company の場合は404を返す
	user, err := u.userRepository.FindByUserID(ctx, userID)
	if err != nil {
		return nil, nil, err
	}

	company, err := u.companyRepository.FindByCompanyID(ctx, user.CompanyID)
	if err != nil {
		return nil, nil, err
	}

	invoices, err := u.invoiceRepository.FindInvoicesByCompanyIDAndPaidDueDateRange(ctx, company.ID, startDate, endDate)
	if err != nil {
		return nil, nil, err
	}

	partnerCompanies, err := u.partnerCompanyRepository.FindPartnerCompanyEmbedListByPartnerCompanyIDs(ctx, invoices.UniquePartnerCompanyIDs())
	if err != nil {
		return nil, nil, err
	}

	// 暗号化して保存していた銀行口座情報を復号
	for _, partnerCompany := range partnerCompanies {
		if partnerCompany.PartnerCompanyBankAccount != nil {
			if _, err := partnerCompany.PartnerCompanyBankAccount.SetDecryptedAccountNumber(config.Env.EncryptKey); err != nil {
				return nil, nil, err
			}
			if _, err := partnerCompany.PartnerCompanyBankAccount.SetDecryptedAccountHolderName(config.Env.EncryptKey); err != nil {
				return nil, nil, err
			}
		}
	}

	return invoices, partnerCompanies, nil
}
