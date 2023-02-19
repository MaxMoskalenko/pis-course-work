package profile

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DBAdapter struct {
	db *gorm.DB
}

func NewDBAdapter(gormDb *gorm.DB) *DBAdapter {
	return &DBAdapter{
		db: gormDb.WithContext(context.Background()),
	}
}

func (Profile) TableName() string {
	return "profiles"
}

func (g *DBAdapter) Create(p *Profile) error {
	return g.db.Model(&Profile{}).Create(&p).Error
}

func (g *DBAdapter) UpdateByID(id uuid.UUID, p *Profile) error {
	return g.db.Model(&Profile{}).Where("id = ?", id).Updates(&p).Error
}

func (g *DBAdapter) GetByID(id uuid.UUID) (*Profile, error) {
	var p Profile
	err := g.db.Where("id = ?", id).Find(&p).Error
	return &p, err
}
