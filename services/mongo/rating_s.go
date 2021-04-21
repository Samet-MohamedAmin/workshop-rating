package rating

import (
	"context"
	"fmt"
	"log"
	"os"
	models "rating/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	rs         = models.Ratings{}
	ctx        context.Context
	client     *mongo.Client
	collection *mongo.Collection
)

const (
	errEmptyParam = "%s param is empty"
	errBadParam   = "%s param must be an integer value between 0 and 5 inclusive"
)

func InitDB() {
	password := os.Getenv("DB_PASSWORD")
	username := os.Getenv("DB_USERNAME")
	uriBase := os.Getenv("DB_URI")
	uri := fmt.Sprintf(uriBase, username, password)
	// TODO: should defer cancel context
	ctx = context.Background()
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	// TODO: client should disconnect
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	collection = client.Database("ExaTech").Collection("NginxDemo")
}

func EmptyCollection() {
	collection.Drop(ctx)
}

func GetAll() (ratings []models.RatingItem) {
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalln(err)
	}
	if err := cursor.All(ctx, &ratings); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(ratings)
	return
}

func GetOne(ip string) (ri models.RatingItem, ok bool) {
	filter := bson.D{primitive.E{Key: "ip", Value: ip}}
	err := collection.FindOne(ctx, filter).Decode(&ri)
	if err != nil {
		log.Println(err)
		ok = false
	} else {
		log.Println(ri)
		ok = true
	}
	return
}

func AddOne(ip string, rating int) {
	// if there is no entity with that ip
	ri, ok := GetOne(ip)
	if !ok {
		record := models.RatingItem{Ip: ip, Rating: rating}
		r, err := collection.InsertOne(ctx, record)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("object %v has been created", r)
		}
		return
	}

	// if there is an entity with that ip
	update := bson.M{"$set": bson.M{"rating": rating}}
	if _, err := collection.UpdateByID(ctx, ri.ID, update); err != nil {
		log.Fatalln(err)
	} else {
		log.Printf("object with id %v has been updated", ri.ID)
	}

}

type AverageQuery struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AvgValue float64            `bson:"AvgValue,omitempty"`
}

func GetAverage() (avg float64) {
	avgKey := "AvgValue"
	groupStage := bson.D{
		{
			Key: "$group",
			Value: bson.D{
				{
					Key:   "_id",
					Value: nil},
				{
					Key: avgKey,
					Value: bson.D{
						{Key: "$avg", Value: "$rating"},
					},
				},
			},
		},
	}
	cursor, err := collection.Aggregate(ctx, mongo.Pipeline{groupStage})
	if err != nil {
		log.Fatalln(err)
		return
	}
	var result AverageQuery
	cursor.Next(ctx)
	if err = cursor.Decode(&result); err != nil {
		log.Fatalln(err)
		return
	}

	return result.AvgValue
}
