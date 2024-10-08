package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/IsraelTeo/api-task/db"
	"github.com/IsraelTeo/api-task/models"
	"github.com/IsraelTeo/api-task/routes"
	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Connecting...")

	err := db.DBConnection()
	if err != nil {
		log.Fatalf("error trying to connect with database: %v", err)
	}

	fmt.Println("conecction ok")

	errorMigrate := db.GDB.AutoMigrate(&models.User{},
		&models.Task{})

	if errorMigrate != nil {
		log.Fatalf("error trying to migrate database (User): %v", errorMigrate)
	}

	fmt.Println("Migración exitosa")

	r := mux.NewRouter()
	r.HandleFunc("/users", routes.GetAllUsersHandler).Methods("GET")
	r.HandleFunc("/user/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/user/{id}", routes.DeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/user", routes.CreateUserHandler).Methods("POST")
	r.HandleFunc("/user/{id}", routes.UpdateUserHandler).Methods("PUT")

	http.ListenAndServe(":8080", r)

}
