package web

import (
	"github.com/gin-gonic/gin"
)

func LoadAPI(engi *gin.Engine, handlers ...gin.HandlerFunc) {
	route := engi.Group("/api/web/v1", handlers...)

	open := route.Group("/open")
	open.GET("/ping", Ping)
}
