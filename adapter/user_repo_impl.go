package adapter

import (
	"context"
	"errors"
	"net/http"

	"github.com/kohge2/upsdct-server/database"
	"github.com/kohge2/upsdct-server/domain/models"
	"github.com/kohge2/upsdct-server/utils"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *database.DB
}

func NewUserRepository(db *database.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (a *UserRepository) FindByUserID(ctx context.Context, userID string) (*models.User, error) {
	user := &models.User{}
	if err := a.db.GetNewTxnOrContext(ctx).
		Where("id = ?", userID).
		First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewAppErr(utils.ErrTypeNotfound, utils.ErrMsgNotfound, http.StatusNotFound, err)
		}
		return nil, err
	}
	return user, nil
}
