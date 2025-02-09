package repository

import (
	"context"

	"github.com/kohge2/upsdct-server/domain/models"
)

type InvoiceRepository interface {
	CreateInvoice(ctx context.Context, invoice *models.Invoice) error
}
