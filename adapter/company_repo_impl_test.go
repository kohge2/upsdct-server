package adapter

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kohge2/upsdct-server/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestFindByCompanyID(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer sqlDB.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	testDB := database.NewDB(gormDB)
	repo := NewCompanyRepository(testDB)

	t.Run("正常系", func(t *testing.T) {
		companyID := "12345"
		companyName := "Test Company"

		rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(companyID, companyName)
		query := regexp.QuoteMeta("SELECT * FROM `companies` WHERE id = ? AND `companies`.`deleted_at` IS NULL ORDER BY `companies`.`id` LIMIT ?")
		mock.ExpectQuery(query).
			WithArgs(companyID, 1).
			WillReturnRows(rows)

		ctx := context.Background()
		company, err := repo.FindByCompanyID(ctx, companyID)

		assert.NoError(t, err)
		assert.NotNil(t, company)
		assert.Equal(t, companyName, company.Name)
	})
	t.Run("正常系_NotfoundErrorを返す", func(t *testing.T) {
		companyID := "cp99999"

		query := regexp.QuoteMeta("SELECT * FROM `companies` WHERE id = ? AND `companies`.`deleted_at` IS NULL ORDER BY `companies`.`id` LIMIT ?")
		mock.ExpectQuery(query).
			WithArgs(companyID).
			WillReturnError(gorm.ErrRecordNotFound)

		ctx := context.Background()
		_, err = repo.FindByCompanyID(ctx, "cp99999")

		assert.Error(t, err)
	})
}
