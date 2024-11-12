package main

import (
	"log"
	"net/http"

	"github.com/amartinezh/ms-crosslisting/pkg/controller"
	"github.com/amartinezh/ms-crosslisting/pkg/service"
	"github.com/amartinezh/sdk-db/sdk_postgres"
	"github.com/gorilla/mux"
)

func main() {
	// Configuración de la conexión a la base de datos
	connString := "postgres://postgres:s3rv3r@localhost:5432/crosslist"

	// Inicializar el SDK de PostgreSQL
	dbSDK := sdk_postgres.NewPostgresSDK()
	if err := dbSDK.Connect(connString); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer dbSDK.Close()

	// Inicializar el controlador de personas
	personService := service.NewPersonService(dbSDK)
	personController := controller.NewPersonController(personService)

	// Configurar rutas
	r := mux.NewRouter()
	r.HandleFunc("/persons", personController.AddPerson).Methods("POST")
	r.HandleFunc("/persons", personController.ListPersons).Methods("GET")

	// Iniciar el servidor
	log.Println("Server started on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
