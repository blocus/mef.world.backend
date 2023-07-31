package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mef.world/backend/models"
)

type HabitController struct {
	DB *gorm.DB
}

func NewHabitController(DB *gorm.DB) HabitController {
	return HabitController{DB}
}

func (hc *HabitController) GetUserCurrentHabits(ctx *gin.Context) {
	var habits []models.Habit
	userId, exist := ctx.Get("user_id")
	var limit = 7

	response := []models.HabitActivityResponse{}

	if !exist {
		ctx.AbortWithStatus(http.StatusForbidden)
		return
	}

	if result := hc.DB.Find(&habits, "user_id = ? AND active = 'true'", userId); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	for _, h := range habits {
		tmp := models.HabitActivityResponse{}
		tmp.ID = h.ID
		tmp.Title = h.Title
		tmp.StartDate = h.StartDate
		tmp.Active = h.Active
		tmp.Activity = []models.ActivityResponse{}
		var activities []models.HabitActivity

		date := ctx.Query("date")

		var currentTime time.Time
		if date != "" {
			parsedDate, errorParse := time.Parse("2006-01-02", date)
			if errorParse != nil {
				ctx.AbortWithStatus(http.StatusNotFound)
				return
			}
			currentTime = parsedDate
		} else {
			currentTime = time.Now()
		}

		if intVar, err := strconv.Atoi(ctx.Query("limit")); err == nil {
			limit = intVar
		}

		result := hc.DB.Order("date desc").Find(&activities, "habit_id = ? AND date > ? AND date <= ?", h.ID, currentTime.AddDate(0, 0, -7), currentTime)
		if result.Error != nil {
			ctx.AbortWithError(http.StatusNotFound, result.Error)
			break
		}

		for day := 0; day < limit; day++ {
			dayToCheck := currentTime.AddDate(0, 0, -day)
			index := -1
			for i, act := range activities {
				if act.Date.Day() == dayToCheck.Day() && act.Date.Month() == dayToCheck.Month() && act.Date.Year() == dayToCheck.Year() {
					index = i
					break
				}
			}
			var activity models.ActivityResponse

			if index == -1 {
				activity.Done = false
				activity.Date = &dayToCheck
			} else {
				activity.ID = activities[index].ID
				activity.Done = activities[index].Done
				activity.Date = activities[index].Date
			}
			tmp.Activity = append(tmp.Activity, activity)
		}

		response = append(response, tmp)
	}

	ctx.JSON(http.StatusOK, &response)

}
