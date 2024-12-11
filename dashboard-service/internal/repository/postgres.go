package repository

import (
	"context"
	"dashboard-service/genproto/dashboardpb"
	"dashboard-service/internal/storage"
	"strconv"
)

type DashboardREPO struct {
	queries *storage.Queries
}

func NewForConsumerDashboardRepo(queries *storage.Queries) *DashboardREPO {
	return &DashboardREPO{
		queries: queries,
	}
}
func NewDashboardREPO(queries *storage.Queries) IDashboardRepository {
	return &DashboardREPO{
		queries: queries,
	}
}

func (r *DashboardREPO) GetDashboardInfo(ctx context.Context, req *dashboardpb.GetDashboardInfoReq) (*dashboardpb.GetDashboardInfoResp, error) {
	totalSum, err := r.queries.GetTotalSum(ctx)
	if err != nil {
		return nil, err
	}
	sold_out, err := r.queries.GetSoldCount(ctx)
	if err != nil {
		return nil, err
	}
	add_debt, err := r.queries.GetAddDebtCount(ctx)
	if err != nil {
		return nil, err
	}
	given_debt, err := r.queries.GetGivenDebtCount(ctx)
	if err != nil {
		return nil, err
	}
	total_client, err := r.queries.GetTotalClient(ctx)
	if err != nil {
		return nil, err
	}
	profit, err := r.queries.GetProfit(ctx)
	if err != nil {
		return nil, err
	}
	body_price, err := r.queries.GetBodyPrice(ctx)
	if err != nil {
		return nil, err
	}
	paid_debt, err := r.queries.GetPaidDebt(ctx)
	if err != nil {
		return nil, err
	}

	return &dashboardpb.GetDashboardInfoResp{
		Status:  true,
		Message: "Get Info In Dashboard Success",
		Dashboard: &dashboardpb.Dashboard{
			Total:       strconv.FormatInt(totalSum, 10),
			SoldOut:     int32(sold_out),
			AddDebt:     int32(add_debt),
			GivenDebt:   int32(given_debt),
			TotalClient: int32(total_client),
			Profit:      profit,
			BodyPrice:   body_price,
			PaidDebt:    paid_debt,
		},
	}, nil
}

func (r *DashboardREPO) SellProductInsertDashboard(ctx context.Context, req *ProductDashboard) error {
	err := r.queries.InsertSellProduct(ctx, storage.InsertSellProductParams{
		ProductID: req.ProductID,
		Color:     req.Color,
		Price:     int32(req.Price),
		PaidPrice: int32(req.PaidPrice),
	})
	if err != nil {
		return err
	}
	return nil
}
func (r *DashboardREPO) AddDebtInsertDashboard(ctx context.Context, req *DebtDashboard) error {
	err := r.queries.InsertAddDebtOrPayment(ctx, storage.InsertAddDebtOrPaymentParams{
		DebtID:    req.DebtID,
		DebtPrice: int32(req.DebtPrice),
		PaidPrice: int32(req.PaidPrice),
	})
	if err != nil {
		return err
	}
	return nil
}
