package Order

import (
	"encoding/xml"
	"strings"
	"bytes"
	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
	"time"
)

type Order struct {
	XMLName xml.Name `xml:"Order"`
	Order_id int `xml:"order_id,attr"`
	OrderHREF string `xml:"href,attr"`
	OrderNo string `xml:"OrderNo"`
	Customer struct {
		HREF string `xml:"href,attr"`
		Value string `xml:",chardata"`
	} `xml:"Customer"`
	State string `xml:"State"`
	CreatedTime string `xml:"CreatedTime"`
	CreatedTimestamp time.Time
	ChangedTime string `xml:"ChangedTime"`
	ChangedTimestamp time.Time
}

func (order *Order) Parse(data string) bool {
	reader := bytes.NewReader([]byte(data))
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReader
	err := decoder.Decode(order)
	if err != nil {
		return false
	}
	// Fix dates
	order.CreatedTimestamp, _ = time.Parse(time.RFC3339, order.CreatedTime)
	order.ChangedTimestamp, _ = time.Parse(time.RFC3339, order.ChangedTime)

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

type OrderList struct {
	XMLName xml.Name `xml:"OrderList"`
	Orders []Order `xml:"Order"`
}

func (orderlist *OrderList) Parse(data string) bool {
	reader := bytes.NewReader([]byte(data))
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReader
	err := decoder.Decode(orderlist)
	if err != nil {
		return false
	}
	return true
}

func NewOrderList(payload ...string) *OrderList {
	var orderlist *OrderList
	orderlist = new(OrderList)
	if len(payload) > 0 {
		s := strings.Join(payload, " ")
		orderlist.Parse(s)
	}
	return orderlist
}

func (orders *OrderList) AsXML() string {
	var xml string
	xml = "<?xml version=\"1.0\" encoding=\"utf-8\" ?>"
	return xml
}

