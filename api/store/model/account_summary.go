package model

type AccountSummary struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	MobileNumber *string `json:"mobileNumber"`
	RoleName     string  `json:"roleName"`
}
