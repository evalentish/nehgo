package Connector

type Connector struct {
	username string
	password string
	hostname string
}

func (connector *Connector) Get(address string, payload string) {

}

func (connector *Connector) Put(address string, payload string) {

}

func (connector *Connector) Post(address string, payload string) {

}

func NewConnector(username string, password string, hostname string) *Connector {
	connector := new(Connector)
	connector.username = username
	connector.password = password
	connector.hostname = hostname
	return connector
}