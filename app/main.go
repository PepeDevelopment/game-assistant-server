package main

import (
	"game-assistant-server/internal/api"
	"game-assistant-server/internal/config"
	"game-assistant-server/internal/db"
	"log"
)

func main() {
	cfg := config.Load()
	log.Printf("Running in %s mode, listening on %s", cfg.Env, cfg.ServerAddr)
	dbConnection := db.Connect(cfg.DatabaseURL)
	defer dbConnection.Close()
	httpServer := api.SetupRouter(dbConnection, &cfg)
	log.Fatal(httpServer.ListenAndServe())
}
