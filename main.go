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
	"mef.world/backend/routes"
)

// type currentHabitInterface struct {
// 	ID    string `json:"id"`
// 	Title string `json:"title"`
// 	MON   bool
// 	TUE   bool
// 	WED   bool
// 	THU   bool
// 	FRI   bool
// 	SAT   bool
// 	SUN   bool
// }

// var habits = []currentHabitInterface{
// 	{ID: "4", Title: "teeth bruch every morning", MON: true, TUE: true, WED: true, THU: true, FRI: true, SAT: true, SUN: true},
// 	{ID: "1", Title: "Drink water", MON: true, TUE: true, WED: true, THU: true, FRI: true, SAT: true, SUN: true},
// 	{ID: "2", Title: "Read 10 pages from a book", MON: true, TUE: true, WED: true, THU: true, FRI: true, SAT: true, SUN: true},
// 	{ID: "3", Title: "workout for 1h", MON: true, TUE: true, WED: true, THU: true, FRI: true, SAT: true, SUN: true},
// }

// type getCurrentHabitInterface struct {
// 	Data []currentHabitInterface `json:"data"`
// 	Week int                     `json:"week"`
// }

// func getCurrentHabits(c *gin.Context) {

// 	var currentHabits = getCurrentHabitInterface{
// 		Data: habits,
// 		Week: 10,
// 	}

// 	c.IndentedJSON(http.StatusOK, currentHabits)
// }

// func postAlbums(c *gin.Context) {
// 	var newHabit currentHabitInterface

// 	if err := c.BindJSON(&newHabit); err != nil {
// 		return
// 	}

// 	habits = append(habits, newHabit)
// 	c.IndentedJSON(http.StatusCreated, newHabit)
// }

var (
	DB     *gorm.DB
	server *gin.Engine

	// Auth
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	// User
	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	// Habits
	HabitController      controllers.HabitController
	HabitRouteController routes.HabitRouteController
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
	// DB.AutoMigrate(&models.User{})
	// DB.AutoMigrate(&models.Habit{})
	// DB.AutoMigrate(&models.HabitActivity{})

	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("🚀 Connected Successfully to the Database")

	AuthController = controllers.NewAuthController(DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(DB)
	UserRouteController = routes.NewUserRouteController(UserController)

	HabitController = controllers.NewHabitController(DB)
	HabitRouteController = routes.NewHabitRouteController(HabitController)
	server = gin.Default()

}

func main() {
	port := helpers.GetEnvVariable("PORT")

	// router.GET("/api/habit/current", getCurrentHabits)
	// router.GET("/api/user/:id/avatar", getUserAvatar)

	// corsConfig := cors.DefaultConfig()
	// corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	// corsConfig.AllowCredentials = true

	router := server.Group("/api")
	router.GET("/status", func(ctx *gin.Context) {
		message := "Welcome to mef.world"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	HabitRouteController.HabitRoute(router)
	log.Fatal(server.Run("localhost:" + port))
}
