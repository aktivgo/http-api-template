package user

import (
	"context"
	"database/sql"
	"wire-di/internal/domain"

	"github.com/semichkin-gopkg/uuid"
)

type repository struct {
	db *sql.DB
}

func (r *repository) FetchByUsername(
	ctx context.Context,
	username string,
) (*domain.UserEntity, error) {
	return &domain.UserEntity{
		ID:       uuid.New(),
		Username: "aktivgo",
		Password: "12345",
	}, nil
}
