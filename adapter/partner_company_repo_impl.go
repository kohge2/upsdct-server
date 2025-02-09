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

type PartnerCompanyRepository struct {
	db *database.DB
}

func NewPartnerCompanyRepository(db *database.DB) *PartnerCompanyRepository {
	return &PartnerCompanyRepository{db: db}
}

func (a *PartnerCompanyRepository) FindByPartnerCompanyID(ctx context.Context, partnerCompanyID string) (*models.PartnerCompany, error) {
	partnerCompany := &models.PartnerCompany{}
	if err := a.db.GetNewTxnOrContext(ctx).
		Where("id = ?", partnerCompanyID).
		First(partnerCompany).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewAppErr(utils.ErrTypeNotfound, utils.ErrMsgNotfound, http.StatusNotFound, err)
		}
		return nil, err
	}
	return partnerCompany, nil
}

func (a *PartnerCompanyRepository) FindPartnerCompanyEmbedListByPartnerCompanyIDs(ctx context.Context, partnerCompanyIDs []string) (models.PartnerCompanyEmbedList, error) {
	partnerCompanies := []*models.PartnerCompanyEmbed{}
	tx := a.db.GetNewTxnOrContext(ctx)
	if err := tx.
		Where("id IN (?)", partnerCompanyIDs).
		Preload("PartnerCompanyBankAccount").
		Find(&partnerCompanies).Error; err != nil {
		return nil, err
	}
	return partnerCompanies, nil
}

func (a *PartnerCompanyRepository) CreatePartnerCompany(ctx context.Context, partnerCompany *models.PartnerCompany) error {
	return a.db.GetNewTxnOrContext(ctx).Create(partnerCompany).Error
}

func (a *PartnerCompanyRepository) CreatePartnerCompanyBankAccount(ctx context.Context, partnerCompanyBankAccount *models.PartnerCompanyBankAccount) error {
	return a.db.GetNewTxnOrContext(ctx).Create(partnerCompanyBankAccount).Error
}
