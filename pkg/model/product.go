package model

type Product struct {
	ID          int32 `json:"id"`
	Category    `json:"category"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Discount    int     `json:"discount"`
}

func (p Product) NewProduct(id int32, name, desc string, price float32, discount int) *Product {
	return &Product{
		ID:          id,
		Name:        name,
		Description: desc,
		Price:       price,
		Discount:    discount,
	}
}
