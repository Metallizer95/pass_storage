package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"store_server/internal/controller/middleware"
	"store_server/internal/controller/monitoring"
	passportctrl "store_server/internal/controller/passport"
	routescontroller "store_server/internal/controller/routes"
	routers2 "store_server/internal/domain/routers"
	passportstorage "store_server/internal/storage/passport"
	routestorage "store_server/internal/storage/router"
	"store_server/internal/usecase/passport"
	"store_server/internal/usecase/routers"
	"store_server/pkg/httpserver"
	"syscall"
)

func Run() {
	handler := gin.New()

	// Create repoPassport
	passportStore := passportstorage.New()
	routeStore := routestorage.New()
	// Create manager here
	//TODO Create use-cases
	passportUseCases := passport.NewUseCases(passportStore)
	routeManager := routers2.NewRouteManager(routeStore)
	routeUseCase := routers.NewUseCases(routeManager, passportUseCases)

	//Routing of handler
	middleware.ApplyMiddleware(handler)
	passportctrl.NewPassportHandlers(handler, passportUseCases.SavePassport(), passportUseCases.LoadPassport())
	routescontroller.NewRoutesHandlers(handler, routeUseCase)
	monitoring.AliveController(handler)

	server := httpserver.New(handler)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

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
