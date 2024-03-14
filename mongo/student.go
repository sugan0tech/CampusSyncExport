package mongo

import (
	"CampusSyncExport/attendance"
	eventmanager "CampusSyncExport/event-manager"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetStudents(classCode string) []eventmanager.Student {
	mongo, _ := attendance.ConnectToMongoDB()
	collection := mongo.Database("test").Collection("students")

	projection := bson.M{"_id": 1, "classCode": 1, "attendancePeriodSetMap": 1}
	filter := bson.D{{ "classCode", classCode}}

	findOptions := options.Find()
	findOptions.SetLimit(64)
	findOptions.SetProjection(projection)

	cursor, err := collection.Find(context.TODO(), filter, findOptions)

	if err != nil {
		panic(err)
	}

	students := make([]eventmanager.Student, 0, 64)

	for cursor.Next(context.TODO()) {
		var student eventmanager.Student
		err := cursor.Decode(&student)
		if err != nil {
			panic(err)
		}
		students = append(students, student)
	}
  fmt.Println(students)

	defer cursor.Close(context.TODO())

	return students
}
