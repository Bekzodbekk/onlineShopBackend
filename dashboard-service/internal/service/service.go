package service

import (
	"context"
	"dashboard-service/genproto/dashboardpb"
	"dashboard-service/internal/repository"
)

type Service struct {
	dashboardpb.UnimplementedDashboardServiceServer
	repo repository.IDashboardRepository
}

func NewService(repo repository.IDashboardRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetDashboardInfo(ctx context.Context, req *dashboardpb.GetDashboardInfoReq) (*dashboardpb.GetDashboardInfoResp, error) {
	return s.repo.GetDashboardInfo(ctx, req)
}
