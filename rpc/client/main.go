package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"

	"../rpcdemo"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)

	var result float64
	err = client.Call("DemoService.Div", rpcdemo.Args{10, 3}, &result)
	fmt.Println(result, err)

	err = client.Call("DemoService.Div", rpcdemo.Args{10, 0}, &result)
	fmt.Println(result, err)
}
