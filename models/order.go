package models

type Order struct {
	ID int
	UserID int
	Products []Product
	Total float64
	Status string
}

//Agregar producto al pedido
func (o *Order) AddProduct(p Product) {
	o.Products = append(o.Products, p)
	o.Total += p.Price
}

//Confirmar el pedido
func (o *Order) Confirm() {
	o.Status = "Confirmed"
}
