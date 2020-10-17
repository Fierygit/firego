package controller

import "github.com/gin-gonic/gin"

//RestController 接口
type RestController interface {
	Create(ctx *gin.Context)
	Retrieve(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
