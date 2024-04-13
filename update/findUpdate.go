package update

import (
	"context"
	"fmt"

	"github.com/pradytpk/go-mongodb-atlas/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindUpdate - compound operation
func FindUpdate(collection *mongo.Collection, client *mongo.Client) {

	filter := bson.D{{"studentid", 2}}

	update := bson.D{{"$push", bson.D{{"marks", bson.D{{"science", 89}}}}}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var student models.Student

	err := collection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&student)
	if err != nil {
		panic(err)
	}
	res, _ := bson.MarshalExtJSON(student, false, false)
	fmt.Println(string(res))
}
