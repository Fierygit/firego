/*
 * @Author: Firefly
 * @Date: 2020-10-16 23:33:44
 * @Descripttion:
 * @LastEditTime: 2020-10-17 15:31:20
 */
package dto

//UserPrifix UserPrifix
var UserPrifix = "user"

//DdlUser 不需要名字
type DdlUser struct {
	Name string `from:"name" json:"name" binding:"required"`
}
