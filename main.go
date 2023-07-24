package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mef.world/backend/controllers"
	"mef.world/backend/helpers"
	"mef.world/backend/models"
	"mef.world/backend/routes"
)

var (
	DB     *gorm.DB
	server *gin.Engine

	// Auth
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	// User
	UserController      controllers.UserController
	UserRouteController routes.UserRouteController
)

func init() {
	password := helpers.GetEnvVariable("DATABASE_PASSWORD")
	username := helpers.GetEnvVariable("DATABASE_USERNAME")
	host := helpers.GetEnvVariable("DATABASE_HOSTNAME")
	port := helpers.GetEnvVariable("DATABASE_HOSTPORT")
	dbname := helpers.GetEnvVariable("DATABASE_DATANAME")
	mode := helpers.GetEnvVariable("GIN_MODE")

	gin.SetMode(mode)

	dsn := fmt.Sprintf("host=%v dbname=%v user=%v password=%v port=%v sslmode=disable", host, dbname, username, password, port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("🚀 Connected Successfully to the Database")

	AuthController = controllers.NewAuthController(DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(DB)
	UserRouteController = routes.NewUserRouteController(UserController)
	server = gin.Default()

}

func main() {
	port := helpers.GetEnvVariable("PORT")

	router := server.Group("/api")
	router.GET("/status", func(ctx *gin.Context) {
		message := "Welcome to mef.world"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	log.Fatal(server.Run("localhost:" + port))
}
