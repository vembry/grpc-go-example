package main

import (
	"context"
	"grpc-go-example/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

// MAIN starts
const (
	port = "localhost:3000"
)

func main() {

	//initializing server for listening
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on %v, %v", port, err)
	}

	//initializing grpc server
	grpcServer := grpc.NewServer()

	//registering our grpc
	proto.RegisterChatServiceServer(grpcServer, &Server{})

	//serving grpc
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server port %v, %v", port, err)
	}

}

// MAIN end

//GRPC definition starts
type Server struct {
	proto.UnimplementedChatServiceServer
}

//our rpc command
func (s *Server) SayHello(ctx context.Context, message *proto.Message) (*proto.Message, error) {
	log.Printf("incoming: %s", message.Body)
	return &proto.Message{Body: "Message received"}, nil
}

//GRPC definition ends
