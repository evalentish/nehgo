package Connector

type Connector struct {
	username string
	password string
	hostname string
	basepath string
}

func (connector *Connector) Get(address string, payload string) string {
	return ""
}

func (connector *Connector) Put(address string, payload string) string {

	return ""
}

func (connector *Connector) Post(address string, payload string) string {
	return ""
}

func (connector *Connector) Delete(address string, payload string) string {
	return ""
}

func NewConnector(username string, password string, hostname string) *Connector {
	connector := new(Connector)
	connector.username = username
	connector.password = password
	connector.hostname = hostname
	connector.basepath = "/__API__/"
	return connector
}