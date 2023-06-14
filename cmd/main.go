package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/dundee/grpc-dark-side/proto"
	"github.com/dundee/grpc-dark-side/proto/nested"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 5000, "server port")
)

type server struct {
	pb.UnimplementedSomeServiceServer
}

func (s *server) GetSome(ctx context.Context, in *pb.Some) (*pb.Some, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.Some{
		Name: "xxx",
		Nested: &nested.Nested{
			Name: "yyy",
		},
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSomeServiceServer(s, &server{})
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
