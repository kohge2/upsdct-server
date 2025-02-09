package models

import (
	"fmt"
	"time"
)

type Invoice struct {
	SoftModel
	ID               string
	CompanyID        string
	PartnerCompanyID string
	PublishDate      time.Time // 発行日
	Commission       *int      // 手数料
	CommissionRate   *float64  // 手数料率
	PaidAmount       int       // 支払金額
	BilledAmount     *int      // 請求金額(支払金額 に手数料4%を加えたものに更に手数料の消費税を加えたもの)
	PaidDueDate      time.Time // 支払期限
	InvoiceStatus    InvoiceStatus
	CreatedBy        string // 作成したUserID
}

func (m Invoice) TableName() string {
	return "invoices"
}

func (m Invoice) CalcCommission() (int, error) {
	if m.CommissionRate == nil {
		return 0, fmt.Errorf("commission rate is not set")
	}
	return int(float64(m.PaidAmount) * *m.CommissionRate), nil
}

func (m Invoice) CalcBilledAmount(taxRate float64) (int, error) {
	if m.Commission == nil {
		return 0, fmt.Errorf("commission is not set")
	}
	return m.PaidAmount + int(float64(*m.Commission)*(1+taxRate)), nil
}

func (m *Invoice) SetCommission(commission int) {
	m.Commission = &commission
}

func (m *Invoice) SetBilledAmount(billedAmount int) {
	m.BilledAmount = &billedAmount
}
