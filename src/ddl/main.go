/*
 * @Author: Firefly
 * @Date: 2020-10-15 22:41:28
 * @Descripttion:
 * @LastEditTime: 2020-10-15 23:20:55
 */
package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func initConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("ddl")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

//CollectRoute router
func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.POST("/api/hello", controller.Hello)
	return r
}
