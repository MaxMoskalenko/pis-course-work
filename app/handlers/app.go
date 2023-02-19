package handlers

import (
	"github.com/MaxMoskalenko/pis-course-work/models/profile"
	"github.com/MaxMoskalenko/pis-course-work/pkg/config"
	"github.com/MaxMoskalenko/pis-course-work/pkg/ctx"
	"gorm.io/gorm"
)

func NewApplication(cnf *config.Config, db *gorm.DB) (*ctx.Ctx, error) {
	profileInterface := profile.NewDBAdapter(db)

	app := &ctx.Ctx{
		Config:           cnf,
		DB:               db,
		ProfileInterface: profileInterface,
	}

	return app, nil
}
