package adapter

import (
	"context"
	"time"

	"github.com/kohge2/upsdct-server/database"
	"github.com/kohge2/upsdct-server/domain/models"
)

type InvoiceRepository struct {
	db *database.DB
}

func NewInvoiceRepository(db *database.DB) *InvoiceRepository {
	return &InvoiceRepository{db: db}
}

func (a *InvoiceRepository) CreateInvoice(ctx context.Context, invoice *models.Invoice) error {
	if err := a.db.GetNewTxnOrContext(ctx).Create(invoice).Error; err != nil {
		return err
	}
	return nil
}

func (a *InvoiceRepository) FindInvoicesByCompanyIDAndPaidDueDateRange(ctx context.Context, companyID string, startDate, endDate *time.Time) (models.InvoiceList, error) {
	invoices := []*models.Invoice{}
	tx := a.db.GetNewTxnOrContext(ctx)
	if startDate != nil {
		tx = tx.Where("paid_due_date >= ?", startDate)
	}
	if endDate != nil {
		tx = tx.Where("paid_due_date <= ?", endDate)
	}

	if err := tx.
		Where("company_id = ?", companyID).
		Find(&invoices).Error; err != nil {
		return nil, err
	}
	return invoices, nil
}
