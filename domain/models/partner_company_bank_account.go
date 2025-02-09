package models

type PartnerCompanyBankAccount struct {
	SoftModel
	PartnerCompanyID  string
	BankName          string
	BranchName        string
	AccountType       AccountType
	AccountNumber     string
	AccountHolderName string
}

func (PartnerCompanyBankAccount) TableName() string {
	return "partner_company_bank_accounts"
}

type AccountType string

const (
	AccountTypeChecking AccountType = "checking"
	AccountTypeSavings  AccountType = "savings"
)

func (t AccountType) String() string {
	return string(t)
}

func (t AccountType) Label() string {
	switch t {
	case AccountTypeChecking:
		return "当座預金"
	case AccountTypeSavings:
		return "普通預金"
	}
	return ""
}
