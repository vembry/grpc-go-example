package main

import (
	"context"
	"grpc-go-example/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	//connecting to server grpc address
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("client failed to connect to grpc: %s", err)
	}
	defer conn.Close()

	//connecting to specific grpc scope(?)
	c := proto.NewChatServiceClient(conn)

	//preparing object to be sent to server's grpc
	message := proto.Message{
		Body: "Hello",
	}

	//calling server's rpc command
	response, err1 := c.SayHello(context.Background(), &message)
	if err1 != nil {
		log.Fatalf("error calling grpc. message: %s", err1)
	}

	log.Printf("response: %s", response.Body)
}
