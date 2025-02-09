package request

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kohge2/upsdct-server/utils"
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
		return nil, utils.NewAppValidateByErr("paidDueDate")
	}
	return &paidDueDate, nil
}

type GetInvoicesRequest struct {
	StartDate string `query:"startDate"`
	EndDate   string `query:"endDate"`
}

func (r *GetInvoicesRequest) Bind(c *gin.Context) error {
	r.StartDate = c.Query("startDate")
	r.EndDate = c.Query("endDate")
	return nil
}

func (r *GetInvoicesRequest) GetStartDateAndEndDate() (*time.Time, *time.Time, error) {
	var startDate, endDate *time.Time
	if r.StartDate != "" {
		date, err := time.Parse(time.RFC3339, r.StartDate)
		if err != nil {
			return nil, nil, utils.NewAppValidateByErr("startDate")
		}
		startDate = &date
	}

	if r.EndDate != "" {
		date, err := time.Parse(time.RFC3339, r.EndDate)
		if err != nil {
			return nil, nil, utils.NewAppValidateByErr("endDate")
		}
		endDate = &date
	}
	return startDate, endDate, nil
}
