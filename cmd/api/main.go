package main

import (
	"fmt"
	"log"

	"go-backend-training/internal/server"
)

func main() {
	srv := server.New()

	fmt.Println("Servidor corriendo en http://localhost:8080")

	if err := srv.Start(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
