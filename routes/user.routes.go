package routes

import (
	"github.com/gin-gonic/gin"
	"mef.world/backend/controllers"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewUserRouteController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (rc *UserRouteController) UserRoute(rg *gin.RouterGroup) {
	router := rg.Group("user")
	router.GET("/:id/avatar", rc.userController.GetUserAvatar)
}
