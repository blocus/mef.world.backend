package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

func (ac *UserController) GetUserAvatar(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.File("./avatar/" + id + ".jpg")
}
