package model

import (

)

// User struct
type User struct {
	UserId string 		`json:"id"`
	Name string 		`json:"name"`
	Email string 		`json:"email"`
	Password string 	`json:"password"`
	BirthDate string 	`json:"birth_date"`
	CPF string 			`json:"cpf"`
	Phone string 		`json:"phone"`
	Address string 		`json:"address"`
	City string 		`json:"city"`
	State string 		`json:"state"`
	Country string 		`json:"country"`
}

// Account struct
type Account struct {
	AccountId string 		`json:"id"`
	Name string 			`json:"name"`
	User User 				`json:"user"`
}


//implementar request e response