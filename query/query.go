package query

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pradytpk/go-mongodb-atlas/connection"
	"github.com/pradytpk/go-mongodb-atlas/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Query() {
	client := connection.Connection()

	filter := bson.D{{"grade", "B"}}
	collection := client.Database("university").Collection("students")
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []models.Student

	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, " ", "     ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", output)
	}
	defer func() {
		cursor.Close(context.TODO())
	}()
}
