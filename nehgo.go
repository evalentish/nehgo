package nehgo

import (
	"fmt"
	"nehgo/lib/Order"
	"nehgo/lib/Connector"
)

func init() {
	fmt.Print("Test")
}

func GetOrder(connector *Connector.Connector, order_id int) *Order.Order {
	var apiResult string
	var order *Order.Order
	var address string
	address = fmt.Sprintf("/order/%d", order_id)
	apiResult = connector.Get(address, "")
	order = Order.NewOrder(apiResult)
	order = Order.NewOrder()
	order.Parse(apiResult)
	return order
}