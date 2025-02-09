package models

type Company struct {
	SoftModel
	ID                 string
	Name               string
	RepresentativeName string
	Tel                string
	PostalCode         string
	Address            string
}

func (Company) TableName() string {
	return "companies"
}
