package main

import (
	"store_server/internal/app"
	"store_server/pkg/logging"
)

func main() {
	//init logger
	_, err := logging.New(logging.INFO, "l.log")
	if err != nil {
		panic(err)
	}
	log, err := logging.Get()
	if err != nil {
		panic(err)
	}
	log.Info("I initialized %s", "yeap")
	app.Run()

}

// TODO: Inject logger
// TODO: add error types instead nel returning
// TODO: create structure of error on usecase layer
// TODO: Remove extra pointers
