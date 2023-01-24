package mathservice

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

type MathService struct{}

func New() *MathService {
	return &MathService{}
}

func (ms *MathService) Calculate(text string) (float64, error) {
	expression, err := govaluate.NewEvaluableExpression(text)
	if err != nil {
		return 0, fmt.Errorf("MathService - Calculate - govaluate.NewEvaluableExpression: %w", err)
	}

	result, err := expression.Evaluate(nil)
	if err != nil {
		return 0, fmt.Errorf("MathService - Calculate - expression.Evaluate: %w", err)
	}

	value, ok := result.(float64)
	if !ok {
		return 0, fmt.Errorf("MathService - Calculate - result.(float64)")
	}

	return value, nil
}
