package actions

import (
	"log"
	"net/rpc"
)

func AddAction(client *rpc.Client) (err error) {
	type Args struct {
		Name string
		Type string
	}

	var reply string

	// Create a new instance of Args and assign it to the args variable
	args := Args{
		Name: "Logs",
		Type: "logs",
	}

	err = client.Call("Action.RegisterAction", args, &reply)
	if err != nil {
		log.Fatal("Client invocation error: ", err)
	}

	log.Println(reply)
	return err
}
