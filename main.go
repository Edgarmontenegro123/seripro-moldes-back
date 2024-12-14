package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := "8080"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bienvenido a Seripro moldes!")
	})

	fmt.Printf("Servidor corriendo en http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error al iniciar el servidor: ", err)
	}
}
