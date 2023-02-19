package db

import (
	"fmt"
	"time"

	"github.com/MaxMoskalenko/pis-course-work/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB(cnf *config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s search_path=%s sslmode=disable TimeZone=UTC",
		cnf.Username, cnf.Password, cnf.Host, cnf.Port, cnf.Name, cnf.Schema,
	)

	dial := postgres.Open(dsn)

	db, err := gorm.Open(dial, &gorm.Config{
		NowFunc: time.Now().UTC,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(0)
	sqlDB.SetMaxOpenConns(100)

	return db, nil
}
