package pdf

type InvoiceData struct{
	Title string
	Quantity int
	Price int
	TotalAmount int
}

func (d *InvoiceData)CalculateTotalAmount()int{
	totalAmount:=d.Quantity *d.Price
	return totalAmount
}