package handler

import (
	"fmt"
	"net/http"

	"github.com/gabrielbarretoo/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, gin.H{
		"message": message,
		"status":  code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s was successful", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type CreateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
