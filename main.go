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
)

func init() {
	password := goDotEnvVariable("DATABASE_PASSWORD")
	username := goDotEnvVariable("DATABASE_USERNAME")
	host := goDotEnvVariable("DATABASE_HOSTNAME")
	port := goDotEnvVariable("DATABASE_HOSTPORT")
	dbname := goDotEnvVariable("DATABASE_DATANAME")

	dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("🚀 Connected Successfully to the Database")

	server = gin.Default()

}

func main() {
	mode := goDotEnvVariable("GIN_MODE")
	port := goDotEnvVariable("PORT")

	gin.SetMode(mode)

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Two-Factor Authentication with Golang"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	log.Fatal(server.Run("localhost:" + port))
}
