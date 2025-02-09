package adapter

import (
	"context"

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
