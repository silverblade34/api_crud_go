package repository

import (
	model "api-v2/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UserRepository representa el repositorio de usuarios.
type UserRepository struct {
	client *mongo.Client
}

// NewUserRepository crea una nueva instancia de UserRepository.
func NewUserRepository() (*UserRepository, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}
	return &UserRepository{client: client}, nil
}

// GetByEmailAndPassword obtiene un usuario por su email y password.
func (r *UserRepository) GetByEmailAndPassword(email, password string) (*model.User, error) {
	collection := r.client.Database("test").Collection("users_test")
	var user model.User
	err := collection.FindOne(context.TODO(), bson.M{"email": email, "password": password}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
