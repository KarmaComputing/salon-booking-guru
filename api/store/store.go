package store

import "salon-booking-guru/store/model"

var store Store

type Store interface {
	Account() AccountStore
	Authenticate() AuthenticateStore
	Authorize() AuthorizeStore
	Qualification() QualificationStore
	Token() TokenStore
}

type AccountStore interface {
	GetAll() ([]model.Account, error)
	Get(id int) (model.Account, error)
	GetInfo(id int) (model.AccountInfo, error)
	Create(product *model.Account) error
	Update(product *model.Account) error
	Delete(id int) error
}

type AuthenticateStore interface {
	AuthenticateCredentials(credentials model.Credentials) (model.AuthenticateResponse, error)
}

type AuthorizeStore interface {
	AuthorizeToken(bearerToken string, requiredPermissions []string) error
}

type QualificationStore interface {
	GetAll() ([]model.Qualification, error)
	Get(id int) (model.Qualification, error)
	Create(qualification *model.Qualification) error
	Update(qualification *model.Qualification) error
	Delete(id int) error
}

type TokenStore interface {
	Generate(accountId int) (model.Token, error)
	Delete(id int) error
	DeleteAllByAccountId(accountId int) error
}
