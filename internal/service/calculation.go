package service

import (
	"github.com/MatveyDevs/yandex-calculator/internal/models"
	"github.com/MatveyDevs/yandex-calculator/pkg/calculation"
)

type CalculationService struct{}

func New() *CalculationService {
	return &CalculationService{}
}

func (cs *CalculationService) CalculateExpression(data models.CalculationResponse) (float64, error) {
	return calculation.Calc(data.Expression)
}
