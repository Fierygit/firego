package todolist

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Run() {
	initConfig()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/todo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	port := viper.GetString("server.port")
	if port != "" {
		r.Run(":" + port)
	}
}

func initConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/todolist/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
