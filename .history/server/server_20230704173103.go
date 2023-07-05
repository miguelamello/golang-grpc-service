package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	ppb "go_grpc_service/go_grpc_service"
)

type server struct{}

func (s *server) GRPCService(req *ppb.Request, stream ppb.GRPCServiceServer) error {
	for {
		// Receive the client's message
		fmt.Printf("Received message: %s\n", req.Message)

		// Construct the response message
		response := &ppb.Response{
			Message: "Response: " + req.Message,
		}

		// Send the response message back to the client
		if _, err := stream.SendFeedback(response) {
			if err != nil {
			return err
		}
		}

		// Receive the next ping message from the client
		ping, err := stream.Recv()
		if err != nil {
			return err
		}
	}
}

func main() {
	// Create a TCP listener on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the PingPong service with the server
	ppb.RegisterGRPCServiceServer(s, &server{})

	// Start accepting incoming requests
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
