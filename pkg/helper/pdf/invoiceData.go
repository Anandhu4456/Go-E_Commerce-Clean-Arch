package pdf

import "errors"

type InvoiceData struct {
	Title       string
	Quantity    int
	Price       int
	TotalAmount int
}

func (d *InvoiceData) CalculateTotalAmount() int {
	totalAmount := d.Quantity * d.Price
	return totalAmount
}

func (d *InvoiceData) ReturnItemTotalAmount() float64 {
	totalAmount := d.CalculateTotalAmount()
	converted := float64(totalAmount) / 100
	return converted
}

func (d *InvoiceData) ReturnItemPrice() float64 {
	returnPrice := float64(d.Price) / 100
	return returnPrice
}

func NewInvoiceData(title string, qty int, price interface{}) (*InvoiceData, error) {
	var convertedPrice int

	switch priceValue := price.(type) {
	case int:
		convertedPrice = priceValue * 100
	case float32:
		convertedPrice = int(priceValue) * 100
	case float64:
		convertedPrice = int(priceValue) * 100
	default:
		return nil, errors.New("type not permitted")
	}
	return &InvoiceData{
		Title:    title,
		Quantity: qty,
		Price:    convertedPrice,
	}, nil
}
