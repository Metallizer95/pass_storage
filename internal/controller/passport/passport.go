package passportctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"store_server/internal/usecase/passport"
	"store_server/pkg/logging"
)

type controller struct {
	SaveUseCase      passport.SavePassportUseCase
	LoadUseCase      passport.LoadPassportUseCase
	GetTowersUseCase passport.GetTowersUseCase
	logger           logging.Logger
}

func NewPassportHandlers(handler *gin.Engine, uc passport.UseCases) {
	r := controller{
		SaveUseCase:      uc.SavePassportUseCase(),
		LoadUseCase:      uc.LoadPassportUseCase(),
		GetTowersUseCase: uc.GetTowersUseCase(),
	}
	gr := handler.Group("/passport")
	{
		gr.POST("/", r.SavePassport)
		gr.GET("/:id", r.LoadPassport)
		gr.GET("/:id/towers", r.PassportTowers)
	}
}

func (ctrl *controller) SavePassport(c *gin.Context) {
	var request passport.Model
	ctrl.logger.Info("\nget request to save passport")

	err := c.ShouldBindXML(&request)
	if err != nil {
		ctrl.logger.Error(fmt.Sprintf("Could not parse request: %v", err))
		c.XML(http.StatusBadRequest, nil)
		return
	}

	pass := ctrl.SaveUseCase.Save(request)
	c.XML(http.StatusOK, pass)
	ctrl.logger.Info("return statusOk")
}

func (ctrl *controller) LoadPassport(c *gin.Context) {
	passportId := c.Params.ByName("id")
	ctrl.logger.Infof("\nget request to load passport with id: %s", passportId)

	p := ctrl.LoadUseCase.Load(passportId)
	if p == nil {
		ctrl.logger.Warnf("there are not passport with id %s in database", passportId)
		c.XML(http.StatusInternalServerError, nil)
		return
	}
	c.XML(http.StatusOK, p)
	ctrl.logger.Info("return statusOk")
}

func (ctrl *controller) PassportTowers(c *gin.Context) {
	passportId := c.Params.ByName("id")
	ctrl.logger.Infof("\nget request to get towers of passport with id %s", passportId)
	towers := ctrl.GetTowersUseCase.LoadTowers(passportId)
	if towers == nil {
		ctrl.logger.Infof("there is not passport with id %s in database", passportId)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, towers)
	ctrl.logger.Info("return statusOK")
}
