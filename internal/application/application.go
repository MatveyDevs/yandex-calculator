package application

import (
	"github.com/MatveyDevs/yandex-calculator/internal/api/route"
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
	route.NewCalculationRoute(mux)
	log.Println("Server started on port :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		return err
	}
	return nil
}
