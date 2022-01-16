package monitoring

import (
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"       // swagger embed files
	_ "github.com/swaggo/gin-swagger" // gin-swagger middleware
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	_ "store_server/docs"
)

// @Summary Alive
// @Tags monitoring
// @Description check server status
func AliveController(handler *gin.Engine) {
	handler.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
