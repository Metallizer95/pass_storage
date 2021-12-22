package middleware

import (
	"github.com/gin-gonic/gin"
)

func ApplyMiddleware(handler *gin.Engine) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
}
