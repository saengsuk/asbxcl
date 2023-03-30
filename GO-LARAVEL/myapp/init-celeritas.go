package main

import (
	"github.com/saengsuk/celeritas"
	"log"
	"os"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init Celeritas
	cel := &celeritas.Celeritas{}
	cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "MYAPP"
	cel.Debug = true

	app := &application{App: cel}
	return app
}
