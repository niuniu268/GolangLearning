package main

import "net"

type Client struct {
	Name    string
	Address string
	C       chan string
	conn    net.Conn
}

// Create a listener to deliver message to users
func (this *Client) ListenMsg() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}

//Create a user's API

func NewClient(conn net.Conn) *Client {

	clientAddr := conn.RemoteAddr().String()

	client := &Client{
		Name:    clientAddr,
		Address: clientAddr,
		C:       make(chan string),
		conn:    conn,
	}
	// launch listener

	go client.ListenMsg()

	return client

}
