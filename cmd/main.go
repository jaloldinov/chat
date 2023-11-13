package main

import (
	"auth/api"
	"auth/api/handler"
	"auth/config"
	"auth/pkg/logger"
	"auth/storage/postgres"

	"context"
	"fmt"
)

func main() {
	cfg := config.Load()
	log := logger.NewLogger("chat app", logger.LevelInfo)
	strg, err := postgres.NewStorage(context.Background(), cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	hub := handler.NewHub()
	go hub.Run()

	h := handler.NewHandler(strg, hub, log)

	r := api.NewServer(h)
	r.Run(fmt.Sprintf(":%s", cfg.Port))
}
