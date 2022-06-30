package model

import (
	"time"
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
	CreatedAt          time.Time    `json:"created_at" bson:"created_at"`
	ValidatedAt        time.Time    `json:"validated_at" bson:"validated_at,omitempty"`
}

// Account struct
type Account struct {
	AccountId string 		`json:"id"`
	Name string 			`json:"name"`
	User User 				`json:"user"`
}


//implementar request e response