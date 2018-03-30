package Order

import (
	Connector "github.com/evalentish/nehgo/lib/Connector"
	"go/types"
)

type Order struct {
	Order_id int
}

type Orders struct {
	List types.Array
}

func GetOrder(order_id int, connector Connector.Connector) *Order {
	connector.Get("", "")
	var order = new(Order)
	order.Order_id = order_id
	return order
}

func GetOrders(connector Connector.Connector) *Orders {
	var orders = new(Orders)

	connector.Get("", "")

	return orders
}