package lib

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidateLogin(username, password string) (User, error) {
	user, err := FindUser(username, password)
	return user, err
}

func AddPolygon(compressed_polygon_list []string, addedBy, addedTime string, isInDanger bool) error {
	godotenv.Load()

	database_users := NewDB(os.Getenv("DB_NAME"))
	database_locations := NewDB(os.Getenv("DB_NAME"))

	database_users.ConnectDB(os.Getenv("DB_COLLECTION_USER"))
	database_locations.ConnectDB(os.Getenv("DB_COLLECTION_LOCATION"))

	result, err := database_users.FindOneUser(bson.D{
		{Key: "username", Value: addedBy},
	})
	LogError(err)

	if (err != nil) || (result.Username == "NOT_FOUND") {
		return errors.New("we cannot able to add item")
	}

	polygon_list := [][]float64{}

	for _, item := range compressed_polygon_list {
		polygon_item := strings.Split(item, ":")
		lat, err := strconv.ParseFloat(polygon_item[0], 64)
		LogError(err)
		lng, err := strconv.ParseFloat(polygon_item[1], 64)
		LogError(err)
		polygon_list = append(polygon_list, []float64{lat, lng})
	}

	err = database_locations.AddLocation(bson.D{
		{Key: "polygon", Value: polygon_list},
		{Key: "added_time", Value: addedTime},
		{Key: "added_by", Value: addedBy},
		{Key: "is_in_danger", Value: isInDanger},
	})

	if err != nil {
		return err
	}

	return nil
}

func FetchAllPolygons() ([]Locations, error) {
	godotenv.Load()

	database_locations := NewDB(os.Getenv("DB_NAME"))
	database_locations.ConnectDB(os.Getenv("DB_COLLECTION_LOCATION"))

	locations, err := database_locations.DumpLocations()

	return locations, err
}

func FindUser(username, password string) (User, error) {
	godotenv.Load()

	database := NewDB(os.Getenv("DB_NAME"))

	database.ConnectDB(os.Getenv("DB_COLLECTION_USER"))

	result, err := database.FindOneUser(bson.D{
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
	LogError(err)
	result, err := database.FindOneUser(bson.D{{Key: "_id", Value: ObjectID_user}})
	LogError(err)

	return result, err

}

func GeoCodeQuery(latitude, longitude string) (string, int, error) {

	query_url := fmt.Sprintf(
		"https://api.geoapify.com/v1/geocode/reverse?lat=%s&lon=%s&apiKey=%s",
		latitude,
		longitude,
		config.GetApiKeys("geocode"),
	)

	response, err := http.Get(query_url)
	LogError(err)

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	LogError(err)

	return string(body), response.StatusCode, err
}
