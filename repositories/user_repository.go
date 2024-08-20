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
	filter := bson.D{{Key: "username", Value: username}}
	log.Printf("Querying for username: %s with filter: %v", username, filter)
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("No user found for username: %s", username)
			return nil, nil
		}
		log.Printf("Error fetching user by username: %v", err)
		return nil, err
	}
	log.Printf("Found user: %v", user)
	return &user, nil
}

func (r *UserRepository) GetByID(ctx context.Context, ID string) (*models.User, error) {
	var user models.User
	filter := bson.D{{Key: "user_id", Value: ID}}
	log.Printf("Querying for user ID: %s with filter: %v", ID, filter)
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("No user found for ID: %s", ID)
			return nil, nil
		}
		log.Printf("Error fetching user by ID: %v", err)
		return nil, err
	}
	log.Printf("Found user: %v", user)
	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(ctx, user)
}
