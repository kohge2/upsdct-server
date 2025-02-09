package request

import (
	"time"

	"github.com/gin-gonic/gin"
)

type CreateInvoiceRequest struct {
	PartnerCompanyID string `json:"partnerCompanyID" validate:"required"`
	PaidDueDate      string `json:"paidDueDate" validate:"required"`
	PaidAmount       int    `json:"paidAmount" validate:"required"`
}

func (r *CreateInvoiceRequest) Bind(c *gin.Context) error {
	if err := c.BindJSON(r); err != nil {
		return err
	}
	return nil
}

// RFC3339(指定した日付フォーマット, ex 2025-05-31T22:00:00+09:00)かのバリデーション
func (r *CreateInvoiceRequest) GetPaidDueDate() (*time.Time, error) {
	paidDueDate, err := time.Parse(time.RFC3339, r.PaidDueDate)
	if err != nil {
		return nil, err
	}
	return &paidDueDate, nil
}
