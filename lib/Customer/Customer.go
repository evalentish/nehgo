package Customer

import (
	"bytes"
	"encoding/xml"
	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
	"strings"
)

type Customer struct {
	Customer_id int
	Name string
	Address string
	Customer_no string
}

func (customer *Customer) Parse(data string) bool {
	reader := bytes.NewReader([]byte(data))
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReader
	err := decoder.Decode(customer)
	if err != nil {
		return false
	}
	return true
}

func (customer *Customer) AsXML() string {
	var xml string
	xml = "<?xml version=\"1.0\" encoding=\"utf-8\" ?>"
	return xml
}

func NewCustomer(payload ...string) *Customer {
	var customer *Customer
	customer = new(Customer)
	if len(payload) > 0 {
		s := strings.Join(payload, " ")
		customer.Parse(s)
	}
	return customer
}
