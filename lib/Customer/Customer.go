package Customer

import (
	Connector "github.com/evalentish/nehgo/lib/Connector"
)

type Customer struct {
	Customer_id int
	Name string
	Address string
	Customer_no string
}

func (customer *Customer) GetCustomer(customer_id int, connector Connector.Connector) {
	connector.Put("", "")
	customer.Customer_id = customer_id
	customer.Name = "Nisse Hult"
	customer.Address = "Testvägen 1"
}

func (customer *Customer) PutCustomer() {
	// Send this object to the server
}

func (customer *Customer) PostCustomer() {
	// Send this object to the server
}