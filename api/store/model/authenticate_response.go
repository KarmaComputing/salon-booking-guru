package model

type AuthenticateResponse struct {
	AccountInfo AccountInfo `json:"accountInfo"`
	Token       Token       `json:"token"`
}
