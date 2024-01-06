package models

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var COLLECTION_NAME = "urls"

type UrlPayload struct {
	Id     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Link   string             `json:"link" required:"true"`
	Label  string             `json:"label"`
	Active bool               `json:"active" default:"true"`
	UserId string             `json:"user_id"`
}

type UrlResult struct {
	Id     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Label  string             `json:"label"`
	Active bool               `json:"active" default:"true"`
}

func (u *UrlPayload) Create(db *mongo.Client) error {
	data := bson.M{
		"link":    u.Link,
		"label":   u.Label,
		"active":  true,
		"user_id": u.UserId,
	}
	collection := db.Database(os.Getenv("DATABASE_NAME")).Collection(COLLECTION_NAME)
	_, insertError := collection.InsertOne(context.Background(), data)
	if insertError != nil {
		fmt.Println(insertError)
		return insertError
	}
	return nil
}

func (u *UrlPayload) Get(db *mongo.Client) error {
	collection := db.Database(os.Getenv("DATABASE_NAME")).Collection(COLLECTION_NAME).FindOne(context.Background(), bson.M{"_id": u.Id})
	decodeError := collection.Decode(&u)
	if decodeError != nil {
		fmt.Println(decodeError)
		return decodeError
	}
	return nil
}

func (u *UrlPayload) Update(id *primitive.ObjectID, db *mongo.Client) error {
	return nil
}

func (u *UrlPayload) Delete(db *mongo.Client) error {
	collection := db.Database(os.Getenv("DATABASE_NAME")).Collection(COLLECTION_NAME)
	result := collection.FindOneAndDelete(context.Background(), bson.M{"_id": u.Id})
	if result.Err() != nil {
		fmt.Println(result.Err())
		return result.Err()
	}
	return nil
}

func (u *UrlPayload) GetAll(db *mongo.Client) (urls []UrlResult, err error) {
	collection := db.Database(os.Getenv("DATABASE_NAME")).Collection(COLLECTION_NAME)
	cursor, err := collection.Find(context.Background(), bson.M{"user_id": u.UserId})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var url = UrlResult{}
		error := cursor.Decode(&url)
		if error != nil {
			fmt.Println(error)
			return nil, error
		}
		urls = append(urls, url)
	}
	return urls, nil
}
