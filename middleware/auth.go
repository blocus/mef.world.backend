package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"mef.world/backend/helpers"
)

type MyCustomClaims struct {
	Username  string `json:"username"`
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	jwt.RegisteredClaims
}

func VerifyAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hmacSecret := helpers.GetEnvVariable("JSON_WEB_TOKEN_HMAC_SECRET")
		authorisation := strings.Split(ctx.GetHeader("authorisation"), " ")

		if authorisation[0] != "Bearer" || len(authorisation) != 2 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		token := authorisation[1]
		res, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(hmacSecret), nil
		})

		if claims, ok := res.Claims.(*MyCustomClaims); ok && res.Valid && err == nil {
			ctx.Set("user_id", claims.Id)
		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Next()
	}
}
