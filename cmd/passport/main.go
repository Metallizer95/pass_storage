package main

import (
	"fmt"
	"os"
	"store_server/internal/app"
	"store_server/pkg/logging"
)

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

// TODO: validation of values (Longitude and Latitude)