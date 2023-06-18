package lib

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongo_ctx context.Context = context.TODO()

type DatabaseManagement interface {
	NewDB(db_name string)
	ConnectDB(collection_name string)
	FindOneUser(filter bson.D) bson.M
	DeleteOneQuery(filter bson.D) bool
}

type Database struct {
	DBUri                   string
	DBClient                *mongo.Client
	DBCurrentCollection     *mongo.Collection
	DBName                  string
	DBCurrentCollectionName string
}

// CREATE NEW DB
func NewDB(db_name string) Database {
	if db_name == "" {
		db_name = "posidonia"
	}

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	return Database{
		DBUri:                   os.Getenv("DB_URI"),
		DBClient:                &mongo.Client{},
		DBCurrentCollection:     &mongo.Collection{},
		DBName:                  db_name,
		DBCurrentCollectionName: "",
	}
}

// DB CONNECTION
func (db *Database) ConnectDB(collection_name string) *mongo.Collection {

	db.DBCurrentCollectionName = collection_name

	if db.DBUri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(mongo_ctx, options.Client().ApplyURI(db.DBUri))
	if err != nil {
		panic(err)
	}

	db.DBClient = client

	collection := client.Database(db.DBName).Collection(db.DBCurrentCollectionName)
	db.DBCurrentCollection = collection

	return collection
}

/// DB OPERATIONS:

func (db *Database) FindOneUser(document bson.D) (User, error) {
	var result User
	err := db.DBCurrentCollection.FindOne(mongo_ctx, document).Decode(&result)

	if err == mongo.ErrNoDocuments {
		LogError(fmt.Errorf("ERR: Document not found in the collection: %s", db.DBCurrentCollectionName))
		return User{
			UserID:   primitive.NewObjectID(),
			Username: "NOT_FOUND",
			Password: "NOT_FOUND",
		}, errors.New("ERROR: DOCUMENT NOT FOUNDED")
	}
	LogError(err)

	return result, nil
}

func (db *Database) DeleteOneQuery(filter bson.D) bool {
	_, err := db.DBCurrentCollection.DeleteOne(mongo_ctx, filter)
	if err != nil {
		LogError(errors.New("ERROR: WE CAN'T DELETE DOCUMENT FROM DB FOR NOW"))
		return false
	}
	return true
}

func (db *Database) AddLocation(doc bson.D) error {
	result, err := db.DBCurrentCollection.InsertOne(mongo_ctx, doc)
	log.Println(result)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) DumpLocations() ([]Locations, error) {
	cursor, err := db.DBCurrentCollection.Find(mongo_ctx, bson.M{})
	if err != nil {
		non_found_oid, _ := primitive.ObjectIDFromHex("0")
		return []Locations{
			{
				LocationID: non_found_oid,
				Polygon:    []Coordinates{{float64(0), float64(0)}},
				AddedTime:  "",
				AddedBy:    "",
			},
		}, err
	}

	defer cursor.Close(mongo_ctx)

	result := []Locations{}

	for cursor.Next(mongo_ctx) {
		var locations Locations

		if err = cursor.Decode(&locations); err != nil {
			panic(err)
		}

		result = append(result, locations)

	}
	return result, nil
}
