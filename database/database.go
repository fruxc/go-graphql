package database

import (
	"context"
	"log"
	"time"

	"github.com/fruxc/go-graphql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB Structure
type DB struct {
	client *mongo.Client
}

//Connect : for connection to database
func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	return &DB{
		client: client,
	}
}

//Save : for inserting data into database
func (db *DB) Save(input *model.NewBike) *model.Bike {
	collection := db.client.Database("vehicle").Collection("bikes")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	return &model.Bike{
		ID:         res.InsertedID.(primitive.ObjectID).Hex(),
		Name:       input.Name,
		IsGoodBike: input.IsGoodBike,
	}
}

//FindByID finds one document from the database
func (db *DB) FindByID(ID string) *model.Bike {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}
	collection := db.client.Database("vehicle").Collection("bikes")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	bike := model.Bike{}
	res.Decode(&bike)
	return &bike
}

//All : to get all the data from the database
func (db *DB) All() []*model.Bike {
	collection := db.client.Database("vehicle").Collection("bikes")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var bikes []*model.Bike
	for cur.Next(ctx) {
		var bike *model.Bike
		err := cur.Decode(&bike)
		if err != nil {
			log.Fatal(err)
		}
		bikes = append(bikes, bike)
	}
	return bikes
}
