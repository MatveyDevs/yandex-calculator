package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/MatveyDevs/yandex-calculator/internal/models"
	"io"
	"log"
	"net/http"
)

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var body bytes.Buffer
		tee := io.TeeReader(r.Body, &body)
		var calcResp models.CalculationResponse
		if err := json.NewDecoder(tee).Decode(&calcResp); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("Received request: %s %s\nBody: %s", r.Method, r.URL.Path, calcResp)

		r.Body = io.NopCloser(&body)
		next(w, r)
	}
}
