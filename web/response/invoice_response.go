package response

import (
	"time"

	"github.com/kohge2/upsdct-server/domain/models"
	"github.com/kohge2/upsdct-server/utils"
)

type GetInvoicesResponse struct {
	Invoices []getInvoiceResponseItem `json:"invoices"`
}

type getInvoiceResponseItem struct {
	ID                        string                            `json:"id"`
	PublishedDate             string                            `json:"publishedDate"`
	PaidDueDate               string                            `json:"paidDueDate"`
	InvoiceStatus             string                            `json:"invoiceStatus"`
	PaidAmount                int                               `json:"paidAmount"`
	BilledAmount              int                               `json:"billedAmount"`
	Commission                int                               `json:"commission"`
	Tax                       int                               `json:"tax"`
	PartnerCompanyID          string                            `json:"partnerCompanyID"`
	PartnerCompanyName        string                            `json:"partnerCompanyName"`
	PartnerCompanyBankAccount partnerCompanyBankAccountResponse `json:"partnerCompanyBankAccount"`
}

type partnerCompanyBankAccountResponse struct {
	BankName          string `json:"bankName"`
	BranchName        string `json:"branchName"`
	AccountType       string `json:"accountType"`
	AccountNumber     string `json:"accountNumber"`
	AccountHolderName string `json:"accountHolderName"`
}

func NewGetInvoicesResponse(invoices models.InvoiceList, partnerCompanies models.PartnerCompanyEmbedList) GetInvoicesResponse {
	invoiceResponseItems := make([]getInvoiceResponseItem, 0, len(invoices))

	partnerCompaniesMap := make(map[string]models.PartnerCompanyEmbed)
	for _, partnerCompany := range partnerCompanies {
		partnerCompaniesMap[partnerCompany.ID] = *partnerCompany
	}

	for _, invoice := range invoices {
		partnerCompany := partnerCompaniesMap[invoice.PartnerCompanyID]
		invoiceResponseItem := getInvoiceResponseItem{
			ID:                        invoice.ID,
			PublishedDate:             invoice.PublishedDate.Format("2006-01-02"),
			PaidDueDate:               utils.TimeJST(invoice.PaidDueDate).Format(time.RFC3339),
			InvoiceStatus:             invoice.InvoiceStatus.Label(),
			PaidAmount:                invoice.PaidAmount,
			BilledAmount:              *invoice.BilledAmount,
			Commission:                *invoice.Commission,
			Tax:                       *invoice.Tax,
			PartnerCompanyID:          partnerCompany.ID,
			PartnerCompanyName:        partnerCompany.Name,
			PartnerCompanyBankAccount: partnerCompanyBankAccountResponse{},
		}

		if partnerCompany.PartnerCompanyBankAccount != nil {

			invoiceResponseItem.PartnerCompanyBankAccount = partnerCompanyBankAccountResponse{
				BankName:          partnerCompany.PartnerCompanyBankAccount.BankName,
				BranchName:        partnerCompany.PartnerCompanyBankAccount.BranchName,
				AccountType:       partnerCompany.PartnerCompanyBankAccount.AccountType.Label(),
				AccountNumber:     partnerCompany.PartnerCompanyBankAccount.AccountNumber,
				AccountHolderName: partnerCompany.PartnerCompanyBankAccount.AccountHolderName,
			}
		}

		invoiceResponseItems = append(invoiceResponseItems, invoiceResponseItem)
	}
	return GetInvoicesResponse{
		Invoices: invoiceResponseItems,
	}
}

type PostResponse struct {
	OK int `json:"ok"`
}

func NewPostResponse() PostResponse {
	return PostResponse{
		OK: 1,
	}
}
