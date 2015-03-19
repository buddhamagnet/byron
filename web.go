package main

import (
	"log"
	"net"
	"net/http"
	"time"
)

// Function to use for transport dial.
func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, time.Duration(2*time.Second))
}

func newClient() http.Client {
	transport := http.Transport{
		Dial: dialTimeout,
	}
	client := http.Client{
		Transport: &transport,
	}
	return client
}

func get(endpoint string, client http.Client) *http.Response {
	resp, err := client.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func getInstanceId(client http.Client) *http.Response {
	return get("http://169.254.169.254/latest/meta-data/instance-id", client)
}

func getDevices(token string, client http.Client) *http.Response {
	return get("https://api.serverdensity.io/inventory/devices?token="+token, client)
}

func getInstaller(script string, client http.Client) *http.Response {
	return get("https://www.serverdensity.com/downloads/"+script, client)
}
