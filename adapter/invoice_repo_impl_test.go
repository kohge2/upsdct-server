package adapter

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kohge2/upsdct-server/database"
	"github.com/kohge2/upsdct-server/domain/models"
	"github.com/kohge2/upsdct-server/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestCreateInvoice(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer sqlDB.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	testDB := database.NewDB(gormDB)
	repo := NewInvoiceRepository(testDB)

	assert.NoError(t, err)

	t.Run("正常系", func(t *testing.T) {
		invoice := models.Invoice{
			ID:               "inv12345",
			CompanyID:        "12345",
			PartnerCompanyID: "54321",
			PaidAmount:       1000,
			BilledAmount:     utils.IntPtr(1040),
			CommissionRate:   utils.Float64Ptr(0.04),
			Commission:       utils.IntPtr(40),
			Tax:              utils.IntPtr(40),
			TaxRate:          utils.Float64Ptr(0.1),
			PaidDueDate:      time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			PublishedDate:    time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			InvoiceStatus:    models.InvoiceStatusOpen,
			CreatedBy:        "user12345",
		}

		query := regexp.QuoteMeta("INSERT INTO `invoices` (`created_at`,`updated_at`,`deleted_at`,`id`,`company_id`,`partner_company_id`,`published_date`,`commission`,`commission_rate`,`tax_rate`,`tax`,`paid_amount`,`billed_amount`,`paid_due_date`,`invoice_status`,`created_by`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
		mock.ExpectBegin()
		mock.ExpectExec(query).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, invoice.ID, invoice.CompanyID, invoice.PartnerCompanyID, invoice.PublishedDate,
				invoice.Commission, invoice.CommissionRate, invoice.TaxRate, invoice.Tax,
				invoice.PaidAmount, invoice.BilledAmount, invoice.PaidDueDate, invoice.InvoiceStatus, invoice.CreatedBy).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		ctx := context.Background()
		err := repo.CreateInvoice(ctx, &invoice)

		assert.NoError(t, err)
	})

}

func TestFindInvoicesByCompanyIDAndPaidDueDateRange(t *testing.T) {

	sqlDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer sqlDB.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	testDB := database.NewDB(gormDB)
	repo := NewInvoiceRepository(testDB)

	assert.NoError(t, err)

	t.Run("正常系_期間指定あり(startDate,endDate)", func(t *testing.T) {
		companyID := "cp001"
		startDate := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
		endDate := time.Date(2021, 1, 31, 0, 0, 0, 0, time.UTC)

		query := regexp.QuoteMeta("SELECT * FROM `invoices` WHERE paid_due_date >= ? AND paid_due_date <= ? AND company_id = ? AND `invoices`.`deleted_at` IS NULL")
		mock.ExpectQuery(query).
			WithArgs(startDate, endDate, companyID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "company_id", "partner_company_id", "published_date", "commission", "commission_rate", "tax_rate", "tax", "paid_amount", "billed_amount", "paid_due_date", "invoice_status", "created_by"}))

		ctx := context.Background()
		invoices, err := repo.FindInvoicesByCompanyIDAndPaidDueDateRange(ctx, companyID, &startDate, &endDate)

		assert.NoError(t, err)
		assert.NotNil(t, invoices)
	})

	t.Run("正常系_期間指定あり(startDateのみ)", func(t *testing.T) {
		companyID := "cp001"
		startDate := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)

		query := regexp.QuoteMeta("SELECT * FROM `invoices` WHERE paid_due_date >= ? AND company_id = ? AND `invoices`.`deleted_at` IS NULL")
		mock.ExpectQuery(query).
			WithArgs(startDate, companyID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "company_id", "partner_company_id", "published_date", "commission", "commission_rate", "tax_rate", "tax", "paid_amount", "billed_amount", "paid_due_date", "invoice_status", "created_by"}))

		ctx := context.Background()
		invoices, err := repo.FindInvoicesByCompanyIDAndPaidDueDateRange(ctx, companyID, &startDate, nil)

		assert.NoError(t, err)
		assert.NotNil(t, invoices)
	})

	t.Run("正常系_期間指定あり(endDateのみ)", func(t *testing.T) {
		companyID := "cp001"
		endDate := time.Date(2021, 1, 31, 0, 0, 0, 0, time.UTC)

		query := regexp.QuoteMeta("SELECT * FROM `invoices` WHERE paid_due_date <= ? AND company_id = ? AND `invoices`.`deleted_at` IS NULL")
		mock.ExpectQuery(query).
			WithArgs(endDate, companyID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "company_id", "partner_company_id", "published_date", "commission", "commission_rate", "tax_rate", "tax", "paid_amount", "billed_amount", "paid_due_date", "invoice_status", "created_by"}))

		ctx := context.Background()
		invoices, err := repo.FindInvoicesByCompanyIDAndPaidDueDateRange(ctx, companyID, nil, &endDate)

		assert.NoError(t, err)
		assert.NotNil(t, invoices)
	})

	t.Run("正常系_期間指定なし", func(t *testing.T) {
		companyID := "cp001"

		query := regexp.QuoteMeta("SELECT * FROM `invoices` WHERE company_id = ? AND `invoices`.`deleted_at` IS NULL")
		mock.ExpectQuery(query).
			WithArgs(companyID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "company_id", "partner_company_id", "published_date", "commission", "commission_rate", "tax_rate", "tax", "paid_amount", "billed_amount", "paid_due_date", "invoice_status", "created_by"}))

		ctx := context.Background()
		invoices, err := repo.FindInvoicesByCompanyIDAndPaidDueDateRange(ctx, companyID, nil, nil)

		assert.NoError(t, err)
		assert.NotNil(t, invoices)
	})

}
