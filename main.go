package main

import (
	"log"
	"net/rpc"
)

type Args struct {
	Name string
	Type string
}

func main() {
	// Address to this variable will be sent to the RPC server
	// Type of reply should be same as that specified on server
	var reply string
	args := Args{
		Name: "Logs",
		Type: "logs",
	}

	// DialHTTP connects to an HTTP RPC server at the specified network
	client, err := rpc.DialHTTP("tcp", "localhost"+":1234")
	if err != nil {
		log.Fatal("Client connection error: ", err)
	}

	// TODO: add action name to rpc call
	err = client.Call("Action.RegisterAction", &args, &reply)
	if err != nil {
		log.Fatal("Client invocation error: ", err)
	}

	log.Println(reply)
}
