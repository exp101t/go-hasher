package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/exp101t/go-hasher/api/internal/controllers/hash"
	"github.com/exp101t/go-hasher/api/internal/middlewares"
)

func main() {
	router := mux.NewRouter()

	router.PathPrefix("/api/hash").Handler(
		http.StripPrefix("/api/hash", hash.NewHasherRouter()))
	router.Use(middlewares.JsonHeaderMiddleware)

	http.ListenAndServe(":8000", router)
}
