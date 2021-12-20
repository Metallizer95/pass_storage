package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"store_server/internal/controller/monitoring"
	passportctrl "store_server/internal/controller/passport"
	storage "store_server/internal/storage/passport"
	"store_server/internal/usecase/passport"
	"store_server/pkg/httpserver"
	"syscall"
)

func Run() {
	handler := gin.New()

	// Create repository
	repository := storage.New()
	// Create use-cases
	useCases := passport.NewUseCases(repository)

	//Routing of handler. Maybe need to create new interface for this
	passportctrl.NewPassportHandlers(handler, useCases.SavePassport(), useCases.LoadPassport())
	monitoring.AliveController(handler)

	server := httpserver.New(handler)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// TODO: add logger
	var err error
	select {
	case s := <-interrupt:
		fmt.Println(s)
	case err = <-server.Notify():
		fmt.Printf("server error: %v", err)
	}

	err = server.Shutdown()
	if err != nil {
		fmt.Printf("server shutdown error: %v", err)
	}
}
