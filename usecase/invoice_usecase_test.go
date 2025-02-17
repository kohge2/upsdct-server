package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/kohge2/upsdct-server/domain/models"
	mock_repository "github.com/kohge2/upsdct-server/testmock/domain/repository"
	mock_usecase "github.com/kohge2/upsdct-server/testmock/usecase"
	"github.com/kohge2/upsdct-server/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateInvoice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockInvoiceRepository := mock_repository.NewMockInvoiceRepository(ctrl)
	mockUserRepository := mock_repository.NewMockUserRepository(ctrl)
	mockCompanyRepository := mock_repository.NewMockCompanyRepository(ctrl)
	mockPartnerCompanyRepository := mock_repository.NewMockPartnerCompanyRepository(ctrl)
	mockTransaction := mock_usecase.NewMockTransaction(ctrl)

	usecase := NewInvoiceUseCase(mockInvoiceRepository, mockUserRepository, mockCompanyRepository, mockPartnerCompanyRepository, mockTransaction)
	t.Run("正常系", func(t *testing.T) {
		userID := "u0001"
		user := &models.User{
			ID:        userID,
			CompanyID: "cp0001",
		}
		partnerCompanyID := "pc0001"
		ctx := context.Background()
		mockTransaction.EXPECT().RunTxn(ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, f func(ctx context.Context) error) error {
			return f(ctx)
		})
		mockUserRepository.EXPECT().FindByUserID(ctx, userID).Return(user, nil)
		mockCompanyRepository.EXPECT().FindByCompanyID(ctx, user.CompanyID).Return(nil, nil)
		mockPartnerCompanyRepository.EXPECT().FindByPartnerCompanyID(ctx, partnerCompanyID).Return(nil, nil)
		mockInvoiceRepository.EXPECT().CreateInvoice(ctx, gomock.Any()).Return(nil)

		date := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
		err := usecase.CreateInvoice(1000, partnerCompanyID, userID, date, date)

		assert.NoError(t, err)
	})
}

func TestGetInvoices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockInvoiceRepository := mock_repository.NewMockInvoiceRepository(ctrl)
	mockUserRepository := mock_repository.NewMockUserRepository(ctrl)
	mockCompanyRepository := mock_repository.NewMockCompanyRepository(ctrl)
	mockPartnerCompanyRepository := mock_repository.NewMockPartnerCompanyRepository(ctrl)
	mockTransaction := mock_usecase.NewMockTransaction(ctrl)

	usecase := NewInvoiceUseCase(mockInvoiceRepository, mockUserRepository, mockCompanyRepository, mockPartnerCompanyRepository, mockTransaction)

	t.Run("正常系", func(t *testing.T) {
		// ここにテストケースを書く
		userID := "u0001"
		stareDate := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
		endDate := time.Date(2021, 1, 31, 0, 0, 0, 0, time.UTC)

		user := &models.User{
			ID:        userID,
			CompanyID: "cp0001",
		}
		company := &models.Company{}

		invoices := models.InvoiceList{
			{
				ID:               "inv0001",
				CompanyID:        "cp0001",
				PartnerCompanyID: "pc0001",
				PaidAmount:       1000,
				BilledAmount:     utils.IntPtr(1040),
				CommissionRate:   utils.Float64Ptr(0.04),
				Commission:       utils.IntPtr(40),
				Tax:              utils.IntPtr(40),
				TaxRate:          utils.Float64Ptr(0.1),
				PaidDueDate:      time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				PublishedDate:    time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				InvoiceStatus:    models.InvoiceStatusOpen,
				CreatedBy:        "u0001",
			},
		}

		partnerCompany := models.PartnerCompanyEmbed{
			PartnerCompany: models.PartnerCompany{
				ID: "pc0001",
			},
		}
		partnerCompanies := models.PartnerCompanyEmbedList{&partnerCompany}

		ctx := context.Background()

		mockUserRepository.EXPECT().FindByUserID(ctx, userID).Return(user, nil)
		mockCompanyRepository.EXPECT().FindByCompanyID(ctx, user.CompanyID).Return(company, nil)
		mockInvoiceRepository.EXPECT().FindInvoicesByCompanyIDAndPaidDueDateRange(ctx, company.ID, &stareDate, &endDate).Return(invoices, nil)
		mockPartnerCompanyRepository.EXPECT().FindPartnerCompanyEmbedListByPartnerCompanyIDs(ctx, invoices.UniquePartnerCompanyIDs()).Return(partnerCompanies, nil)

		actualInvoices, err := usecase.GetInvoices(userID, &stareDate, &endDate)
		assert.NoError(t, err)
		assert.Equal(t, models.NewInvoiceEmbedList(invoices, partnerCompanies), actualInvoices)
	})
}
