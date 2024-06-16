package routes

import (
	"file-storage-system/controller"

	"github.com/gin-gonic/gin"
)

func FileRoutes(r *gin.RouterGroup) {
	r.POST("/file", controller.UploadFile())
}