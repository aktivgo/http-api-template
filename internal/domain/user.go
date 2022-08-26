package domain

import (
	"context"
	"net/http"

	"github.com/semichkin-gopkg/uuid"
)

type (
	User struct {
		ID       uuid.UUID `json:"id"`
		Username string    `json:"username"`
	}

	UserEntity struct {
		ID       uuid.UUID
		Username string
		Password string
	}

	UserRepository interface {
		FetchByUsername(ctx context.Context, username string) (*UserEntity, error)
	}

	UserService interface {
		FetchByUsername(ctx context.Context, username string) (*User, error)
	}

	UserHandler interface {
		FetchByUsername() http.HandlerFunc
	}
)
