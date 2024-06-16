package controller

import (
	"context"
	"file-storage-system/database"
	"file-storage-system/model"
	"file-storage-system/response"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var filesCollection *mongo.Collection = database.GetCollection(database.DB, "files")
var validate = validator.New()

func UploadFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var file model.File

		formFile, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, response.FileResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer formFile.Close()

		content, err := io.ReadAll(formFile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.FileResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		file.FileName = header.Filename
		file.FileType = http.DetectContentType(content)
		file.Size = header.Size
		file.UploadDate = time.Now()
		file.LastModifiedDate = time.Now()
		file.Content = content
		file.ContentType = header.Header.Get("Content-Type")

		if validationErr := validate.Struct(&file); validationErr != nil {
			c.JSON(http.StatusBadRequest, response.FileResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		result, err := filesCollection.InsertOne(ctx, file)

		if err != nil {
			c.JSON(http.StatusInternalServerError, response.FileResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, response.FileResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}
