package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"store_server/internal/usecase/routers"
)

type controller struct {
	useCases routers.UseCases
}

func NewRoutesHandlers(handler *gin.Engine, cases routers.UseCases) {
	r := controller{useCases: cases}

	gr := handler.Group("/route")
	{
		gr.POST("/save", r.Save)
		gr.GET("/load/:id", r.LoadByID)
		gr.GET("/load/all", r.LoadAll)
	}
}

func (ctrl *controller) Save(c *gin.Context) {
	var body routers.RouteModel

	err := c.ShouldBindXML(&body)
	if err != nil {
		c.XML(http.StatusBadRequest, nil)
		return
	}

	result := ctrl.useCases.SaveRouter().Save(body)
	if result == nil {
		c.XML(http.StatusInternalServerError, nil)
		return
	}
	c.XML(http.StatusOK, result)
}

func (ctrl *controller) LoadByID(c *gin.Context) {
	id := c.Params.ByName("id")
	result := ctrl.useCases.LoadRouterByID().Load(id)

	if result == nil {
		c.XML(http.StatusInternalServerError, nil)
		return
	}
	c.XML(http.StatusOK, result)
}

func (ctrl *controller) LoadAll(c *gin.Context) {
	result := ctrl.useCases.LoadRouters().Load()
	if result == nil {
		c.XML(http.StatusInternalServerError, nil)
		return
	}
	c.XML(http.StatusOK, result)
}
