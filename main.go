package main

import (
	"github.com/pradytpk/go-mongodb-atlas/aggregation"
	"github.com/pradytpk/go-mongodb-atlas/connection"
)

func main() {
	client := connection.Connection()
	collection := client.Database("university").Collection("students")

	// insert.Insert()
	//insert.InsertMany()
	// update.UpdateMany()
	//update.UpdateMany()
	// query.Query()
	// delete.Delete()
	// update.FindReplace(collection, client)
	// update.FindUpdate(collection, client)
	// update.FindDelete(collection, client)
	aggregation.MatchGroupAggreation(collection, client)
	aggregation.SortProject(collection, client)
}
