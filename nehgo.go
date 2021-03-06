package nehgo

import (
	"fmt"
	"nehgo/lib/Order"
	"nehgo/lib/Connector"
	"nehgo/lib/Customer"
	"strings"
)

func GetOrder(connector *Connector.Connector, order_id int) *Order.Order {
	var apiResult *Connector.ConnectorResponse
	var order *Order.Order
	var address string
	address = fmt.Sprintf("/order/%d?view=long", order_id)
	apiResult = connector.Get(address)
	if apiResult == nil {
		return nil
	}
	order = Order.NewOrder(apiResult.Content)
	return order
}

func GetOrderHREF(connector *Connector.Connector, address string) *Order.Order {
	var apiResult *Connector.ConnectorResponse
	var order *Order.Order

	if strings.Index(address, "/__API__") < 0 {
		address = "/__API__" + address
	}
	address += "?view=long"
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
	address = fmt.Sprintf("/customer/%d?view=long", customer_id)
	apiResult = connector.Get(address)
	if apiResult == nil {
		return nil
	}
	customer = Customer.NewCustomer(apiResult.Content)
	return customer
}

func GetCustomerOrders(connector *Connector.Connector, customer_id int) *Order.OrderList {
	var apiResult *Connector.ConnectorResponse
	var orderlist *Order.OrderList
	var address string
	address = fmt.Sprintf("/customer/%d/orders?view=long", customer_id)
	apiResult = connector.Get(address)
	if apiResult == nil {
		return nil
	}
	orderlist = Order.NewOrderList(apiResult.Content)
	return orderlist
}
