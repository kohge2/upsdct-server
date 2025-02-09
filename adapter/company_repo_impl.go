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

type CompanyRepository struct {
	db *database.DB
}

func NewCompanyRepository(db *database.DB) *CompanyRepository {
	return &CompanyRepository{db: db}
}

func (a *CompanyRepository) FindByCompanyID(ctx context.Context, companyID string) (*models.Company, error) {
	company := &models.Company{}
	if err := a.db.GetNewTxnOrContext(ctx).Where("id = ?", companyID).First(company).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewAppErr(utils.ErrTypeNotfound, utils.ErrMsgNotfound, http.StatusNotFound, err)
		}
		return nil, err
	}
	return company, nil
}
