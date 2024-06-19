package product

import (
	"database/sql"
	"go_ecom/internal/models"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router, db *sql.DB) {
	route := router.PathPrefix("/api/v1/product").Subrouter()
	route.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		createProductHandler(w, r, db)
	}).Methods("POST")
}

func createProductHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var product models.Product
	token := r.Header.Get("token")
	if token == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

}
