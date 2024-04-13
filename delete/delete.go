package delete

import (
	"context"
	"fmt"

	"github.com/pradytpk/go-mongodb-atlas/connection"
	"go.mongodb.org/mongo-driver/bson"
)

func Delete() {
	client := connection.Connection()

	filter := bson.D{{"studentid", bson.D{{"$gt", 1}}}}
	collection := client.Database("university").Collection("students")
	result, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Println("no of documents deleted:", result.DeletedCount)
}
