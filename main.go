package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/mralanlee/hello-world-protobuf/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) HelloWorld(ctx context.Context, in *pb.Greeting) (*pb.Hello, error) {
	return &pb.Hello{Message: "Hello World!"}, nil
}

func main() {
	// Create Server
	fmt.Printf("Listening on port %s", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
