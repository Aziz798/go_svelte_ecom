package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go_ecom/internal/models"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router, db *sql.DB) {
	route := router.PathPrefix("/api/v1/user").Subrouter()
	route.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		registerUser(w, r, db)
	}).Methods("POST")
	route.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		loginUser(w, r, db)
	}).Methods("POST")
}

func registerUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := CreateUser(user, db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func loginUser(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var user models.LoginUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(user)
	token, err := LoginUser(user, db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
