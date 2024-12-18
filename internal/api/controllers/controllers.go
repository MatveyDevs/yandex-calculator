package controllers

import (
	"encoding/json"
	"github.com/MatveyDevs/yandex-calculator/internal/models"
	"net/http"
)

type CalculationService interface {
	CalculateExpression(data models.CalculationResponse) (float64, error)
}

type CalculationController struct {
	service CalculationService
}

func New(service CalculationService) *CalculationController {
	return &CalculationController{
		service: service,
	}
}

func (c *CalculationController) GetCalc(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var data models.CalculationResponse

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(internalServerErr)
		return
	}
	if data.Expression == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(internalServerErr)
		return
	}

	calc, err := c.service.CalculateExpression(data)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(invalidExpressionErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	resp := map[string]float64{"result": calc}
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(internalServerErr)
		return
	}
}
