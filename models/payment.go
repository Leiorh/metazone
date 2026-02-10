package models

import "fmt"

type Payment struct {
	ID int
	OrderID int
	Amount float64
	Method string
	Status string
	
}

// Procesar el pago con validación
func (p *Payment) Process() error {
	if p.Amount <= 0 {
		return fmt.Errorf("el monto del pago debe ser mayor a cero")
	}

	p.Status = "APROBADO"
	return nil
}