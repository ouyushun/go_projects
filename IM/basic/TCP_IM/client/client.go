package main

import "net"

type Client struct {
	IP string
	Port string
	Name string
	Conn net.Conn
}

func NewClient(IP string, Port string) *Client {
	client := &Client{
		IP: IP,
		Port: Port,
	}


	return client

}

func main() {

}
