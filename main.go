package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Edgarmontenegro123/seripro-moldes-back/config"
	"github.com/Edgarmontenegro123/seripro-moldes-back/internal/db"
)

func main() {
	// Carga configuración
	config.LoadEnv()

	// Conecta a MongoDB
	if err := db.Connect(config.MongoURI); err != nil {
		log.Fatalf("No se pudo conectar a MongoDB: %v", err)
	}
	defer db.Disconnect()

	log.Println("Servidor corriendo... Presiona Ctrl+C para salir")

	// Espera señal para salir
	waitForExit()
}

func waitForExit() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Apagando servidor...")
}
