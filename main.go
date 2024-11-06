package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/riz-it/go-jwt-mux/controller/authcontroller"
	"github.com/riz-it/go-jwt-mux/controller/productcontroller"
	"github.com/riz-it/go-jwt-mux/middleware"
	"github.com/riz-it/go-jwt-mux/model"
)

func main() {

	model.DatabaseConnection()

	r := mux.NewRouter()

	auth := r.PathPrefix("/auth").Subrouter()

	auth.HandleFunc("/login", authcontroller.Login).Methods("POST")
	auth.HandleFunc("/register", authcontroller.Register).Methods("POST")
	auth.HandleFunc("/logout", authcontroller.Logout).Methods("DELETE")

	product := r.PathPrefix("/products").Subrouter()
	product.Use(middleware.JWTMiddleware)
	product.HandleFunc("", productcontroller.Index).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))

}
