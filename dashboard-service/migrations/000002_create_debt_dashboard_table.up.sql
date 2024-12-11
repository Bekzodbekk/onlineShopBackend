CREATE TABLE IF NOT EXISTS debt_dashboard(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    debt_id TEXT NOT NULL,
    debt_price INT NOT NULL,
    paid_price INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
)