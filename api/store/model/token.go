package model

type Token struct {
	Id        int    `json:"id"`
	AccountId int    `json:"accountId"`
	Token     string `json:"token"`
}
