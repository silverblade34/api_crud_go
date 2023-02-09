package repository

import (
	m "api-v2/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UserRepository es una interfaz para la persistencia de usuarios.
type UserRepository interface {
	ValidUser(email string, password string) (*m.User, error)
}

// UserMongoRepository es una implementación de la interfaz UserRepository que utiliza MongoDB como almacenamiento.
type UserMongoRepository struct {
	client *mongo.Client
}

// NewUserMongoRepository crea una nueva instancia de UserMongoRepository.
func NewUserMongoRepository(client *mongo.Client) *UserMongoRepository {
	return &UserMongoRepository{client: client}
}

// ValidUser verifica si un usuario existe en la base de datos.
func (r *UserMongoRepository) ValidUser(email, password string) (*m.User, error) {
	// Obtiene la colección "users_test" de la base de datos "test".
	collection := r.client.Database("test").Collection("users_test")
	// Estructura que se usará para almacenar el usuario que se obtiene de la base de datos.
	var user m.User
	// Busca el usuario en la base de datos.
	err := collection.FindOne(context.TODO(), bson.M{"email": email, "password": password}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func InsertUser(user m.User) error {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	err = client.Connect(context.TODO())
	if err != nil {
		return err
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("test").Collection("users_test")

	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}
