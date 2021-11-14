package controller

import (
	"firego/app/ddl/dto"
	"firego/comm/response"
	"firego/server/kv/client"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//IDdlController 这个controller的接口
type IDdlController interface {
	RestController
	AddUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

//DdlController 还可以用 service 封装
type DdlController struct {
	DB client.Leveldb
}

//NewDdlController A
func NewDdlController() IDdlController {
	db := client.NewConnector().SetSize(2).Connect(client.PRE_DDL, "123456")
	return DdlController{DB: db}
}

func (controller DdlController) GetImg(ctx *gin.Context) {

}

//AddUser c
func (controller DdlController) AddUser(ctx *gin.Context) {

	user, _ := ctx.Get("user")
	logrus.Info(user)

	var ddlUser dto.DdlUser
	// 数据验证
	if err := ctx.BindJSON(&ddlUser); err != nil {
		logrus.Info(err.Error())
		response.Error(ctx, "数据验证错误", nil)
		return
	}

	logrus.Info(ddlUser)
	// if controller.DB.Has(dto.UserPrifix + ddlUser.Name) {
	// 	response.Error(ctx, "用户已经存在", nil)
	// } else {
	// 	controller.DB.Put1(dto.UserPrifix, ddlUser.Name, "1")
	// 	response.Success(ctx, nil, "增加成功")
	// }

}

//DeleteUser c
func (controller DdlController) DeleteUser(ctx *gin.Context) {

}

//Create c
func (controller DdlController) Create(ctx *gin.Context) {

}

//Retrieve r
func (controller DdlController) Retrieve(ctx *gin.Context) {

}

//Update u
func (controller DdlController) Update(ctx *gin.Context) {

}

//Delete d
func (controller DdlController) Delete(ctx *gin.Context) {

}
