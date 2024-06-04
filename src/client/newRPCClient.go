package client

import (
	"log"
	"net/rpc"
)

func NewRPCClient() (client *rpc.Client, err error) {
	client, err = rpc.DialHTTP("tcp", "127.0.0.1:9854")
	if err != nil {
		log.Fatal("Failed to create RPC client: ", err)
	}
	return client, err
}
