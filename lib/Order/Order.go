package Order

import (
	Connector "github.com/evalentish/nehgo/lib/Connector"
	"go/types"
)

type Order struct {
	Order_id int
}

func (order *Order) Parse(data string) {
	order.Order_id = 123
}

type Orders struct {
	List types.Array
}

func (orders *Orders) Parse(data string) {

}

func GetOrder(order_id int, connector *Connector.Connector) *Order {
	var order = new(Order)

	result := connector.Get("", "")
	order.Parse(result);

	return order
}

func GetOrders(connector *Connector.Connector) *Orders {
	var orders = new(Orders)

	result := connector.Get("", "")
	orders.Parse(result)

	return orders
}