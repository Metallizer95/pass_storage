package main

import (
	"fmt"
	"os"
	"store_server/internal/app"
	"store_server/pkg/logging"
)

// @title Passports and Routes of railways store server
// @version 1.0
// @description Store server for passports and routes.
// @host localhost:80

func main() {
	//init logger
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0665)
	if err != nil {
		fmt.Println(err)
	}
	logging.New(true, logFile)
	if err != nil {
		panic(err)
	}
	logger, _ := logging.GetLogger()
	logger.Info("Application was started")
	app.Run()
}

// TODO: Maybe need return only status code for each request? And reformat all function for return error
