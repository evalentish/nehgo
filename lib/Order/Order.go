package Order

import "go/types"

type Order struct {
	Order_id int
}

func (order *Order) Parse(data string) {
	order.Order_id = 123
}

func NewOrder(payload ...string) *Order {
	var order *Order
	order = new(Order)
	if len(payload) > 0 {
		order.Parse(payload[0])
	}
	return order
}

func (order *Order) AsXML() string {
	var xml string
	xml = "<?xml version=\"1.0\" encoding=\"utf-8\" ?>"
	return xml
}

type Orders struct {
	List types.Array
}

func (orders *Orders) Parse(data string) {

}

func (orders *Orders) AsXML() string {
	var xml string
	xml = "<?xml version=\"1.0\" encoding=\"utf-8\" ?>"
	return xml
}

