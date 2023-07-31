package routes

import (
	"github.com/gin-gonic/gin"
	"mef.world/backend/controllers"
	"mef.world/backend/middleware"
)

type HabitRouteController struct {
	userController controllers.HabitController
}

func NewHabitRouteController(userController controllers.HabitController) HabitRouteController {
	return HabitRouteController{userController}
}

func (rc *HabitRouteController) HabitRoute(rg *gin.RouterGroup) {
	router := rg.Group("habit")
	router.GET("/", middleware.VerifyAuth(), rc.userController.GetUserCurrentHabits)
}
