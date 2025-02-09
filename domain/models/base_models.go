package models

import (
	"time"

	"gorm.io/gorm"
)

type HardModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SoftModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt
}
