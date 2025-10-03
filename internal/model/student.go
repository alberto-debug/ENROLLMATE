package model

type Student struct {
	ID        string `gorm:"primaryKey" json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Course    string `json:"course"`
}
