package handler

import (
	"fmt"
	"learn-golang/utils"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type NewsHandler struct{}

type PostNewsV1Param struct {
	Title  string `form:"title" binding:"required"`
	Status string `form:"status" binding:"required,oneof=1 2"`
}

func NewNewsHandler() *NewsHandler {
	return &NewsHandler{}
}

func (n *NewsHandler) PostNewsV1(ctx *gin.Context) {
	var param PostNewsV1Param
	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	if file.Size > 5<<20 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "File size large (5MB)"})
		return
	}

	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create directory"})
		return
	}

	dst := fmt.Sprintf("./uploads/%s", filepath.Base(file.Filename))
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot save file"})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Post news (V1)",
		"title":   param.Title,
		"status":  param.Status,
		"image":   dst,
	})
}

func (n *NewsHandler) PostUploadFileNewsV1(ctx *gin.Context) {
	var param PostNewsV1Param
	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	filename, err := utils.ValidateAndSaveFile(file, "./uploads")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Post news (V1)",
		"title":   param.Title,
		"status":  param.Status,
		"image":   filename,
		"path":    "./uploads/" + filename,
	})
}

func (n *NewsHandler) PostUploadMultipleFileNewsV1(ctx *gin.Context) {
	const publicURL = "http://localhost:8080/images/"
	var param PostNewsV1Param
	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.HandleValidationErrors(err))
		return
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid multipart form"})
		return
	}

	images := form.File["images"]
	if len(images) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No file provided"})
		return
	}

	var successFiles []string
	var fieldFile []map[string]string
	for _, image := range images {
		filename, err := utils.ValidateAndSaveFile(image, "./uploads")
		if err != nil {
			fieldFile = append(fieldFile, map[string]string{
				"filename": image.Filename,
				"error":    err.Error(),
			})

			continue
		}

		publicImageURL := publicURL + filename
		successFiles = append(successFiles, publicImageURL)
	}

	resp := gin.H{
		"message":       "Post news (V1)",
		"title":         param.Title,
		"status":        param.Status,
		"success_files": successFiles,
	}

	if len(fieldFile) > 0 {
		resp["message"] = "Upload complete with partial errors"
		resp["error_files"] = fieldFile
	}

	ctx.JSON(200, resp)
}
