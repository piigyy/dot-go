package repository

import (
	"context"
	"database/sql"

	"github.com/piigyy/dot-go/internal/model"
)

type UserRepositoryPGSQL struct {
	db *sql.DB
}

func NewUserRepositoryPGSQL(db *sql.DB) *UserRepositoryPGSQL {
	return &UserRepositoryPGSQL{db: db}
}

func (r UserRepositoryPGSQL) CreateUser(ctx context.Context, user model.User) (id int, err error) {
	if err = r.db.QueryRowContext(ctx, QueryCreateUser,
		user.Fullname,
		user.Email,
		user.BirthDate.Format("2006-01-02"),
	).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
