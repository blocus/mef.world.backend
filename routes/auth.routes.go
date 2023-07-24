package routes

import (
	"github.com/gin-gonic/gin"
	"mef.world/backend/controllers"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("auth")

	router.POST("/login", rc.authController.LoginUser)
	router.GET("/current", rc.authController.GetCurrentUser)
}
