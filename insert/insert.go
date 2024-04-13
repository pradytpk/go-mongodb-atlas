package insert

import (
	"context"
	"fmt"

	"github.com/pradytpk/go-mongodb-atlas/connection"
	"github.com/pradytpk/go-mongodb-atlas/models"
)

func Insert() {
	client := connection.Connection()
	newStudent := models.Student{
		StudentID: 1,
		Name:      "Tes",
		Age:       20,
		Gender:    "Female",
		Grade:     "A",
		Marks: []map[string]int{
			{
				"eng": 89,
			},
			{
				"lang": 90,
			},
			{
				"math": 100,
			},
		},
	}
	collection := client.Database("university").Collection("students")
	result, err := collection.InsertOne(context.TODO(), newStudent)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	fmt.Println(result.InsertedID)

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}

func InsertMany() {
	client := connection.Connection()
	newStudent := []interface{}{
		models.Student{
			StudentID: 22,
			Name:      "Test23",
			Age:       34,
			Gender:    "Female",
			Grade:     "A",
			Marks: []map[string]int{
				{
					"eng": 90,
				},
				{
					"lang": 99,
				},
				{
					"math": 99,
				},
			},
		},
		models.Student{
			StudentID: 33,
			Name:      "Test34",
			Age:       45,
			Gender:    "Male",
			Grade:     "A",
			Marks: []map[string]int{
				{
					"eng": 90,
				},
				{
					"lang": 100,
				},
				{
					"math": 100,
				},
			},
		},
	}
	collection := client.Database("university").Collection("students")
	result, err := collection.InsertMany(context.TODO(), newStudent)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	for _, v := range result.InsertedIDs {
		fmt.Println(v)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}
