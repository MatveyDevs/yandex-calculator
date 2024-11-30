package main

import (
	"github.com/MatveyDevs/yandex-calculator/internal/application"
	"log"
)

func main() {
	app := application.New()
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
