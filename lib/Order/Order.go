package Order

import (
	"encoding/xml"
	"strings"
	"bytes"
	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
	"time"
)

type Item struct {
	XMLName xml.Name `xml:"Item"`
	HREF string `xml:"href,attr"`
	LineNumber int `xml:"LineNumber"`
	Variant struct {
		HREF string `xml:"href,attr"`
		IsNil bool `xml:"xsi:nil,attr"`
	} `xml:"Variant"`
	SKU string `xml:"SKU"`
	Qty float32 `xml:"Qty"`
	Price float32 `xml:"Price"`
	Tax float32 `xml:"Tax"`
}

type OrderItems struct {
	XMLName xml.Name `xml:"OrderItems"`
	HREF string `xml:"href,attr"`
	Items []Item `xml:"Item"`
}

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
	OrderItems OrderItems `xml:"OrderItems"`
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

func (orderlist *OrderList) AsXML() string {
	var buffer []byte
	var err error
	buffer, err = xml.Marshal(orderlist)
	if err != nil {
		return ""
	}
	return string(buffer)
}

