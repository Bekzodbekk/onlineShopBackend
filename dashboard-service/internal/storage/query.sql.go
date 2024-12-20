// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package storage

import (
	"context"
)

const getAddDebtCount = `-- name: GetAddDebtCount :one
SELECT COUNT(id) 
FROM debt_dashboard
WHERE paid_price = 0 AND created_at::DATE = CURRENT_DATE
`

func (q *Queries) GetAddDebtCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getAddDebtCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getBodyPrice = `-- name: GetBodyPrice :one
SELECT SUM(body_price) 
FROM product_dashboard 
WHERE created_at::DATE = CURRENT_DATE
`

func (q *Queries) GetBodyPrice(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getBodyPrice)
	var sum int64
	err := row.Scan(&sum)
	return sum, err
}

const getGivenDebtCount = `-- name: GetGivenDebtCount :one
SELECT 
    COUNT(DISTINCT id) AS unique_id_count
FROM 
    debt_dashboard
WHERE 
    paid_debt != 0 AND created_at::DATE = CURRENT_DATE
`

func (q *Queries) GetGivenDebtCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getGivenDebtCount)
	var unique_id_count int64
	err := row.Scan(&unique_id_count)
	return unique_id_count, err
}

const getPaidDebt = `-- name: GetPaidDebt :one
SELECT SUM(paid_price) AS total_paid_price
FROM debt_dashboard
WHERE paid_price != 0 AND created_at::DATE = CURRENT_DATE
`

func (q *Queries) GetPaidDebt(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getPaidDebt)
	var total_paid_price int64
	err := row.Scan(&total_paid_price)
	return total_paid_price, err
}

const getProfit = `-- name: GetProfit :one
SELECT SUM(paid_price - price) AS total_difference
FROM product_dashboard 
WHERE created_at::DATE = CURRENT_DATE
`

func (q *Queries) GetProfit(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getProfit)
	var total_difference int64
	err := row.Scan(&total_difference)
	return total_difference, err
}

const getSoldCount = `-- name: GetSoldCount :one
SELECT COUNT(id) 
FROM product_dashboard 
WHERE created_at::DATE = CURRENT_DATE
`

func (q *Queries) GetSoldCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getSoldCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getTotalClient = `-- name: GetTotalClient :one
SELECT COUNT(id) AS total_clients
FROM (
    SELECT DISTINCT id FROM debt_dashboard 
    WHERE paid_price = 0 AND created_at::DATE = CURRENT_DATE
    UNION ALL
    SELECT id FROM product_dashboard 
    WHERE created_at::DATE = CURRENT_DATE
) AS allClients
`

func (q *Queries) GetTotalClient(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getTotalClient)
	var total_clients int64
	err := row.Scan(&total_clients)
	return total_clients, err
}

const getTotalSum = `-- name: GetTotalSum :one
SELECT SUM(paid_price) AS total_paid_price
FROM (
    SELECT paid_price FROM product_dashboard WHERE created_at::DATE = CURRENT_DATE
    UNION ALL
    SELECT paid_price FROM debt_dashboard WHERE created_at::DATE = CURRENT_DATE
) AS combined
`

func (q *Queries) GetTotalSum(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getTotalSum)
	var total_paid_price int64
	err := row.Scan(&total_paid_price)
	return total_paid_price, err
}

const insertAddDebtOrPayment = `-- name: InsertAddDebtOrPayment :exec
INSERT INTO debt_dashboard(
    debt_id,
    debt_price,
    paid_price
) VALUES ($1, $2, $3)
`

type InsertAddDebtOrPaymentParams struct {
	DebtID    string
	DebtPrice int32
	PaidPrice int32
}

func (q *Queries) InsertAddDebtOrPayment(ctx context.Context, arg InsertAddDebtOrPaymentParams) error {
	_, err := q.db.ExecContext(ctx, insertAddDebtOrPayment, arg.DebtID, arg.DebtPrice, arg.PaidPrice)
	return err
}

const insertSellProduct = `-- name: InsertSellProduct :exec
INSERT INTO product_dashboard(
    product_id,
    color,
    price,
    paid_price
) VALUES ($1, $2, $3, $4)
`

type InsertSellProductParams struct {
	ProductID string
	Color     string
	Price     int32
	PaidPrice int32
}

func (q *Queries) InsertSellProduct(ctx context.Context, arg InsertSellProductParams) error {
	_, err := q.db.ExecContext(ctx, insertSellProduct,
		arg.ProductID,
		arg.Color,
		arg.Price,
		arg.PaidPrice,
	)
	return err
}
