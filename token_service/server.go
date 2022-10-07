package token_service

import (
	"fmt"
	"grpc1/util"

	pb "github.com/aydogduyusuf/grpc1/pb"
)

// Server serves gRPC requests for our service
type Server struct {
	pb.UnimplementedDBchainServer
	config     util.Config
}

// NewServer creates a new gRPC server
func NewServer(config util.Config) (*Server, error) {
	
	server := &Server{
		config:     config,
	}
	
	return server, nil
}