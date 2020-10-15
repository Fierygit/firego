package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Hello hello
func Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"test": "hello world"})
}
