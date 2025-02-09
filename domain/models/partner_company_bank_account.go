package models

import "github.com/kohge2/upsdct-server/utils"

type PartnerCompanyBankAccount struct {
	SoftModel
	ID                string
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

func (m *PartnerCompanyBankAccount) SetDecryptedAccountNumber(key string) (string, error) {
	accountNumber, err := utils.Decrypt(m.AccountNumber, key)
	if err != nil {
		return "", err
	}

	m.AccountNumber = accountNumber
	return accountNumber, nil
}

func (m *PartnerCompanyBankAccount) EncryptedAccountNumber(key string) error {
	encrypted, err := utils.Encrypt(m.AccountNumber, key)
	if err != nil {
		return err
	}
	m.AccountNumber = encrypted
	return nil
}

func (m *PartnerCompanyBankAccount) SetDecryptedAccountHolderName(key string) (string, error) {
	name, err := utils.Decrypt(m.AccountHolderName, key)
	if err != nil {
		return "", err
	}

	m.AccountHolderName = name
	return name, nil
}

func (m *PartnerCompanyBankAccount) EncryptedAccountHolderName(key string) error {
	encrypted, err := utils.Encrypt(m.AccountHolderName, key)
	if err != nil {
		return err
	}
	m.AccountHolderName = encrypted
	return nil
}
