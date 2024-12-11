package postgres

import (
	"dashboard-service/internal/pkg/config"
	"dashboard-service/internal/storage"
	"database/sql"
	"fmt"
)

func InitDB(cfg *config.Config) (*sql.DB, error) {
	target := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Database, cfg.Postgres.User, cfg.Postgres.Password)

	db, err := sql.Open("postgres", target)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func Queries(db *sql.DB) *storage.Queries {
	return storage.New(db)
}
