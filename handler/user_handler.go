package handler

import (
	"context"
	"net/http"

	m "api-v2/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ValidUser verifica que un usuario existe en la base de datos.
func ValidUser(c *gin.Context) {
	// Estructura que se espera recibir en el cuerpo de la petición.
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Se verifica que la estructura enviada en el cuerpo sea válida.
	if err := c.ShouldBindJSON(&req); err != nil {
		// Si la estructura no es válida, se envía un error con el código 400 Bad Request.
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Se obtienen los valores de email y password.
	email := req.Email
	password := req.Password
	// Se crea un nuevo cliente de MongoDB.
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		// Si no se pudo conectar a la base de datos, se envía un error con el código 500 Internal Server Error.
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo conectar a la base de datos"})
		return
	}
	err = client.Connect(context.TODO())
	if err != nil {
		// Si no se pudo conectar a la base de datos, se envía un error con el código 500 Internal Server Error.
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo conectar a la base de datos"})
		return
	}
	// Se cierra la conexión a la base de datos al finalizar la función.
	defer client.Disconnect(context.TODO())

	// Se obtiene la colección "users_test" de la base de datos "test".
	collection := client.Database("test").Collection("users_test")
	// Estructura que se usará para almacenar el usuario que se obtiene de la base de datos.
	var user m.User
	// Se busca el usuario en la base de datos.
	err = collection.FindOne(context.TODO(), bson.M{"email": email, "password": password}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Si el usuario no existe, devuelve un error de tipo "Usuario no encontrado".
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el usuario"})
		return
	}
	// Si se encontró el usuario, se devuelve en formato JSON.
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user m.User
	// Se realiza un "bind" de los datos del usuario recibidos en el cuerpo de la petición.
	if err := c.ShouldBindJSON(&user); err != nil {
		// Si ocurre un error en el proceso de "bind", se devuelve un error de tipo "Solicitud inválida".
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Se conecta a la base de datos MongoDB.
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		// Si no se puede conectar a la base de datos, se devuelve un error de tipo "Error interno".
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo conectar a la base de datos"})
		return
	}
	err = client.Connect(context.TODO())
	if err != nil {
		// Si no se puede conectar a la base de datos, se devuelve un error de tipo "Error interno".
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo conectar a la base de datos"})
		return
	}
	defer client.Disconnect(context.TODO())

	// Se obtiene la colección "users_test" en la base de datos "test".
	collection := client.Database("test").Collection("users_test")

	// Se inserta el nuevo usuario en la colección.
	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		// Si no se puede insertar el usuario, se devuelve un error de tipo "Error interno".
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo insertar el usuario en la base de datos"})
		return
	}

	// Si el usuario se crea exitosamente, se devuelve un mensaje de éxito.
	c.JSON(http.StatusCreated, gin.H{"message": "Usuario creado exitosamente"})

}

func UpdateUser(c *gin.Context) {
	// lógica para actualizar un usuario
}

func DeleteUser(c *gin.Context) {
	// lógica para borrar un usuario
}
