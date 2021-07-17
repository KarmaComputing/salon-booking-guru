package model

type Account struct {
	Id           int     `json:"id"`
	RoleId       int     `json:"roleId"`
	FirstName    string  `json:"firstName"`
	LastName     string  `json:"lastName"`
	Email        string  `json:"email"`
	Password     string  `json:"password"`
	MobileNumber *string `json:"mobileNumber"`
}
