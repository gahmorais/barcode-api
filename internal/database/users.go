package database

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserModel struct {
	DB *mongo.Database
}
type User struct{}
