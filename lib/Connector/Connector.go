package Connector

import (
	"net/http"
	"strings"
	"log"
	"io/ioutil"
	"fmt"
)

type Connector struct {
	username string
	password string
	hostname string
	basepath string
}

func (connector *Connector) Get(address string, payload string) string {
	var url string
	url = "https://" + connector.hostname + "/__API__" + address
	client := &http.Client{}
	request, err := http.NewRequest("PUT", url, strings.NewReader(payload))
	request.SetBasicAuth(connector.username, connector.password)
	//request.ContentLength = int64(len(payload))
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("The calculated length is:", len(string(contents)), "for the url:", url)
		fmt.Println("   ", response.StatusCode)
		hdr := response.Header
		for key, value := range hdr {
			fmt.Println("   ", key, ":", value)
		}
		fmt.Println(string(contents))
	}
	return ""
}

func (connector *Connector) Put(address string, payload string) string {
	var url string
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
		fmt.Println("The calculated length is:", len(string(contents)), "for the url:", url)
		fmt.Println("   ", response.StatusCode)
		hdr := response.Header
		for key, value := range hdr {
			fmt.Println("   ", key, ":", value)
		}
		fmt.Println(contents)
	}
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