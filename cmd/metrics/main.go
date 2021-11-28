package main

import (
	"github.com/MaxFando/go-observability/internal/domain1"
	domain1 "github.com/MaxFando/go-observability/internal/domain2"
	"log"
)

func main() {
	app1 := domain1.App{}
	app2 := domain2.App{}

	if err := app1.Init(); err != nil {
		log.Fatal(err)
	}

	if err := app1.Serve(); err != nil {
		log.Fatal(err)
	}

	if err := app2.Init(); err != nil {
		log.Fatal(err)
	}

	if err := app2.Serve(); err != nil {
		log.Fatal(err)
	}
}
