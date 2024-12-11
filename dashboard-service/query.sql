-- name: GetTotalSum :one
SELECT SUM(paid_price) AS total_paid_price
FROM (
    SELECT paid_price FROM product_dashboard WHERE created_at::DATE = CURRENT_DATE
    UNION ALL
    SELECT paid_price FROM debt_dashboard WHERE created_at::DATE = CURRENT_DATE
) AS combined;

-- name: GetSoldCount :one
SELECT COUNT(id) 
FROM product_dashboard 
WHERE created_at::DATE = CURRENT_DATE;


-- name: GetAddDebtCount :one
SELECT COUNT(id) 
FROM debt_dashboard
WHERE paid_price = 0 AND created_at::DATE = CURRENT_DATE;


-- name: GetGivenDebtCount :one
SELECT 
    COUNT(DISTINCT id) AS unique_id_count
FROM 
    debt_dashboard
WHERE 
    paid_debt != 0 AND created_at::DATE = CURRENT_DATE;



-- name: GetTotalClient :one
SELECT COUNT(id) AS total_clients
FROM (
    SELECT DISTINCT id FROM debt_dashboard 
    WHERE paid_price = 0 AND created_at::DATE = CURRENT_DATE
    UNION ALL
    SELECT id FROM product_dashboard 
    WHERE created_at::DATE = CURRENT_DATE
) AS allClients;


-- name: GetProfit :one
SELECT SUM(paid_price - price) AS total_difference
FROM product_dashboard 
WHERE created_at::DATE = CURRENT_DATE;

-- name: GetBodyPrice :one
SELECT SUM(body_price) 
FROM product_dashboard 
WHERE created_at::DATE = CURRENT_DATE;

-- name: GetPaidDebt :one
SELECT SUM(paid_price) AS total_paid_price
FROM debt_dashboard
WHERE paid_price != 0 AND created_at::DATE = CURRENT_DATE;


-- name: InsertSellProduct :exec
INSERT INTO product_dashboard(
    product_id,
    color,
    price,
    paid_price
) VALUES ($1, $2, $3, $4);

-- name: InsertAddDebtOrPayment :exec
INSERT INTO debt_dashboard(
    debt_id,
    debt_price,
    paid_price
) VALUES ($1, $2, $3);
