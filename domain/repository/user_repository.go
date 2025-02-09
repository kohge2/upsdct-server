package repository

import (
	"context"

	"github.com/kohge2/upsdct-server/domain/models"
)

type UserRepository interface {
	FindByUserID(ctx context.Context, userID string) (*models.User, error)
}
