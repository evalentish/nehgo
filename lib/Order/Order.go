package Order

import (
	"go/types"
	"encoding/xml"
	"fmt"
)

type Order struct {
	XMLName xml.Name `xml:"Order"`
	Order_id int `xml:"order_id,attr"`
	OrderHREF string `xml:"href,attr"`
	OrderNo int `xml:"OrderNo"`
	CustomerNo int `xml:"Customer"`
}

func (order *Order) Parse(data string) {
	o := &Order{}
	xml.Unmarshal([]byte(data), o)
	fmt.Println(data)
	fmt.Println(o)
	order = o
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

