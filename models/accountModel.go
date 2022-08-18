package models

type Address struct {
	AddressLine1 string `json:"address_line_1" bson:"address_line_1"`
	City         string `json:"city" bson:"city"`
	District     string `json:"district" bson:"district"`
	State        string `json:"state" bson:"state"`
	Country      string `json:"country" bson:"country"`
	Pincode      int64  `json:"pincode" bson:"pincode"`
}

type Owner struct {
	FirstName    string `json:"first_name" bson:"first_name"`
	LastName     string `json:"last_name" bson:"last_name"`
	EmailId      string `json:"email_id" bson:"email_id"`
	MobileNumber string `json:"mobile_number" bson:"mobile_number"`
	Password     string `json:"-"`
	Address      `json:"address" bson:"address"`
}

type Account struct {
	Id          string  `json:"id" bson:"_id"`
	Balance     float64 `json:"balance" bson:"balance"`
	AccountType string  `json:"account_type" bson:"account_type"`
	Owner       `json:"owner" bson:"owner"`
}

type AccountController interface {
	CreateAccount(Account) error
	GetAccount() (Account, error)
	UpdateUserDetails(Account) error
	DeleteAccount() error
}
