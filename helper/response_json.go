package helper

import (
	"github.com/gin-gonic/gin"
)

func ErrJSON(ctx *gin.Context, httpStatus int, err string) {
	ctx.JSON(httpStatus, gin.H{
		"status":  "fail",
		"message": err,
	})
}
