package main

import (
	"context"
	"fmt"
	"log"
	"logger/data"
)

type RPCServer struct{}

type RPCPayload struct {
	Name string
	Data string
}

func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name: payload.Name,
		Data: payload.Data,
	})
	if err != nil {
		log.Println("Error writing the log:", err)
		return err
	}

	*resp = fmt.Sprintf("Processed payload via RPC:%s", payload.Name)
	return nil
}
