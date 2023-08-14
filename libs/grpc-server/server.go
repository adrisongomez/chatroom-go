package grpcserver

import (
	config_parser "config-parser"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Server is a wrapper on top of grpc.Server and net.Listener
// to make it easier to start and stop a grpc server.
//
// It also all the plugins needed for the grpc server.
type Server struct {
	config     config_parser.ServiceConfig
	grpcServer *grpc.Server
}

func (s *Server) GetRef() *grpc.Server {
	return s.grpcServer
}

func (s Server) Listen() {
	port := s.config.ServicePort
	if port == 0 {
		log.Printf("service: %s port is not set, using default port 5000", s.config.ServiceName)
		port = 5000
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatalf("service: %s failed to listen in port %d: %v", s.config.ServiceName, s.config.ServicePort, err)
		panic(err)
	}

	if s.grpcServer == nil {
		log.Fatalf("service: %s grpc server is not initialized", s.config.ServiceName)
		panic(err)
	}

	log.Printf("service: %s listening in port %d", s.config.ServiceName, s.config.ServicePort)
	if err := s.grpcServer.Serve(lis); err != nil {
		log.Fatalf("service: %s failed to serve: %v", s.config.ServiceName, err)
		panic(err)
	}
}

func (s Server) Stop() {
	log.Printf("service: %s stopping", s.config.ServiceName)
	s.grpcServer.GracefulStop()
	log.Printf("service: %s stopped", s.config.ServiceName)
}

func NewServer(config *config_parser.Config) Server {
	grpcServer := grpc.NewServer()

	return Server{
		config:     config.Service,
		grpcServer: grpcServer,
	}
}
