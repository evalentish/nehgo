package Customer

import (
	"bytes"
	"encoding/xml"
	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
	"strings"
)

type Customer struct {
	XMLName xml.Name `xml:"Customer"`
	Customer_id int `xml:"customer_id,attr"`
	Type string `xml:"Type"`
	OneTimeCustomer bool `xml:"OneTimeCustomer"`
	Mode string `xml:"Mode"`
	CustomerNo string `xml:"CustomerNo"`
	FirstName string `xml:"FirstName"`
	LastName string `xml:"LastName"`
	Address string `xml:"Address1"`
	Address2 string `xml:"Address2"`
	ZIP string `xml:"ZIP"`
	City string `xml:"City"`
	State string `xml:"State"`
	Language string `xml:"Language"`
	PhoneNo string `xml:"PhoneNo"`
	CellPhoneNo string `xml:"CellPhoneNo"`
	FaxNo string `xml:"FaxNo"`
	EMail string `xml:"EMail"`
	Country string `xml:"Country"`
	Company string `xml:"Company"`
	VatNo string `xml:"VatNo"`
	OrgNo string `xml:"OrgNo"`
	ExtraText1 string `xml:"ExtraText1"`
	ExtraText2 string `xml:"ExtraText2"`
	ExtraText3 string `xml:"ExtraText3"`
	ExtraText4 string `xml:"ExtraText4"`
	ExtraText5 string `xml:"ExtraText5"`
	ExtraText6 string `xml:"ExtraText6"`
	ExtraSelector1 int `xml:"ExtraSelector1"`
	ExtraSelector2 int `xml:"ExtraSelector2"`
	ExtraSelector3 int `xml:"ExtraSelector3"`
	ExtraSelector4 int `xml:"ExtraSelector4"`
	ExtraSelector5 int `xml:"ExtraSelector5"`
	ExtraSelector6 int `xml:"ExtraSelector6"`
	ExtraSelector7 int `xml:"ExtraSelector7"`
	ExtraSelector8 int `xml:"ExtraSelector8"`
	ExtraSelector9 int `xml:"ExtraSelector9"`
	ExtraSelector10 int `xml:"ExtraSelector10"`
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
