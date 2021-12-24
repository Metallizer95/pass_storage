package routescontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"store_server/internal/usecase/errs"
	"store_server/internal/usecase/routers"
	"store_server/pkg/logging"
)

type controller struct {
	useCases routers.UseCases
	logger   *logging.Logger
}

func NewRoutesHandlers(handler *gin.Engine, cases routers.UseCases) {
	logger, _ := logging.GetLogger()
	r := controller{useCases: cases, logger: logger}

	gr := handler.Group("route")
	{
		gr.POST("/", r.Save)
		gr.GET("/:id", r.LoadByID)
		gr.GET("/all", r.LoadAll)
		gr.GET("/:id/passports", r.GetPassportsByRoute)
	}
}

func (ctrl *controller) Save(c *gin.Context) {
	var body routers.RouteModel
	ctrl.logger.Info("get query to save route")
	err := c.ShouldBindXML(&body)
	if err != nil {
		ctrl.logger.Error(err)
		c.XML(http.StatusBadRequest, errs.NewErrModel(err))
		return
	}

	result := ctrl.useCases.SaveRouter().Save(body)
	if result == nil {
		errResponse := errs.NewErrModel(errs.ErrObjectAlreadyExists)
		c.XML(http.StatusOK, errResponse)
		return
	}
	c.XML(http.StatusOK, result)
	ctrl.logger.Info("return statusOk")
}

func (ctrl *controller) LoadByID(c *gin.Context) {
	id := c.Params.ByName("id")
	ctrl.logger.Infof("\nget query load route with id: %s", id)
	result := ctrl.useCases.LoadRouterByID().Load(id)

	if result == nil {
		ctrl.logger.Error(fmt.Sprintf("not exist route with id %s", id))
		errResponse := errs.NewErrModel(errs.ErrObjectNotFound)
		c.XML(http.StatusBadRequest, errResponse)
		return
	}
	c.XML(http.StatusOK, result)
	ctrl.logger.Info("return statusOk")
}

func (ctrl *controller) LoadAll(c *gin.Context) {
	result := ctrl.useCases.LoadRouters().Load()
	ctrl.logger.Infof("get query load all")
	if result == nil {
		ctrl.logger.Warn("there are not routes in database")
		errResponse := errs.NewErrModel(errs.ErrNotFoundRoutes)
		c.XML(http.StatusOK, errResponse)
		return
	}
	c.XML(http.StatusOK, result)
	ctrl.logger.Info("return statusOk")
}

func (ctrl *controller) GetPassportsByRoute(c *gin.Context) {
	routeid := c.Params.ByName("id")
	ctrl.logger.Infof("\nget query load passports by route with id %s", routeid)

	result := ctrl.useCases.LoadPassportsByRoute().Load(routeid)
	if result == nil {
		ctrl.logger.Warnf("there is not route with id %s", routeid)
		errResponse := errs.NewErrModel(errs.ErrObjectNotFound)
		c.XML(http.StatusBadRequest, errResponse)
		return
	}
	c.XML(http.StatusOK, result)
	ctrl.logger.Info("return statusOk")
}
