package commonctrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"store_server/internal/usecase/common"
)

type controller struct {
	useCase common.UseCases
}

func NewCommonHandler(handler *gin.Engine, cases common.UseCases) {
	r := controller{useCase: cases}

	group := handler.Group("/common")
	{
		group.GET("/passports/:id", r.GetPassportsByRoute)
	}
}

func (ctrl *controller) GetPassportsByRoute(c *gin.Context) {
	routeid := c.Params.ByName("id")
	response := ctrl.useCase.LoadPassportsByRouteUseCase().Load(routeid)

	c.XML(http.StatusOK, response)
}
