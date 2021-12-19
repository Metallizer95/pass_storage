package monitoring

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AliveController(handler *gin.Engine) {
	handler.GET("/alive", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
}
