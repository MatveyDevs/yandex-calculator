package middleware

import (
	"encoding/json"
	"github.com/MatveyDevs/yandex-calculator/internal/models"
	"log"
	"net/http"
)

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body models.CalculationResponse
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Printf("Received request: %s %s\nBody: %s", r.Method, r.URL.Path, body)
		next(w, r)
	}
}
