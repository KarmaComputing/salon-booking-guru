package store

import "salon-booking-guru/store/model"

var store Store

type Store interface {
	Account() AccountStore
	Authenticate() AuthenticateStore
	Authorize() AuthorizeStore
	Availability() AvailabilityStore
	Product() ProductStore
	ProductCategory() ProductCategoryStore
	Qualification() QualificationStore
	Role() RoleStore
	Token() TokenStore
}

type AccountStore interface {
	GetAll() ([]model.Account, error)
	Get(id int) (model.Account, error)
	GetInfo(id int) (model.AccountInfo, error)
	Create(product *model.Account) error
	Update(product *model.Account) error
	Delete(id int) error
	UpsertQualification(id int, qualificationIds []int) error
}

type AuthenticateStore interface {
	AuthenticateCredentials(credentials model.Credentials) (model.AuthenticateResponse, error)
}

type AuthorizeStore interface {
	AuthorizeToken(bearerToken string, requiredPermissions []string) error
}

type AvailabilityStore interface {
	GetAll() ([]model.Availability, error)
	GetAllByAccountId(accountId int) ([]model.Availability, error)
	Get(id int) (model.Availability, error)
	CreateMultiple(availability []model.Availability) error
	Update(availability *model.Availability) error
	Delete(id int) error
}

type ProductStore interface {
	GetAll() ([]model.Product, error)
	Get(id int) (model.Product, error)
	Create(product *model.Product) error
	Update(product *model.Product) error
	Delete(id int) error
}

type ProductCategoryStore interface {
	GetAll() ([]model.ProductCategory, error)
	Get(id int) (model.ProductCategory, error)
	Create(qualification *model.ProductCategory) error
	Update(qualification *model.ProductCategory) error
	Delete(id int) error
}

type QualificationStore interface {
	GetAll() ([]model.Qualification, error)
	GetAllNameByAccountId(accountId int) ([]string, error)
	GetAllNameByProductId(productId int) ([]string, error)
	Get(id int) (model.Qualification, error)
	Create(qualification *model.Qualification) error
	Update(qualification *model.Qualification) error
	Delete(id int) error
}

type RoleStore interface {
	GetAll() ([]model.Role, error)
}

type TokenStore interface {
	Generate(accountId int) (model.Token, error)
	Delete(id int) error
	DeleteAllByAccountId(accountId int) error
}
