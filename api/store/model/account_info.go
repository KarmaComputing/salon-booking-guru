package model

type AccountInfo struct {
	Email       string   `json:"email"`
	FirstName   string   `json:"firstName"`
	LastName    string   `json:"lastName"`
	RoleName    string   `json:"roleName"`
	Permissions []string `json:"permissions"`
}
