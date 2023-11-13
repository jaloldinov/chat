package storage

import (
	"auth/models"
	"context"
)

type StorageI interface {
	User() UsersI
	// WS() WSI
}

type UsersI interface {
	CreateUser(context.Context, *models.CreateUserReq) (*models.CreateUserRes, error)
	GetUserByEmail(context.Context, *models.LoginUserReq) (*models.User, error)
}
