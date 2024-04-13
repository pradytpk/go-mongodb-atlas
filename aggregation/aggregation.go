package aggregation

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func MatchGroupAggreation(collection *mongo.Collection, client *mongo.Client) {

	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "grade", Value: "A"}}}}

	groupStage := bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: "$gender"}, {Key: "count", Value: bson.D{{Key: "$count", Value: bson.D{}}}}}}}

	cursor, err := collection.Aggregate(context.TODO(), mongo.Pipeline{matchStage, groupStage})
	if err != nil {
		panic(err)
	}
	var results []bson.M

	err = cursor.All(context.TODO(), &results)

	for _, result := range results {
		fmt.Println("No of ", result["_id"], "students:", result["count"])
	}
}

func SortProject(collection *mongo.Collection, client *mongo.Client) {

	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "grade", Value: "A"}}}}
	sortStage := bson.D{{Key: "$sort", Value: bson.D{{"name", 1}}}}
	projectStage := bson.D{
		{Key: "$project",
			Value: bson.D{
				{"studentid", 1},
				{"name", 1},
				{"grade", 1},
				{"_id", 0},
				{"totalMarks",
					bson.D{
						{"$sum", bson.A{
							bson.D{{"$sum", "$marks.eng"}},
							bson.D{{"$sum", "$marks.lang"}},
							bson.D{{"$sum", "$marks.math"}},
						},
						},
					},
				},
			},
		},
	}
	cursor, err := collection.Aggregate(context.TODO(), mongo.Pipeline{matchStage, sortStage, projectStage})
	if err != nil {
		panic(err)
	}
	var results []bson.M

	err = cursor.All(context.TODO(), &results)
	for _, result := range results {
		fmt.Println("StudentID:", result["studentid"], "Name:", result["name"], "Grade:", result["grade"], "TotalMarks:", result["totalMarks"])
	}
}
