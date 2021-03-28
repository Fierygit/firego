/*
 * @Author: Firefly
 * @Date: 2020-10-16 23:13:04
 * @Descripttion:
 * @LastEditTime: 2020-11-17 11:02:35
 */
package ddl

import (
	mid "firego/src/common/middleware"
	"firego/src/home/ddl/controller"

	"github.com/gin-gonic/gin"
)

//CollectRoute routers
func CollectRoute(r *gin.Engine) *gin.Engine {

	r.Use(mid.CORSMiddleware(), mid.RecoveryMiddleware())

	r.GET("ddl/hello", controller.Hello)

	ddlController := controller.NewDdlController()

	ddlUserRoutes := r.Group("/ddl/user")
	ddlUserRoutes.Use(mid.AuthMiddleware())
	ddlUserRoutes.POST("/add", ddlController.AddUser)
	ddlUserRoutes.DELETE(":id", ddlController.Delete)

	// ddlRoutes := r.Group("/ddl")
	// ddlRoutes.Use(mid.AuthMiddleware())
	// ddlRoutes.POST("", ddlController.Create)
	// ddlRoutes.GET(":id", ddlController.Retrieve)
	// ddlRoutes.PUT(":id", ddlController.Update)
	// ddlRoutes.DELETE(":id", ddlController.Delete)
	return r
}
