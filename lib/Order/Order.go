package Order

import (
	"go/types"
	"encoding/xml"
)

type Order struct {
	XMLName xml.Name `xml:"Order"`
	Order_id int `xml:"order_id,attr"`
	OrderNo int `xml:"OrderNo"`
	CustomerNo int `xml:"Customer"`
}

func (order *Order) Parse(data string) {
	xml.Unmarshal([]byte(data), order)
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

