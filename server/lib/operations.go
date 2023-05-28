package lib

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidateLogin(username, password string) (User, error) {
	user, err := FindUser(username, password)
	return user, err
}

func FindUser(username, password string) (User, error) {
	godotenv.Load()

	database := NewDB(os.Getenv("DB_NAME"))

	database.ConnectDB(os.Getenv("DB_COLLECTION_USER"))

	result, err := database.FindOneQuery(bson.D{
		{Key: "username", Value: username},
		{Key: "password", Value: password},
	})
	LogError(err)
	return result, err
}

func GetTime() string {
	// RFC1123 because i liked a little bit much than others.
	// Format: Mon, 02 Jan 2006 15:04:05 MST
	return time.Now().Format(time.RFC1123)
}

func GetUserByID(user_id string) (User, error) {
	godotenv.Load()

	database := NewDB(os.Getenv("DB_NAME"))
	database.ConnectDB(os.Getenv("DB_COLLECTION_USER"))

	ObjectID_user, err := primitive.ObjectIDFromHex(user_id)
	result, err := database.FindOneQuery(bson.D{{Key: "_id", Value: ObjectID_user}})
	LogError(err)

	return result, err

}
