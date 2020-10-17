package ddl

/*
 * @Author: Firefly
 * @Date: 2020-10-15 22:41:28
 * @Descripttion:
 * @LastEditTime: 2020-10-16 23:15:49
 */

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//Run run
func Run() {
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
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/ddl/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}


