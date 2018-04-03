package nehgo

import (
	"fmt"
	"github.com/evalentish/nehgo/lib/Order"
	"github.com/evalentish/nehgo/lib/Connector"
	"nehgo/lib/Customer"
)

func GetOrder(connector *Connector.Connector, order_id int) *Order.Order {
	var apiResult *Connector.ConnectorResponse
	var order *Order.Order
	var address string
	address = fmt.Sprintf("/order/%d", order_id)
	apiResult = connector.Get(address)
	if apiResult == nil {
		return nil
	}
	order = Order.NewOrder(apiResult.Content)
	return order
}

func GetCustomer(connector *Connector.Connector, customer_id int) *Customer.Customer {
	var apiResult *Connector.ConnectorResponse
	var customer *Customer.Customer
	var address string
	address = fmt.Sprintf("/customer/%d", customer_id)
	apiResult = connector.Get(address)
	if apiResult == nil {
		return nil
	}
	customer = Customer.NewCustomer(apiResult.Content)
	return customer
}
