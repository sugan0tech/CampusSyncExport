package routes

import (
	eventmanager "CampusSyncExport/event-manager"
	"CampusSyncExport/mongo"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RequestPayload struct {
	ClassCode string    `json:"classCode" binding:"required"`
	From      string `json:"from" time_format:"2006-01-02" binding:"required"`
	To        string `json:"to" time_format:"2006-01-02" binding:"required"`
}

func export(context *gin.Context) {
	requestPayload := RequestPayload{}

	if err := context.ShouldBindJSON(&requestPayload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(requestPayload.ClassCode)
	students := mongo.GetStudents(requestPayload.ClassCode)
	dates := filterDates(eventmanager.GetAcademicDates(), requestPayload.From, requestPayload.To)
  fmt.Println(len(eventmanager.GetAcademicDates()))
  fmt.Println(len(dates))

	file_name := eventmanager.Export(students, dates, requestPayload.ClassCode)

	context.File("./" + file_name)
}

func filterDates(dates []eventmanager.AcamedicDate, from, to string) []eventmanager.AcamedicDate {
	fromDate, err := time.Parse("2006-01-02", from)
	toDate, err := time.Parse("2006-01-02", to)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	filteredDates := make([]eventmanager.AcamedicDate, 0)
	for _, dateType := range dates {
		date, err := time.Parse("2006-01-02", dateType.Date)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			continue
		}

		if date.After(fromDate) && date.Before(toDate) {
			filteredDates = append(filteredDates, dateType)
		}
	}

	return filteredDates
}
