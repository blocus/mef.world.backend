package routes

import (
	"github.com/gin-gonic/gin"
	"mef.world/backend/controllers"
	"mef.world/backend/middleware"
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
	router.GET("/current", middleware.VerifyAuth(), rc.authController.GetCurrentUser)
}
