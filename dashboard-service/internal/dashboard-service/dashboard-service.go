package dashboardservice

import (
	"dashboard-service/genproto/dashboardpb"
	"dashboard-service/internal/pkg/config"
	"dashboard-service/internal/service"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type RunGRPC struct {
	serv *service.Service
}

func NewRunGRPC(serv *service.Service) *RunGRPC {
	return &RunGRPC{
		serv: serv,
	}
}

func (s *RunGRPC) RUN(cfg *config.Config) error {
	target := fmt.Sprintf("%s:%d", cfg.DashboardServiceHost, cfg.DashboardServicePort)
	listener, err := net.Listen("tcp", target)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	dashboardpb.RegisterDashboardServiceServer(server, s.serv)
	return server.Serve(listener)
}
