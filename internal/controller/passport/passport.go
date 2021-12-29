package passportctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"store_server/internal/usecase/errs"
	"store_server/internal/usecase/passport"
	"store_server/pkg/logging"
	"strconv"
)

type controller struct {
	SaveUseCase      passport.SavePassportUseCase
	LoadUseCase      passport.LoadPassportUseCase
	GetTowersUseCase passport.GetTowersUseCase
	FindTower        passport.FindTowerByIdAndCoordinateUseCase
	ExpiredPassports passport.FindExpiredPassportsUseCase
	logger           *logging.Logger
}

func NewPassportHandlers(handler *gin.Engine, uc passport.UseCases) {
	logger, _ := logging.GetLogger()
	r := controller{
		SaveUseCase:      uc.SavePassportUseCase(),
		LoadUseCase:      uc.LoadPassportUseCase(),
		GetTowersUseCase: uc.GetTowersUseCase(),
		FindTower:        uc.FindTowerByIdAndCoordinateUseCase(),
		ExpiredPassports: uc.FindExpiredPassportsUseCase(),
		logger:           logger,
	}
	gr := handler.Group("/passport")
	{
		gr.POST("/", r.savePassport)
		gr.GET("/:passportId", r.loadPassport)
		gr.GET("/:passportId/towers", r.passportTowers)
		gr.GET("/:passportId/towers/findtower", r.findTowerByIdAndCoordinate)
		gr.GET("/:passportId/towers/:towerId", r.getPassportTowerById)
		gr.GET("/expired", r.findExpiredPassports)
	}
}

func (ctrl *controller) savePassport(c *gin.Context) {
	var request passport.Model
	ctrl.logger.Info("get request to save passport")

	err := c.ShouldBindXML(&request)
	if err != nil {
		ctrl.logger.Error(fmt.Sprintf("Could not parse request: %v", err))
		c.XML(http.StatusBadRequest, errs.NewErrModel(err))
		return
	}

	pass := ctrl.SaveUseCase.Save(request)
	if pass == nil {
		errResponse := errs.NewErrModel(errs.ErrObjectAlreadyExists)
		c.XML(http.StatusOK, errResponse)
		return
	}
	c.XML(http.StatusOK, pass)
	ctrl.logger.Info("return statusOk")
}

func (ctrl *controller) loadPassport(c *gin.Context) {
	passportId := c.Params.ByName("passportId")
	ctrl.logger.Infof("get request to load passport with id: %s", passportId)

	p := ctrl.LoadUseCase.Load(passportId)
	if p == nil {
		ctrl.logger.Warnf("there are not passport with id %s in database", passportId)
		errResponse := errs.NewErrModel(errs.ErrObjectNotFound)
		c.XML(http.StatusOK, errResponse)
		return
	}
	c.XML(http.StatusOK, p)
	ctrl.logger.Info("return status 200")
}

func (ctrl *controller) passportTowers(c *gin.Context) {
	passportId := c.Params.ByName("passportId")
	ctrl.logger.Infof("get request to get towers of passport with id %s", passportId)

	towers := ctrl.GetTowersUseCase.LoadAllTowerByPassportId(passportId)
	if towers == nil {
		ctrl.logger.Infof("there is not passport with id %s in database", passportId)
		errResponse := errs.NewErrModel(errs.ErrObjectNotFound)
		c.JSON(http.StatusOK, errResponse)
		return
	}

	c.JSON(http.StatusOK, towers)
	ctrl.logger.Info("return status 200")
}

func (ctrl *controller) getPassportTowerById(c *gin.Context) {
	passportId := c.Params.ByName("passportId")
	towerId := c.Params.ByName("towerId")

	ctrl.logger.Infof("request to get tower of passport; passportId: %s, towerId:%s", passportId, towerId)

	result := ctrl.GetTowersUseCase.LoadTowerById(passportId, towerId)
	if result == nil {
		c.XML(http.StatusOK, errs.NewErrModel(errs.ErrObjectNotFound))
		ctrl.logger.Warn("not found object")
		return
	}

	c.XML(http.StatusOK, result)
	ctrl.logger.Info("return status 200")
}

func (ctrl *controller) findTowerByIdAndCoordinate(c *gin.Context) {
	values := c.Request.URL.Query()
	ctrl.logger.Info("get request to find tower by coordinate")
	longitudeString, ok := values["longitude"]
	if !ok {
		ctrl.logger.Warn("user did not type longitude parameter")
		c.XML(http.StatusBadRequest, errs.NewErrModel(ErrNotFindLongitude))
		return
	}

	longitude, err := strconv.ParseFloat(longitudeString[0], 64)
	if err != nil {
		ctrl.logger.Error(err)
		c.XML(http.StatusBadRequest, errs.NewErrModel(ErrWrongTypeLongitudeParameter))
		return
	}

	latitudeString, ok := values["latitude"]
	if !ok {
		ctrl.logger.Info("user did not type latitude parameter")
		c.XML(http.StatusBadRequest, errs.NewErrModel(ErrNotFindLatitude))
		return
	}

	latitude, err := strconv.ParseFloat(latitudeString[0], 64)
	if err != nil {
		ctrl.logger.Error(err)
		c.XML(http.StatusBadRequest, errs.NewErrModel(ErrWrongTypeLatitudeParameter))
		return
	}

	passportId := c.Params.ByName("passportId")
	tower := ctrl.FindTower.FindTower(passportId, longitude, latitude)
	if tower == nil {
		c.XML(http.StatusOK, errs.NewErrModel(errs.ErrObjectNotFound))
		return
	}

	c.XML(http.StatusOK, tower)
	ctrl.logger.Info("return status 200")
}

func (ctrl *controller) findExpiredPassports(c *gin.Context) {
	c.XML(http.StatusOK, ctrl.ExpiredPassports.FindPassports())
}
