package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/mudrikah/go-chi-playground/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedPersonServiceServer
}

func (s *Server) GetPerson(ctx context.Context, req *pb.GetPersonRequest) (*pb.Person, error) {
	// This is a simple implementation that returns a hardcoded person
	// In a real application, you would typically fetch this from a database
	return &pb.Person{
		Name: req.Name,
		Age:  30,
	}, nil
}

func StartGRPCServer(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPersonServiceServer(s, &Server{})

	log.Printf("Starting gRPC server on :%d", port)
	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}
