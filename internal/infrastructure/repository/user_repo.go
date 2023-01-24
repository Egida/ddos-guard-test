package repository

import (
	"fmt"

	"github.com/Shteyd/ddos-guard-test/internal/entity"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	*sqlx.DB
}

func New(db *sqlx.DB) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) GetMetric() (entity.Metric, error) {
	var metric entity.Metric

	rows, err := r.Query("SELECT COUNT(id) FROM users")
	if err != nil {
		return metric, fmt.Errorf("UserRepo - GetMetric - r.Query: %w", err)
	}
	defer rows.Close()

	var count int

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return metric, fmt.Errorf("UserRepo - GetMetric - rows.Scan: %w", err)
		}
	}

	metric = entity.Metric{UserCount: count}

	return metric, nil
}

func (r *UserRepo) Store(username string) error {
	_, err := r.Exec("INSERT INTO (username) VALUES ($1)", username)
	if err != nil {
		return fmt.Errorf("UserRepo - Store - r.Exec: %w", err)
	}

	return nil
}

func (r *UserRepo) GetUserID(username string) (int, error) {
	var id int
	err := r.Get(&id, "SELECT id FROM users WHERE username = $1", username)
	if err != nil {
		return 0, fmt.Errorf("UserRepo - Check - r.Get: %w", err)
	}

	return id, nil
}
