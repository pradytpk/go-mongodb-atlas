package update

import (
	"context"
	"fmt"

	"github.com/pradytpk/go-mongodb-atlas/connection"
	"github.com/pradytpk/go-mongodb-atlas/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Update() {
	client := connection.Connection()

	filter := bson.D{{Key: "studentid", Value: 1}}
	replaceDoc := models.Student{
		StudentID: 1,
		Name:      "Kumar",
		Age:       30,
		Gender:    "male",
		Grade:     "B",
		Marks: []map[string]int{
			{
				"eng": 99,
			},
			{
				"lang": 79,
			},
			{
				"math": 60,
			},
		},
	}
	collection := client.Database("university").Collection("students")

	result, err := collection.ReplaceOne(context.TODO(), filter, replaceDoc)
	if err != nil {
		panic(err)
	}
	if result.ModifiedCount != 0 {
		fmt.Println("No document modified:", result.ModifiedCount)
	}
}

func UpdateMany() {
	client := connection.Connection()

	filter := bson.D{{Key: "age", Value: 21}}

	update := bson.D{{"$set", bson.D{{"age", 22}}}}
	colllection := client.Database("university").Collection("students")
	result, err := colllection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Println("no of documents updated:", result.ModifiedCount)
}
