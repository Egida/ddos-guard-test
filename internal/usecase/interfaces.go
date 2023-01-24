package usecase

import (
	"github.com/Shteyd/ddos-guard-test/internal/entity"
)

type (
	User interface {
		Metric() (entity.Metric, error)
		Store(username string) error
		GetUserID(username string) (int, error)
	}

	Math interface {
		Calculate(text string) (float64, error)
	}

	UserRepo interface {
		GetUserID(username string) (int, error)
		Store(username string) error
		GetMetric() (entity.Metric, error)
	}

	MathService interface {
		Calculate(text string) (float64, error)
	}
)
