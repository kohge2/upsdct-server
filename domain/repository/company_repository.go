package repository

import (
	"context"

	"github.com/kohge2/upsdct-server/domain/models"
)

type CompanyRepository interface {
	FindByCompanyID(ctx context.Context, companyID string) (*models.Company, error)
}
