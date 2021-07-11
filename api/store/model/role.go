package model

type Role struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}
