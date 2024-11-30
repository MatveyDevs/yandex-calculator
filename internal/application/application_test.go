package application_test

import (
	"github.com/MatveyDevs/yandex-calculator/internal/application"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApplication_Run(t *testing.T) {
	app := application.New()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	app.Run()

	resp, err := http.Get(ts.URL + "/invalid")
	if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("expected status %v, but got %v", http.StatusNotFound, resp.Status)
	}
}
