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

type DatabaseManagement interface {
	NewDB(db_name string)
	ConnectDB(collection_name string)
	FindOneQuery(filter bson.D) bson.M
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

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBUri))
	if err != nil {
		panic(err)
	}

	db.DBClient = client

	collection := client.Database(db.DBName).Collection(db.DBCurrentCollectionName)
	db.DBCurrentCollection = collection

	return collection
}

// DB OPERATIONS

func (db *Database) FindOneQuery(document any) (User, error) {
	var result User
	err := db.DBCurrentCollection.FindOne(context.TODO(), document).Decode(&result)

	if err == mongo.ErrNoDocuments {
		LogError(errors.New(fmt.Sprintf("ERR: Document not found in the collection: %s", db.DBCurrentCollection)))
		return User{
			UserID:   primitive.NewObjectID(),
			Username: "NOT_FOUND",
			Password: "NOT_FOUND",
		}, errors.New("ERROR DOCUMENT NOT FOUNDED!")
	}
	LogError(err)

	return result, nil
}

func (db *Database) DeleteOneQuery(filter bson.D) bool {
	_, err := db.DBCurrentCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		LogError(errors.New("WE CAN'T DELETE DOCUMENT FROM DB FOR NOW!"))
		return false
	}
	return true
}
