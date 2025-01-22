package application

import (
	"github.com/MatveyDevs/yandex-calculator/internal/api/route"
	"github.com/MatveyDevs/yandex-calculator/internal/config"
	"log"
	"net/http"
)

type Application struct {
}

func New() *Application {
	return &Application{}
}

func (a *Application) Run() error {
	mux := http.NewServeMux()
	cfg := config.New()
	route.NewCalculationRoute(mux)
	log.Printf("Server started on port :%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		return err
	}
	return nil
}
