package nehgo

type Customer struct {
	Customer_id int
	Name string
	Address string
	Customer_no string
}

func (customer *Customer) GetCustomer(customer_id int) {
	customer.Customer_id = customer_id
	customer.Name = "Nisse Hult"
	customer.Address = "Testvägen 1"
}