package postgres

import (
	"auth/models"
	"context"
	"fmt"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (b *userRepo) CreateUser(c context.Context, req *models.CreateUserReq) (*models.CreateUserRes, error) {
	user := models.CreateUserRes{}
	var lastInsertId int
	query := `INSERT INTO users(
					username, 
					password, 
					email
					) VALUES ($1, $2, $3) returning id`
	_, err := b.db.Exec(context.Background(), query,
		req.Username,
		req.Password,
		req.Email,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	user.ID = lastInsertId
	user.Email = req.Email
	user.Username = req.Username

	return &user, nil
}

func (b *userRepo) GetUserByEmail(c context.Context, req *models.LoginUserReq) (resp *models.User, err error) {

	u := models.User{}
	query := `SELECT 
				id, 
				email, 
				username, 
				password 
				FROM users WHERE email = $1`

	err = b.db.QueryRow(context.Background(), query, req.Email).Scan(
		&u.ID,
		&u.Email,
		&u.Username,
		&u.Password,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &u, nil
}

/*
func (r *userRepo) AdminCreate(ctx context.Context, data AdminCreateRequest) (AdminCreateResponse, *pkg.Error) {
	var detail AdminCreateResponse

	dataCtx, er := r.CheckCtx(ctx)
	if er != nil {
		return AdminCreateResponse{}, er
	}

	if data.File != nil {
		var folder string
		if data.Type == 1 {
			folder = "post/docs"
		} else if data.Type == 2 {
			folder = "post/images"
		} else if data.Type == 3 {
			folder = "post/videos"
		}
		fileLink, err := r.fileService.Upload(ctx, data.File, folder)
		if err != nil {
			return AdminCreateResponse{}, err
		}

		data.Link = &fileLink
	}

	detail.Id = uuid.NewString()
	detail.Type = data.Type
	detail.Link = data.Link
	detail.PostId = data.PostId

	detail.CreatedAt = time.Now()
	detail.CreatedBy = dataCtx.UserId

	_, err := r.NewInsert().Model(&detail).Exec(ctx)
	if err != nil {
		er := r.fileService.Delete(ctx, *detail.Link)
		if er != nil {
			return AdminCreateResponse{}, er
		}
		return AdminCreateResponse{}, &pkg.Error{
			Err:    pkg.WrapError(err, "creating post file"),
			Status: http.StatusInternalServerError,
		}
	}

	return detail, nil
}
*/
