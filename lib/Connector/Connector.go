package Connector

import (
	"net/http"
	"strings"
	"log"
	"io/ioutil"
)

type Connector struct {
	username string
	password string
	hostname string
	basepath string
}

type ConnectorResponse struct {
	StatusCode int
	Headers map[string]string
	Content string
}

func (connector *Connector) Get(address string) *ConnectorResponse {
	var url string
	var result = new(ConnectorResponse)
	result.Headers = make(map[string]string)
	url = "https://" + connector.hostname + "/__API__" + address
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.SetBasicAuth(connector.username, connector.password)
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		hdr := response.Header
		var key, value string
		for key, value = range hdr {
			result.Headers[key] = value
		}
		result.StatusCode = response.StatusCode
		result.Content = string(contents)
	}
	return nil
}

func (connector *Connector) Put(address string, payload string) *ConnectorResponse {
	var url string
	var result = new(ConnectorResponse)
	url = "https://" + connector.hostname + address
	client := &http.Client{}
	request, err := http.NewRequest("PUT", url, strings.NewReader(payload))
	request.SetBasicAuth(connector.username, connector.password)
	request.ContentLength = int64(len(payload))
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		hdr := response.Header
		var key, value string
		for key, value = range hdr {
			result.Headers[key] = value
		}
		result.StatusCode = response.StatusCode
		result.Content = string(contents)
		return result
	}
	return nil
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