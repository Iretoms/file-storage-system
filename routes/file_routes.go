package routes

import "github.com/gin-gonic/gin"

func FileRoutes(r *gin.RouterGroup) {
	r.POST("/file")
}