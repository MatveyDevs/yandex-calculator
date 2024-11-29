package application

import (
	"bufio"
	"github.com/MatveyDevs/yandex-calculator/pkg/rpn"
	"log"
	"os"
	"strings"
)

type Application struct {
}

func New() *Application {
	return &Application{}
}

func (a *Application) Run() error {
	for {
		log.Println("input expression")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println("failed to read expression from console")
		}

		text = strings.TrimSpace(text)

		if text == "exit" {
			log.Println("application was successfully closed")
			return nil
		}

		result, err := rpn.Calc(text)
		if err != nil {
			log.Println(text, " calculation failed with error: ", err)
		} else {
			log.Println(text, "=", result)
		}
	}
}
