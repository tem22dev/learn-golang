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
