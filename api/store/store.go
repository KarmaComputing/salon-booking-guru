package store

import "salon-booking-guru/store/model"

var store Store

type Store interface {
	Account() AccountStore
	Auth() AuthStore
	Permission() PermissionStore
	Role() RoleStore
	Token() TokenStore
}

type AccountStore interface {
	GetAll() ([]model.Account, error)
	Get(id int) (model.Account, error)
	Create(product *model.Account) error
	Update(product *model.Account) error
	Delete(id int) error
	IsAuthorized(id int, permissionName string) (bool, error)
}

type AuthStore interface {
	LogIn(account *model.Account) (model.Token, error)
}

type PermissionStore interface {
	Get(id int) (model.Permission, error)
	GetAll() ([]model.Permission, error)
	Create(permission *model.Permission) error
	Update(permission *model.Permission) error
	Delete(id int) error
	GetAllNameByAccountId(accountId int) ([]string, error)
}

type RoleStore interface {
	Get(id int) (model.Role, error)
	GetAll() ([]model.Role, error)
	Create(role *model.Role) error
	Update(role *model.Role) error
	Delete(id int) error
}

type TokenStore interface {
	Delete(id int) error
	DeleteAllByAccountId(accountId int) error
	Check(token *model.Token) error
}
