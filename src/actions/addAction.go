package actions

import (
	"log"
	"net/rpc"
)

type Args struct {
	Name string
	Type string
}

func AddAction(client *rpc.Client, args Args) (err error) {
	var reply string

	err = client.Call("Action.RegisterAction", args, &reply)
	if err != nil {
		log.Fatal("Client invocation error: ", err)
	}

	log.Println(reply)
	return err
}
