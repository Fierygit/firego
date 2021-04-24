package user

import (
	"encoding/json"
	"firego/src/common/kv/client"
	mid "firego/src/common/middleware"
	"firego/src/common/util"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	DB client.Leveldb
}

func NewUserController() UserController {
	db := client.NewConnector().SetSize(2).Connect(client.PRE_USER, "123456")
	return UserController{DB: db}
}

const (
	max_age = 60 * 60 * 24 * 7
)

func genAndSetToken(c *gin.Context, user_id string) {
	customClaims := &mid.CustomClaims{
		UserId: user_id, //用户id
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(max_age) * time.Second).Unix(), // 过期时间，必须设置
		},
	}
	//采用HMAC SHA256加密算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenString, err := token.SignedString([]byte(mid.SECRETKEY))
	if util.CheckAndResponseError(err, c) {
		return
	}

	c.SetCookie("token", tokenString, max_age, "", "firego.cn", true, true)

	c.JSON(http.StatusOK, gin.H{
		"msg":   "success",
		"token": tokenString,
	})
}

func (ctl *UserController) Login(c *gin.Context) {
	type LoginReq struct {
		Name string `form:"name" json:"name" binding:"required"`
	}
	req := &LoginReq{}
	err := c.BindJSON(&req)
	if util.CheckAndResponseError(err, c) {
		return
	}

	if req.Name == "" {
		logrus.Error("name can not be empty")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "name can not be empty"})
	}

	hasBeen := HasUser(ctl.DB, req.Name)
	uid := ""

	// 用户不存在
	if !hasBeen {
		// 添加新用户
		uid = util.GetSnowflake().Base36()
		user, err := AddUser(ctl.DB, uid, req.Name)
		if util.CheckAndResponseError(err, c) {
			return
		}

		logrus.Info("make a new user ", user.Name)
	} else { // 用户已存在
		user := &UserModel{}
		payload := ctl.DB.Get(kv_user_key, req.Name)

		err := json.Unmarshal([]byte(payload), user)
		if util.CheckAndResponseError(err, c) {
			return
		}

		uid = user.Uid
	}

	//生成token
	genAndSetToken(c, uid)
}
