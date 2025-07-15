package api

import (
	"database/sql"
	"game-assistant-server/internal/config"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter(db *sql.DB, config *config.Config) *http.Server {
	router := mux.NewRouter()
	setupCardRouter(router, db)
	return &http.Server{
		Addr:    config.ServerAddr,
		Handler: router,
	}
}

func setupCardRouter(router *mux.Router, db *sql.DB) {
	h := &CardHandler{db}
	router.HandleFunc("/api/cards", h.GetCards).Methods("GET")
	router.HandleFunc("/api/cards/{id}", h.GetCard).Methods("GET")
	router.HandleFunc("/api/cards", h.InsertCard).Methods("POST")
}
