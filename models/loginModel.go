package models

type Login struct {
	AccountId string `json:"account_id" bson:"account_id"`
	Password  string `json:"password" bson:"password"`
}
