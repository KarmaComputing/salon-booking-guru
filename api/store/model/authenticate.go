package model

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticateResponse struct {
	AccountInfo AccountInfo `json:"accountInfo"`
	Token       Token       `json:"token"`
}
