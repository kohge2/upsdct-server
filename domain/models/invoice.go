package models

import (
	"fmt"
	"time"

	"github.com/kohge2/upsdct-server/utils"
)

type Invoice struct {
	SoftModel
	ID               string
	CompanyID        string
	PartnerCompanyID string
	PublishedDate    time.Time // 発行日
	Commission       *int      // 手数料
	CommissionRate   *float64  // 手数料率
	TaxRate          *float64  // 消費税率
	Tax              *int      // 消費税
	PaidAmount       int       // 支払金額
	BilledAmount     *int      // 請求金額(支払金額 に手数料4%を加えたものに更に手数料の消費税を加えたもの)
	PaidDueDate      time.Time // 支払期限
	InvoiceStatus    InvoiceStatus
	CreatedBy        string // 作成したUserID
}

func (m Invoice) TableName() string {
	return "invoices"
}

func (m *Invoice) CalcCommission() (int, error) {
	if m.CommissionRate == nil {
		return 0, fmt.Errorf("commission rate is not set")
	}
	return utils.MultiplyIntByDecimal(m.PaidAmount, *m.CommissionRate), nil
}

func (m *Invoice) CalcBilledAmount(taxRate float64) (int, int, error) {
	if m.Commission == nil {
		return 0, 0, fmt.Errorf("commission is not set")
	}

	tax := utils.MultiplyIntByDecimal(*m.Commission, taxRate)

	return m.PaidAmount + *m.Commission + tax, tax, nil
}

type InvoiceList []*Invoice

func (l InvoiceList) UniquePartnerCompanyIDs() []string {
	uniquePartnerCompanyIDs := make([]string, 0, len(l))
	uniquePartnerCompanyIDsMap := make(map[string]struct{}, len(l))
	for _, invoice := range l {
		if _, ok := uniquePartnerCompanyIDsMap[invoice.PartnerCompanyID]; !ok {
			uniquePartnerCompanyIDs = append(uniquePartnerCompanyIDs, invoice.PartnerCompanyID)
			uniquePartnerCompanyIDsMap[invoice.PartnerCompanyID] = struct{}{}
		}
	}
	return uniquePartnerCompanyIDs
}

type InvoiceEmbed struct {
	Invoice
	PartnerCompany *PartnerCompanyEmbed
}

type InvoiceEmbedList []*InvoiceEmbed

func NewInvoiceEmbedList(invoices InvoiceList, partnerCompanies PartnerCompanyEmbedList) InvoiceEmbedList {
	partnerCompanyMap := make(map[string]*PartnerCompanyEmbed, len(partnerCompanies))
	for _, partnerCompany := range partnerCompanies {
		partnerCompanyMap[partnerCompany.ID] = partnerCompany
	}

	embedList := make(InvoiceEmbedList, 0, len(invoices))
	for _, invoice := range invoices {
		embedList = append(embedList, &InvoiceEmbed{
			Invoice:        *invoice,
			PartnerCompany: partnerCompanyMap[invoice.PartnerCompanyID],
		})
	}
	return embedList
}
