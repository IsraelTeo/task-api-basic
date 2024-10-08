package routes

import (
	"encoding/json"
	"net/http"

	"github.com/IsraelTeo/api-task/db"
	"github.com/IsraelTeo/api-task/models"
	"github.com/gorilla/mux"
)

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.GDB.Find(&users)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.GDB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var userSave models.User
	json.NewDecoder(r.Body).Decode(&userSave)
	createdUser := db.GDB.Create(&userSave)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusCreated)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var userUpdate models.User
	params := mux.Vars(r)
	db.GDB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	err := json.NewDecoder(r.Body).Decode(&userUpdate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid request"))
		return
	}
	user.FirstName = userUpdate.FirstName
	user.LastName = userUpdate.LastName
	user.Email = userUpdate.Email
	db.GDB.Save(&user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.GDB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	db.GDB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusNoContent)
}
