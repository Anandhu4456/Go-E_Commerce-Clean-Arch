package pdf

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
