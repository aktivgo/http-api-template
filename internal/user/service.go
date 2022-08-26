package user

import (
	"context"
	"wire-di/internal/domain"
)

type service struct {
	repo domain.UserRepository
}

func (s *service) FetchByUsername(
	ctx context.Context,
	username string,
) (*domain.User, error) {
	userEntity, err := s.repo.FetchByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:       userEntity.ID,
		Username: userEntity.Username,
	}, nil
}
