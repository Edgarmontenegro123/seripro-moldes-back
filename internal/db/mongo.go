package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

// Conecta con MongoDB
func Connect(uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return fmt.Errorf("error conectando a MongoDB: %w", err)
	}

	// Verifica la conexión
	if err := client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("error al verificar conexión: %w", err)
	}

	log.Println("Conexión exitosa con MongoDB")
	MongoClient = client
	return nil
}

// Desconecta de MongoDB
func Disconnect() {
	if MongoClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := MongoClient.Disconnect(ctx); err != nil {
			log.Println("Error al desconectar MongoDB:", err)
		} else {
			log.Println("MongoDB desconectado correctamente")
		}
	}
}
