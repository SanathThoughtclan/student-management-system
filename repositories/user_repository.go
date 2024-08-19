package repositories

import (
	"context"
	"log"
	"student-management-system/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		collection: db.Collection("users"),
	}
}

func (r *UserRepository) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	filter := bson.D{{Key: "username", Value: username}} // Correct usage of bson.E
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByID(ctx context.Context, ID string) (*models.User, error) {
	var user models.User
	filter := bson.D{{Key: "username", Value: ID}} // Correct usage of bson.E
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(ctx, user)
}
