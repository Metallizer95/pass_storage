package app

import (
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"store_server/internal/controller/middleware"
	"store_server/internal/controller/monitoring"
	passportctrl "store_server/internal/controller/passport"
	routescontroller "store_server/internal/controller/routes"
	passport2 "store_server/internal/domain/passport"
	routersentity "store_server/internal/domain/routers"
	routestorage "store_server/internal/storage/inmem/router"
	"store_server/internal/storage/mongorepo"
	"store_server/internal/usecase/passport"
	"store_server/internal/usecase/routers"
	"store_server/pkg/httpserver"
	"store_server/pkg/logging"
	"syscall"
)

func Run() {
	handler := gin.New()

	// Create mongo client
	mongourl := os.Getenv("MONGO_URL")
	mongoport := os.Getenv("MONGO_PORT")
	mongourl = "mongodb://" + mongourl + ":" + mongoport
	logger, _ := logging.GetLogger()
	logger.Infof("MONGO_URL: %s", mongourl)
	repoClient, err := mongorepo.NewClient(&mongorepo.Config{Path: mongourl})
	if err != nil {
		panic(err)
	}

	// Create passports repository, manager and use-cases
	passportManager := passport2.NewPassportManager(repoClient.PassportRepository())
	passportUseCases := passport.NewUseCases(passportManager)

	routeStore := routestorage.New()
	routeManager := routersentity.NewRouteManager(routeStore)
	routeUseCase := routers.NewUseCases(routeManager, passportUseCases)

	//Routing of handler
	middleware.ApplyMiddleware(handler)
	passportctrl.NewPassportHandlers(handler, passportUseCases)
	routescontroller.NewRoutesHandlers(handler, routeUseCase)
	monitoring.AliveController(handler)

	server := httpserver.New(handler, httpserver.Option(httpserver.Port(os.Getenv("APP_PORT"))))
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Infof("server is interrupted: %v", s)
	case err = <-server.Notify():
		logger.Errorf("server error: %v", err)
	}

	err = server.Shutdown()
	if err != nil {
		logger.Infof("server shutdown: %v", err)
	}
}
