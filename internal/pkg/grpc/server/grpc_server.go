package grpcserver

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"vladazn/wow/internal/service"
	"vladazn/wow/proto/gen/go/proto/wow"
)

type Server struct {
	port     int
	server   *grpc.Server
	services *service.Services
}

func NewGrpcServer(
	port int,
	services *service.Services,
) Server {

	return Server{
		port:     port,
		server:   grpc.NewServer(),
		services: services,
	}
}

func (s *Server) Serve() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", s.port))
	if err != nil {
		return err
	}

	wow.RegisterWowServer(s.server, NewWowServer(s.services))

	err = s.server.Serve(lis)

	if err != nil {
		return err
	}
	return nil
}

func (s *Server) Stop() {
	s.server.GracefulStop()
}
