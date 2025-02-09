package models

type User struct {
	SoftModel
	ID        string
	CompanyID string
	Email     string
	Password  string
}
