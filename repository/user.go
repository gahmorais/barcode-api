package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/barcode-api/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	database   *mongo.Database
	collection string
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return UserRepository{
		database:   db,
		collection: "users",
	}
}

func (u *UserRepository) Login(username string, password string) error {
	user, err := u.GetByUsername(username)
	if err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("usuário ou senha incorretos")
	}

	hashedPassword := user.Password
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return fmt.Errorf("usuário ou senha incorretos")
	}
	return nil
}

func (u *UserRepository) Create(username string, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defaultCost := bcrypt.DefaultCost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), defaultCost)
	if err != nil {
		return err
	}
	newUser := models.User{
		UserName: username,
		Password: string(hashedPassword),
	}
	_, err = u.database.Collection(u.collection).InsertOne(ctx, newUser)
	return err
}

func (u *UserRepository) updatePassword(newPassword string) error {
	return nil
}

func (u *UserRepository) GetByUsername(username string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{{Key: "username", Value: username}}
	var user models.User
	err := u.database.Collection(u.collection).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.User{}, fmt.Errorf("nenhum usuário encontrado")
		}
		return models.User{}, err
	}
	return user, nil
}
