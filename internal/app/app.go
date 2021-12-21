package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	commonctrl "store_server/internal/controller/common"
	"store_server/internal/controller/monitoring"
	passportctrl "store_server/internal/controller/passport"
	routescontroller "store_server/internal/controller/routes"
	commonstore "store_server/internal/storage/common"
	passportstorage "store_server/internal/storage/passport"
	routestorage "store_server/internal/storage/router"
	"store_server/internal/usecase/common"
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
	commonStore := commonstore.NewRepository(passportStore, routeStore)
	// Create use-cases
	passportUseCases := passport.NewUseCases(passportStore)
	routeUseCase := routers.NewUseCases(routeStore)
	commonUseCase := common.NewUseCases(commonStore)

	//Routing of handler
	passportctrl.NewPassportHandlers(handler, passportUseCases.SavePassport(), passportUseCases.LoadPassport())
	routescontroller.NewRoutesHandlers(handler, routeUseCase)
	commonctrl.NewCommonHandler(handler, commonUseCase)
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
