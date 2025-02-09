package repository

import (
	"context"
	"time"

	"github.com/kohge2/upsdct-server/domain/models"
)

type InvoiceRepository interface {
	CreateInvoice(ctx context.Context, invoice *models.Invoice) error
	FindInvoicesByCompanyIDAndPaidDueDateRange(ctx context.Context, companyID string, startDate, endDate *time.Time) (models.InvoiceList, error)
}
