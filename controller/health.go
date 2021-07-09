package controller

import (
	loggerObj "HELLO-GO/utility/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheck struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

var (
	logger = loggerObj.GetLogger()
)

func GetHealthStatus(ctx *gin.Context) {
	header := ctx.Request.Header
	logger.Error(header)
	finalResponse := HealthCheck{
		Message: "OK",
		Status:  200,
	}
	ctx.JSON(http.StatusOK, finalResponse)
}

func EmpDetails() string {
	return "ashish"
}
