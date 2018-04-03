package Order

import (
	"go/types"
	"encoding/xml"
	"strings"
	"bytes"
	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
)

type Order struct {
	XMLName xml.Name `xml:"Order"`
	Order_id int `xml:"order_id,attr"`
	OrderHREF string `xml:"href,attr"`
	OrderNo string `xml:"OrderNo"`
	CustomerNo int `xml:"Customer"`
	State string `xml:"State"`
}

func (order *Order) Parse(data string) bool {
	reader := bytes.NewReader([]byte(data))
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReader
	err := decoder.Decode(order)
	if err != nil {
		return false
	}
	return true
}

func NewOrder(payload ...string) *Order {
	var order *Order
	order = new(Order)
	if len(payload) > 0 {
		s := strings.Join(payload, " ")
		order.Parse(s)
	}
	return order
}

func (order *Order) AsXML() string {
	var buffer []byte
	var err error
	buffer, err = xml.Marshal(order)
	if err != nil {
		return ""
	}
	return string(buffer)
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

