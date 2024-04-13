package update

import (
	"context"
	"fmt"

	"github.com/pradytpk/go-mongodb-atlas/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindReplace(collection *mongo.Collection, client *mongo.Client) {

	filter := bson.D{{"studentid", 1}}

	replaceDoc := models.Student{
		StudentID: 2,
		Name:      "Test2",
		Age:       20,
		Gender:    "Female",
		Grade:     "F",
		Marks: []map[string]int{
			{"eng": 89},
			{"lang": 89},
			{"maths": 100},
		},
	}
	opts := options.FindOneAndReplace().SetReturnDocument(options.After)
	var student models.Student
	err := collection.FindOneAndReplace(context.TODO(), filter, replaceDoc, opts).Decode(&student)
	if err != nil {
		panic(err)
	}
	result, _ := bson.MarshalExtJSON(student, false, false)

	fmt.Println(string(result))

}
