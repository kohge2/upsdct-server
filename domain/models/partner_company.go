package models

type PartnerCompany struct {
	SoftModel
	ID                 string
	Name               string
	RepresentativeName string
	Tel                string
	PostalCode         string
	Address            string
}

func (PartnerCompany) TableName() string {
	return "partner_companies"
}

type PartnerCompanyEmbed struct {
	PartnerCompany
	PartnerCompanyBankAccount *PartnerCompanyBankAccount `gorm:"foreignKey:ID;references:PartnerCompanyID"`
}

type PartnerCompanyEmbedList []*PartnerCompanyEmbed
