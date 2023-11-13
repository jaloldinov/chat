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
	// GetAllActiveUser(context.Context, *models.GetAllUserRequest) (*models.GetAllUser, error)
	// UpdateUser(context.Context, *models.UpdateUser) (string, error)
	// DeleteUser(context.Context, *models.IdRequest) (string, error)

}

// type WSI interface {
// 	CreateUser(context.Context, *models.CreateUserReq) (*models.CreateUserRes, error)
// 	GetUserByEmail(context.Context, *models.LoginUserReq) (*models.User, error)
// 	// GetAllActiveUser(context.Context, *models.GetAllUserRequest) (*models.GetAllUser, error)
// 	// UpdateUser(context.Context, *models.UpdateUser) (string, error)
// 	// DeleteUser(context.Context, *models.IdRequest) (string, error)

// }
