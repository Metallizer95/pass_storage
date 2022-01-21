package passportctrl

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"store_server/internal/usecase/errs"
	"store_server/internal/usecase/passport"
	"store_server/pkg/logging"
	"strconv"

	_ "github.com/swaggo/files"       // swagger embed files
	_ "github.com/swaggo/gin-swagger" // gin-swagger middleware
	_ "store_server/docs"
)

const (
	xmlTypeHeader = "application/xml"
	zipTypeHeader = "multipart/form-data"
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
		gr.POST("/", r.save)
		gr.GET("/:passportId", r.loadPassport)
		gr.GET("/:passportId/towers", r.passportTowers)
		gr.GET("/:passportId/towers/findtower", r.findTowerByIdAndCoordinate)
		gr.GET("/:passportId/towers/:towerId", r.getPassportTowerById)
		gr.GET("/expired", r.findExpiredPassports)
	}
}

// @Summary SavePassport
// @Tags passports
// @Description save passport in database
// @Param input body passport.PassportModel true "xml structure of passport or zip archive of xml files"
// @Success 200 {object} passport.PassportModel
// @Failure 400 {object} errs.ErrorModel
// @Router /passport [post]
func (ctrl *controller) save(c *gin.Context) {
	var request passport.Model
	ctrl.logger.Info("get request to save passport")

	contentType := c.ContentType()
	fmt.Println(contentType)
	switch contentType {
	case xmlTypeHeader:
		err := c.ShouldBindXML(&request)
		if err != nil {
			ctrl.logger.Error(fmt.Sprintf("Could not parse request: %v", err))
			c.XML(http.StatusBadRequest, errs.NewErrModel(err))
			return
		}
		ctrl.savePassport(request, c)

	case zipTypeHeader:
		files, err := ctrl.uploadFile(c)
		if err != nil {
			c.XML(http.StatusBadRequest, errs.NewErrModel(err))
			return
		}
		var models []passport.Model
		for _, f := range files {
			fil, err := readZipFile(f)
			if err != nil {
				fmt.Printf("error reading: %v", err)
			}
			err = xml.Unmarshal(fil, &request)
			models = append(models, request)
		}
		err = ctrl.SaveUseCase.SaveMany(models)
		if err != nil {
			c.XML(http.StatusBadRequest, errs.NewErrModel(err))
		}
		c.XML(http.StatusOK, models)
	}
}

// @Summary GetPassportByID
// @Tags passports
// @Description return passport by ID from database if there is one, or return error object with status code 200
// @Param passportId path string true "passport ID"
// @Success 200 {object} passport.PassportModel
// @Router /:passportId [get]
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

// @Summary GetTowersOfPassport
// @Tags towers
// @Description return all towers of passport by id
// @Param passportId path string true "passport ID"
// @Success 200 {object} passport.TowersModel
// @Router /:passportId/towers [get]
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

// @Summary GetPassportTowerByID
// @Tags towers
// @Description return certain tower of the passport by ID
// @Param passportId path string true "passport ID"
// @Param towerId path string true "tower ID"
// @Success 200 {object} passport.TowerModel
// @Router /:passportId/towers/:towerId [get]
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

// @Summary FindTowerByCoordinate
// @Tags towers
// @Description return the closest tower belonged the passport by coordinates
// @Param latitude query number true "latitude"
// @Param longitude query number true "longitude"
// @Success 200 {object} passport.TowerModel
// @Failure 400 {object} errs.ErrorModel
// @Router /:passportId/towers/findtower [get]
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
