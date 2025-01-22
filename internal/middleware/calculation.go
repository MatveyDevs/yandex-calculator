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
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Println(err)
			}
		}(r.Body)
		var body bytes.Buffer
		tee := io.TeeReader(r.Body, &body)
		var calcReq models.CalculationRequest
		if err := json.NewDecoder(tee).Decode(&calcReq); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("Received request: %s %s\nBody: %s", r.Method, r.URL.Path, calcReq)

		r.Body = io.NopCloser(&body)
		next(w, r)
	}
}
