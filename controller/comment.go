package controller

import (
	"HELLO-GO/config"
	"HELLO-GO/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllComments(ctx *gin.Context) {
	var header = make(map[string]string)
	header["Content-Type"] = "application/json"
	header["Accept"] = "application/json"
	url := config.HttpConfigProperty.CommentAPIURL
	commentResponse := service.GetCommentResponse(header, url)
	ctx.JSON(http.StatusOK, &commentResponse)
}
