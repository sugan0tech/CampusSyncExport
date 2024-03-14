package routes

import (
	eventmanager "CampusSyncExport/event-manager"
	"CampusSyncExport/mongo"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestPayload struct {
	ClassCode string `json:"class_code"`
}

func export(context *gin.Context) {
  requestPayload := RequestPayload{}

	if err := context.ShouldBindJSON(&requestPayload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
  fmt.Println(requestPayload.ClassCode)
	students := mongo.GetStudents(requestPayload.ClassCode)
	dates := eventmanager.GetAcademicDates()

	file_name := eventmanager.Export(students, dates, requestPayload.ClassCode)

	context.File("./" + file_name)
}
