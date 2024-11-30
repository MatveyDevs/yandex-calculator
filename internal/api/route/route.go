package route

import (
	"github.com/MatveyDevs/yandex-calculator/internal/api/controllers"
	mw "github.com/MatveyDevs/yandex-calculator/internal/middleware"
	"github.com/MatveyDevs/yandex-calculator/internal/service"
	"net/http"
)

func NewCalculationRoute(router *http.ServeMux) {
	s := service.New()
	c := controllers.New(s)
	router.HandleFunc("/calc", mw.LoggingMiddleware(c.GetCalc))
}
