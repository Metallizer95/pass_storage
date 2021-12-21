package main

import "store_server/internal/app"

func main() {
	app.Run()
}

// TODO: Inject logger
// TODO: add error types instead nel returning
// TODO: create structure of error on usecase layer
