package main

import (
	"CampusSyncExport/attendance"
	eventmanager "CampusSyncExport/event-manager"
	"CampusSyncExport/routes"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func main(){
  server := gin.Default();
  mongo, _ := attendance.ConnectToMongoDB()
  collection := mongo.Database("test").Collection("students")
  var result bson.M


  err := collection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: "20eucs147"}}).Decode(&result)

  dates := eventmanager.GetAcademicDates()
  fmt.Println(dates)

  if ( err != nil){
    panic(err)
  }

  fmt.Println(result["attendancePeriodSetMap"])
  routes.RegisterRoutes(server)
  server.Run(":9090");
}
