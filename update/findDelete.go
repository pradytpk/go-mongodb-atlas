package update

import (
	"context"
	"fmt"

	"github.com/pradytpk/go-mongodb-atlas/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindDelete(collection *mongo.Collection, client *mongo.Client) {
	filter := bson.D{{"studentid", 1}}
	var student models.Student

	err := collection.FindOneAndDelete(context.TODO(), filter).Decode(&student)
	if err != nil {
		panic(err)
	}
	result, _ := bson.MarshalExtJSON(student, false, false)
	fmt.Println(string(result))
}
