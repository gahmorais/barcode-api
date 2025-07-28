package db

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	db *mongo.Database
)

func InitDb(strConn string, database string) error {
	clientOptions := options.Client().ApplyURI(strConn)

	var err error
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return err
	}

	db = client.Database(database)
	return nil
}

func GetDb() *mongo.Database {
	return db
}
