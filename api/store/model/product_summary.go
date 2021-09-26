package model

type ProductSummary struct {
	Id                  int     `json:"id"`
	ProductCategoryName string  `json:"productCategoryName"`
	Name                string  `json:"name"`
	Description         string  `json:"description"`
	Price               float64 `json:"price"`
	Deposit             float64 `json:"deposit"`
	Duration            float64 `json:"duration"`
}
