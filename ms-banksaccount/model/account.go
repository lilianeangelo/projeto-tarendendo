package model

import(
	
)
type Account struct {
	AccountId string 		`json:"id"`
	Name string 			`json:"name"`
	User User 				`json:"user"`
}
