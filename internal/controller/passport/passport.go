package passportctrl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"store_server/internal/usecase/passport"
)

type Controller struct {
	SaveUseCase passport.SavePassportUseCase
	LoadUseCase passport.LoadPassportUseCase
}

// TODO: use UseCases instead use each use case separately

func NewPassportHandlers(handler *gin.Engine, saveUseCase passport.SavePassportUseCase, LoadUseCase passport.LoadPassportUseCase) {
	r := Controller{SaveUseCase: saveUseCase, LoadUseCase: LoadUseCase}
	gr := handler.Group("/passport")
	{
		gr.POST("/", r.SavePassport)
		gr.GET("/:id", r.LoadPassport)
		gr.GET("/:id/towers", r.PassportTowers)
	}
}

func (ctrl *Controller) SavePassport(c *gin.Context) {

	var request passport.Model
	err := c.ShouldBindXML(&request)
	if err != nil {
		fmt.Printf("[save passport]: error occurred: %v", err)
		c.XML(http.StatusBadRequest, nil)
		return
	}

	pass := ctrl.SaveUseCase.Save(request)
	c.XML(http.StatusOK, pass)
}

func (ctrl *Controller) LoadPassport(c *gin.Context) {
	p := ctrl.LoadUseCase.Load(c.Params.ByName("id"))
	if p == nil {
		fmt.Println("[load passport]: return nil pointer of passport")
		c.XML(http.StatusInternalServerError, nil)
		return
	}
	c.XML(http.StatusOK, p)
}

func (ctrl *Controller) PassportTowers(c *gin.Context) {

}
