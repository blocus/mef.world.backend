package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

type currentHabitInterface struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	MON   bool
	TUE   bool
	WED   bool
	THU   bool
	FRI   bool
	SAT   bool
	SUN   bool
}

type JWTStructure struct {
	Jwt string `json:"jwt"`
}

type UserStructure struct {
	ID        string `json:"id"`
	Avatar    string `json:"avatar"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
}

type loginResponse struct {
	JWT  JWTStructure  `json:"jwt"`
	User UserStructure `json:"user"`
}

var habits = []currentHabitInterface{
	{ID: "4", Title: "teeth bruch every morning", MON: true, TUE: true, WED: true, THU: true, FRI: true, SAT: true, SUN: true},
	{ID: "1", Title: "Drink water", MON: true, TUE: true, WED: true, THU: true, FRI: true, SAT: true, SUN: true},
	{ID: "2", Title: "Read 10 pages from a book", MON: true, TUE: true, WED: true, THU: true, FRI: true, SAT: true, SUN: true},
	{ID: "3", Title: "workout for 1h", MON: true, TUE: true, WED: true, THU: true, FRI: true, SAT: true, SUN: true},
}

type getCurrentHabitInterface struct {
	Data []currentHabitInterface `json:"data"`
	Week int                     `json:"week"`
}

func getCurrentHabits(c *gin.Context) {

	var currentHabits = getCurrentHabitInterface{
		Data: habits,
		Week: 10,
	}

	c.IndentedJSON(http.StatusOK, currentHabits)
}

func postAlbums(c *gin.Context) {
	var newHabit currentHabitInterface

	if err := c.BindJSON(&newHabit); err != nil {
		return
	}

	habits = append(habits, newHabit)
	c.IndentedJSON(http.StatusCreated, newHabit)
}

var user = loginResponse{
	JWT: JWTStructure{
		Jwt: "7nCqwyKJ9kh3jWvDng61UgBMXck8yhEme2jqyuceAPCxJ8Ts7jOR0Dy0UVpKuvDx9w7xP90BuKHYMsVWD18ZtcjgWh03IUfAORHyqf7ivZ",
	},
	User: UserStructure{
		ID:        "f691ab0c-225d-11ee-be56-0242ac120002",
		Avatar:    "/api/user/meftah/avatar",
		FirstName: "Ahmed",
		LastName:  "Meftah",
		Username:  "meftah",
	},
}

func login(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, user)
}

func getUserAvatar(c *gin.Context) {
	c.File("./avatar/avatar.jpg")
}

func main() {
	mode := goDotEnvVariable("GIN_MODE")
	port := goDotEnvVariable("PORT")

	gin.SetMode(mode)

	router := gin.Default()
	router.GET("/api/habit/current", getCurrentHabits)
	router.POST("/api/auth/login", login)
	router.GET("/api/auth", login)
	router.GET("/api/user/:id/avatar", getUserAvatar)
	router.Run("localhost:" + port)
}
