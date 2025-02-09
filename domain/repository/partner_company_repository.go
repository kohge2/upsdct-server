package repository

import (
	"context"

	"github.com/kohge2/upsdct-server/domain/models"
)

type PartnerCompanyRepository interface {
	FindByPartnerCompanyID(ctx context.Context, partnerCompanyID string) (*models.PartnerCompany, error)
}
