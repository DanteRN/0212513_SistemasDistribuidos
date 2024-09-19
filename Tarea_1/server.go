package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Estructura para el usuario
type Usuario struct {
	Nombre string `json: "nombre"`
	Edad   int    `json; "edad"`
}

// Variable global para almacenar usuarios
var usuarios []Usuario

func agregarUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usuarios = append(usuarios, usuario)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usuario)
}

// Handler para obtener la lista de datos

func obtenerUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}

func main() {

	//Endpoint para agregar usuario
	http.HandleFunc("/usuario", agregarUsuario)

	//Endpoint para obtener los usuarios
	http.HandleFunc("/usuarios", obtenerUsuarios)

	//Iniciar el servidor en el puerto 8080
	fmt.Println("Servidor iniciado en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
