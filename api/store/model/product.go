package model

type Product struct {
	Id                int     `json:"id"`
	ProductCategoryId int     `json:"productCategoryId"`
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	Price             float64 `json:"price"`
	Deposit           float64 `json:"deposit"`
	Duration          float64 `json:"duration"`
}
