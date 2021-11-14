package middleware

/*
 * @Author: Firefly
 * @Date: 2020-10-15 22:42:15
 * @Descripttion:
 * @LastEditTime: 2020-10-17 15:22:42
 */

import (
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//自定义Claims
type CustomClaims struct {
	UserId string `json:"uid"`
	jwt.StandardClaims
}

const (
	SECRETKEY = "243223ffslsfsldfl412fdsfsdf" //私钥
)

//AuthMiddleware a
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取authorization header
		tokenString, err := ctx.Cookie("token")

		if err != nil {
			logrus.Info("not exists token ", err)
			redirect(ctx)
			ctx.Abort()
			return
		}

		now := time.Now().Unix()

		// 解析jwt
		claims, err := parseToken(tokenString)
		if err != nil {
			logrus.Info("parse token fail", err)
			redirect(ctx)
			ctx.Abort()
			return
		}

		if claims.ExpiresAt < now {
			logrus.Info("jwt timeout ", now, claims.ExpiresAt)
			redirect(ctx)
			ctx.Abort()
		} else {
			logrus.Info("user_id: ", claims.UserId)
			ctx.Set("user_id", claims.UserId)
			ctx.Next()
		}

	}
}

func redirect(ctx *gin.Context) {
	originUrl := url.QueryEscape(ctx.Request.Referer())

	redirectUrl := "https://firego.cn/user/login.html?redirect=" + originUrl
	// redirectUrl := "http://localhost:9527/user/login.html?redirect=" + originUrl
	logrus.Info("redirect to ", redirectUrl)
	// ctx.Redirect(302, redirectUrl) // 这里不可以用302，因为axios不响应

	ctx.JSON(http.StatusUnauthorized, gin.H{
		"redirect": redirectUrl,
	})
}

func parseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logrus.Error("Unexpected signing method: ", token.Header["alg"])
			return nil, errors.New("unexpected signing method")
		}
		return []byte(SECRETKEY), nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}

}
