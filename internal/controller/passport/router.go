package passportctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"store_server/internal/usecase/passport"
)

type passportRoutes struct {
	saveUseCase passport.SavePassportUseCase
	loadUseCase passport.LoadPassportUseCase
}

func NewRouter(handler *gin.Engine, saveUseCase passport.SavePassportUseCase, LoadUseCase passport.LoadPassportUseCase) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.GET("/alive", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	r := passportRoutes{saveUseCase: saveUseCase, loadUseCase: LoadUseCase}
	{
		handler.POST("passport", r.SavePassport)
		handler.GET("passport/:id", r.LoadPassport)
	}
}

func (pr *passportRoutes) SavePassport(c *gin.Context) {
	var request passport.Model
	err := c.ShouldBindXML(&request)
	if err != nil {
		fmt.Printf("[save passport]: error occurred: %v", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	pr.saveUseCase.Save(request)
}

func (pr *passportRoutes) LoadPassport(c *gin.Context) {

}
