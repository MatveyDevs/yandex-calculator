package controllers_test

import (
	"github.com/MatveyDevs/yandex-calculator/internal/api/controllers"
	"github.com/MatveyDevs/yandex-calculator/internal/service"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetCalcHandlerSuccessCase(t *testing.T) {
	s := service.New()
	c := controllers.New(s)
	expected := strings.TrimSpace(`{"result":4}`)
	req := httptest.NewRequest(http.MethodPost, "/calc", strings.NewReader(`{"expression": "2+2"}`))
	w := httptest.NewRecorder()
	c.GetCalc(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if strings.TrimSpace(string(data)) != expected {
		t.Errorf("Expected %s but got %s", expected, string(data))
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("wrong status code")
	}
}
