package nehgo

import (
	"fmt"
	"github.com/evalentish/nehgo/lib/Order"
	"github.com/evalentish/nehgo/lib/Connector"
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
	order.Parse(apiResult.Content)
	return order
}