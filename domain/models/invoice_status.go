package models

type InvoiceStatus string

const (
	InvoiceStatusDraft      InvoiceStatus = "draft"
	InvoiceStatusProcessing InvoiceStatus = "processing"
	InvoiceStatusPaid       InvoiceStatus = "paid"
	InvoiceStatusError      InvoiceStatus = "error"
)

func (s InvoiceStatus) String() string {
	return string(s)
}

func (s InvoiceStatus) IsValid() bool {
	switch s {
	case InvoiceStatusDraft, InvoiceStatusProcessing, InvoiceStatusPaid, InvoiceStatusError:
		return true
	}
	return false
}

func (s InvoiceStatus) Label() string {
	switch s {
	case InvoiceStatusDraft:
		return "未処理"
	case InvoiceStatusProcessing:
		return "処理中"
	case InvoiceStatusPaid:
		return "支払い済み"
	case InvoiceStatusError:
		return "エラー"
	}
	return ""
}
