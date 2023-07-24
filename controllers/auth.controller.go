package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/xlzd/gotp"
	"gorm.io/gorm"
	"mef.world/backend/models"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

type JWTStructure struct {
	Jwt string `json:"jwt"`
}

type LoginResponse struct {
	JWT  JWTStructure                 `json:"jwt"`
	User models.UserResponseStructure `json:"user"`
}

var hmacSampleSecret = []byte{12, 45, 65, 12, 56, 23, 56, 43, 54, 83, 23, 76, 34, 67, 23, 56, 56, 43}

func (ac *AuthController) LoginUser(ctx *gin.Context) {
	var payload *models.OTPInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	message := "Token is invalid or user doesn't exist"

	var user models.User
	result := ac.DB.First(&user, "username = ?", "meftah")
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "message"})
		return
	}

	totp := gotp.NewDefaultTOTP(user.Otp_secret)
	println(totp.Now())

	now := time.Now() 
	sec := now.Unix()

	valid := totp.Verify(payload.Token, sec)
	if !valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": message})
		return
	}

	userResponse := models.UserResponseStructure{
		ID:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        user.ID.String(),
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"username":  user.Username,
	})
	jwToken, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	response := LoginResponse{
		JWT: JWTStructure{
			Jwt: jwToken,
		},
		User: userResponse,
	}
	ctx.IndentedJSON(http.StatusOK, response)
}

func (ac *AuthController) GetCurrentUser(ctx *gin.Context) {
	var user models.User

	result := ac.DB.First(&user, "username = ?", "meftah")
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password", "error": result.Error})
		return
	}

	userResponse := models.UserResponseStructure{
		ID:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        user.ID.String(),
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"username":  user.Username,
	})
	jwToken, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	response := LoginResponse{
		JWT: JWTStructure{
			Jwt: jwToken,
		},
		User: userResponse,
	}
	ctx.IndentedJSON(http.StatusOK, response)
}

