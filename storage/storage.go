package storage

import (
	"auth/models"
	"context"
)

type StorageI interface {
	User() UsersI
	Message() MessageI
}

type UsersI interface {
	CreateUser(context.Context, *models.CreateUserReq) (*models.CreateUserRes, error)
	GetUserByEmail(context.Context, *models.LoginUserReq) (*models.User, error)
}

type MessageI interface {
	CreateMessage(context.Context, *models.Message) (string, error)
}
