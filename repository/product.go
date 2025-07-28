package repository

import (
	"context"
	"time"

	"github.com/barcode-api/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ProductRepository struct {
	database   *mongo.Database
	collection string
}

func NewProductRepository(db *mongo.Database) ProductRepository {
	return ProductRepository{
		database:   db,
		collection: "products",
	}
}

func (p *ProductRepository) create(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := p.database.Collection(p.collection)
	_, err := collection.InsertOne(ctx, user)
	return err
}

func (p *ProductRepository) update() error {
	return nil
}

func (p *ProductRepository) getById(id uint) models.Product {
	return models.Product{}
}

func (p *ProductRepository) getAll() []models.Product {
	return []models.Product{}
}
