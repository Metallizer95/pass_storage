package main

import (
	"io"
	"os"
	"store_server/internal/app"
	"store_server/pkg/logging"
)

func main() {
	//init logger
	_, err := logging.New(logging.INFO, []io.Writer{os.Stdout})
	if err != nil {
		panic(err)
	}
	app.Run()

}

// TODO: Inject logger
// TODO: add error types instead nel returning
// TODO: create structure of error on usecase layer
// TODO: Remove extra pointers
