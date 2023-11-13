package handler

import (
	"auth/pkg/logger"
	"auth/storage"
)

type Handler struct {
	storage storage.StorageI
	hub     *Hub
	log     logger.LoggerI
}

func NewHandler(strg storage.StorageI, hub *Hub, loger logger.LoggerI) *Handler {
	return &Handler{storage: strg, hub: hub, log: loger}
}
