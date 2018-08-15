package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"../rpcdemo"
)

func main() {
	rpc.Register(rpcdemo.DemoService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
		}

		go jsonrpc.ServeConn(conn)
	}
}
