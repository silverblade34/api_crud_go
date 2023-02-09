package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	m "api-v2/model"
	r "api-v2/repository"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ValidUser(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	opts := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al conectarse a la base de datos"})
		return
	}
	// Crea una instancia de UserMongoRepository.
	userRepo := r.NewUserMongoRepository(client)

	// Llamar la función ValidUser con el email y password proporcionados en la request.
	user, err := userRepo.ValidUser(req.Email, req.Password)
	if err != nil {
		log.Println("Error: ", err)
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el usuario"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user m.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := r.InsertUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo insertar el usuario en la base de datos"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Usuario creado exitosamente"})
}

func UpdateUser(c *gin.Context) {
	// lógica para actualizar un usuario
}

func DeleteUser(c *gin.Context) {
	// lógica para borrar un usuario
}
