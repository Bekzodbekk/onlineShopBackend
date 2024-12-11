package repository

type ProductDashboard struct {
	ProductID string `json:"product_id"`
	Color     string `json:"color"`
	Price     int    `json:"price"`
	PaidPrice int    `json:"paid_price"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type DebtDashboard struct {
	DebtID    string `json:"debt_id"`
	DebtPrice int    `json:"debt_price"`
	PaidPrice int    `json:"paid_price"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}
