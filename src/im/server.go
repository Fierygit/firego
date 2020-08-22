



func run(){
	logrus.Info("start to run the im server!!!")
	
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		wechat := c.PostForm("wechat")
		c.String(200, wechat)
	})

	go r.Run(":8080")
}