package model

type Product struct {
	ID          int32   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

func NewProduct(id int32, name string, desc string, price float32) *Product {
	return &Product{
		ID:          id,
		Name:        name,
		Description: desc,
		Price:       price,
	}
}
