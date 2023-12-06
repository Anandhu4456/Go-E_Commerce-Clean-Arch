package pdf

type Invoice struct {
	Name        string
	Address     string
	InvoiceItem []*InvoiceData
}

func CreateInvoice(name string, address string, invoiceItem []*InvoiceData) *Invoice {
	return &Invoice{
		Name:        name,
		Address:     address,
		InvoiceItem: invoiceItem,
	}
}

func (i *Invoice) CalculateInvoiceTotalAmount() float64 {
	var invoiceTotalAmount int = 0
	for _, data := range i.InvoiceItem {
		amount := data.CalculateTotalAmount()
		invoiceTotalAmount += amount
	}
	totalAmount := float64(invoiceTotalAmount) / 100
	return totalAmount
}
