package autosign

import (
	mid "firego/comm/middleware"
	// "encoding/json"
	res "firego/comm/response"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Run(port string) {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.Use(mid.CORSMiddleware())
	r.POST("/autosign/add", add)
	r.GET("/autosign/getall", getall)
	r.StaticFile("/autosign", "../static/index.html")
	r.StaticFile("/", "../static/index.html")

	r.Run(port)
}

func getall(c *gin.Context) {
	f, err := os.Open("./autosign/hnu/user")
	if err != nil {
		res.Error(c, err.Error(), nil)
	}
	user, err := ioutil.ReadAll(f)
	all_user := strings.Split(string(user), "\n")

	ret := ""
	for _, val := range all_user {
		if val == "" {
			continue
		}
		id := strings.Split(string(val), "=")[0]
		ret += id[0:3] + "!@#$%^" + id[9:] + "\n"
	}
	c.String(200, ret)
}

func add(c *gin.Context) {
	type Add struct {
		Id   string `form:"id" json:"id" binding:"required"`
		Name string `form:"pwd" json:"pwd" binding:"required"`
	}
	req := &Add{}
	err := c.BindJSON(&req)
	if err != nil {
		logrus.Error("bind json failed, err", err)
		res.Error(c, err.Error(), nil)
		return
	}
	command := `python3 /home/firefly/firego/src/autosign/hnu/cmdlogin.py ` + req.Id + ` ` + req.Name
	cmd := exec.Command("/bin/bash", "-c", command)

	cmdret, err := cmd.Output()
	if err != nil {
		logrus.Info("Execute Shell:{} failed with error:{}", command, err.Error())
		return
	}

	logrus.Info(string(cmdret))
	cmdLines := strings.Split(string(cmdret), "\n")
	var isSuccess = false
	for _, val := range cmdLines {
		if val == "success" {
			isSuccess = true
			break
		}
	}
	if isSuccess {
		f, err := os.Open("./autosign/hnu/user")
		if err != nil {
			res.Error(c, err.Error(), nil)
		}
		user, err := ioutil.ReadAll(f)
		all_user := strings.Split(string(user), "\n")
		for _, val := range all_user {
			if val == "" {
				continue
			}
			if strings.Split(string(val), "=")[0] == req.Id {
				res.Success(c, gin.H{"msg": "已经添加，不用重复添加！"}, string(cmdret))
				return
			}
		}
		output := string(user) + req.Id + "=" + req.Name + "\n"
		ioutil.WriteFile("./autosign/hnu/user", []byte(output), 0664)
		res.Success(c, gin.H{"msg": "添加成功，将为您定时执行打卡任务"}, string(cmdret))
		logrus.Info(all_user)
		return
	} else {
		res.Error(c, string(cmdret), gin.H{"msg": "添加失败，请查看返回信息"})
		return
	}

}
