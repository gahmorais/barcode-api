package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	db *mongo.Database
)

func InitDb(strConn string, database string) error {
	clientOptions := options.Client().ApplyURI(strConn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	db = client.Database(database)
	return nil
}

func GetDb() *mongo.Database {
	return db
}
