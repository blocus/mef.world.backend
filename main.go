package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mef.world/backend/controllers"
	"mef.world/backend/models"
	"mef.world/backend/routes"
)

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

var (
	DB     *gorm.DB
	server *gin.Engine
	// Auth
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController
)

func init() {
	password := goDotEnvVariable("DATABASE_PASSWORD")
	username := goDotEnvVariable("DATABASE_USERNAME")
	host := goDotEnvVariable("DATABASE_HOSTNAME")
	port := goDotEnvVariable("DATABASE_HOSTPORT")
	dbname := goDotEnvVariable("DATABASE_DATANAME")

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
	server = gin.Default()

}

func main() {
	mode := goDotEnvVariable("GIN_MODE")
	port := goDotEnvVariable("PORT")

	gin.SetMode(mode)

	router := server.Group("/api")
	router.GET("/status", func(ctx *gin.Context) {
		message := "Welcome to mef.world"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	log.Fatal(server.Run("localhost:" + port))
}
