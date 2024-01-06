package models

import (
	"context"
	"fmt"
	"os"

	"devcircle.space/mini-url/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection = "users"

type User struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Username string             `json:"username" bson:"username" required:"true"`
	Email    string             `json:"email" bson:"email" required:"true"`
	Password string             `json:"password" bson:"password" required:"true"`
}

type UserLogin struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Email    string             `json:"email" bson:"email" required:"true"`
	Password string             `json:"password" bson:"password" required:"true"`
}

type UserRegister struct {
	Username string `json:"username" bson:"username" required:"true"`
	Email    string `json:"email" bson:"email" required:"true"`
	Password string `json:"password" bson:"password" required:"true"`
}

type UserResponse struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Username string             `json:"username" bson:"username" required:"true"`
	Email    string             `json:"email" bson:"email" required:"true"`
}

func (u *UserRegister) Create(db *mongo.Client) error {
	collection := db.Database(os.Getenv("DATABASE_NAME")).Collection(userCollection)
	hashPassword, hashError := utils.GeneratePasswordHash(u.Password)
	if hashError != nil {
		fmt.Println(hashError)
		return hashError
	}
	_, insertError := collection.InsertOne(context.Background(), bson.M{
		"username": u.Username,
		"email":    u.Email,
		"password": string(hashPassword),
	})
	if insertError != nil {
		fmt.Println(insertError)
		return insertError
	}
	return nil
}

func (u *UserLogin) Find(db *mongo.Client) (User, error) {
	collection := db.Database(os.Getenv("DATABASE_NAME")).Collection(userCollection)
	var user User
	filter := bson.M{"$or": []interface{}{
		bson.M{"_id": u.Id},
		bson.M{"email": u.Email},
	}}
	decodeError := collection.FindOne(context.Background(), filter).Decode(&user)
	if decodeError != nil {
		fmt.Println(decodeError)
		return User{}, decodeError
	}
	return user, nil
}

func Update() {}

func Delete() {}
