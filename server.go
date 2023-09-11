package main

import (
	"fmt"

	"net"
)

type Server struct {
	Ip   string
	Port int
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}

	return server
}

func (this *Server) handle(conn net.Conn) {

	fmt.Printf("connection is success")
}

func (this *Server) Start() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))

	if err != nil {
		fmt.Printf("net.listen err is ", err)
		return
	}
	defer listen.Close()

	for {
		accept, err := listen.Accept()

		if err != nil {
			fmt.Printf("listen accept error is", err)
			continue
		}
		go this.handle(accept)

	}

}
