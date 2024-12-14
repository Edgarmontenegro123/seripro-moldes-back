package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var MongoURI string

func LoadEnv() {
	// Carga variables del archivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env, usando variables de entorno")
	}

	MongoURI = os.Getenv("MONGO_URI")
	if MongoURI == "" {
		log.Fatal("La variable MONGO_URI no está configurada")
	}
}
