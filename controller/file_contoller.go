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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
			handleError(c, http.StatusBadRequest, err)
			return
		}
		defer formFile.Close()

		content, err := io.ReadAll(formFile)
		if err != nil {
			handleError(c, http.StatusInternalServerError, err)
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
			handleError(c, http.StatusBadRequest, err)
			return
		}

		if file.Size > 16777216 {
			c.JSON(http.StatusBadRequest, response.FileResponse{Status: http.StatusBadRequest, Message: "File too large, should be at most 16MB"})
			return
		}

		result, err := filesCollection.InsertOne(ctx, file)

		if err != nil {
			handleError(c, http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusCreated, response.FileResponse{Status: http.StatusCreated, Message: "success", FileReference: result.InsertedID})
	}
}

func DownloadFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		fileID := c.Param("fileID")
		defer cancel()

		var file model.File

		objID, _ := primitive.ObjectIDFromHex(fileID)

		err := filesCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&file)

		if err != nil {
			handleError(c, http.StatusNotFound, err)
			return
		}

		c.Header("Content-Disposition", "attachment; filename="+file.FileName)
		c.Header("Content-Type", file.ContentType)
		c.Data(http.StatusOK, file.ContentType, file.Content)

	}
}

func handleError(c *gin.Context, status int, e error) {
	c.JSON(status, response.FileResponse{Status: http.StatusBadRequest, Message: e.Error()})
}
