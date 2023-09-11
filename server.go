package main

import (
	"fmt"
	"sync"

	"net"
)

type Server struct {
	Ip   string
	Port int

	//	a list of online clients
	OnlineMap map[string]*Client
	mapLock   sync.RWMutex
	Message   chan string
}

func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*Client),
		Message:   make(chan string),
	}

}

func (this *Server) Broadcast(msg string, client *Client) {
	sendMsg := "[" + client.Address + "]: " + msg

	this.Message <- sendMsg
}

// listen messages of goroutine. when a message is sent, all online users would receive the message

func (this *Server) ListenBroadcast() {
	for {
		msg := <-this.Message
		this.mapLock.Lock()

		for _, cli := range this.OnlineMap {

			cli.C <- msg

		}
		this.mapLock.Unlock()

	}

}

func (this *Server) handle(conn net.Conn) {

	client := NewClient(conn)

	// add to the list of online client

	this.mapLock.Lock()
	this.OnlineMap[client.Name] = client
	this.mapLock.Unlock()

	// broadcast online information
	this.Broadcast("online", client)

	//	receive message
	go func() {
		buff := make([]byte, 4096)
		for {
			n, err := conn.Read(buff)

			if n == 0 {
				this.Broadcast("offline", client)
				return
			}

			if err != nil {
				fmt.Printf("read error: ", err)
				return
			}
			msg := string(buff[:n-1])

			this.Broadcast(msg, client)

		}
	}()

}

func (this *Server) Start() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))

	if err != nil {
		fmt.Printf("net.listen err is ", err, err)
		return
	}
	defer listen.Close()

	go this.ListenBroadcast()

	for {
		accept, err := listen.Accept()

		if err != nil {
			fmt.Printf("listen accept error is", err, err)
			continue
		}
		go this.handle(accept)

	}

}
