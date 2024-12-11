package repository

import (
	"context"
	"dashboard-service/genproto/dashboardpb"
)

type IDashboardRepository interface {
	GetDashboardInfo(ctx context.Context, req *dashboardpb.GetDashboardInfoReq) (*dashboardpb.GetDashboardInfoResp, error)
}
