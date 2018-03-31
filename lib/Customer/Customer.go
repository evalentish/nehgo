package Customer

type Customer struct {
	Customer_id int
	Name string
	Address string
	Customer_no string
}

func (customer *Customer) Parse(data string) bool {
	return true
}

func (customer *Customer) AsXML() string {
	var xml string
	xml = "<?xml version=\"1.0\" encoding=\"utf-8\" ?>"
	return xml
}

func NewCustomer() *Customer {
	var customer *Customer
	customer = new(Customer)
	return customer
}
