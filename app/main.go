package main

import (
	"fmt"

	"github.com/MaxMoskalenko/pis-course-work/handlers"
	"github.com/MaxMoskalenko/pis-course-work/models/profile"
	"github.com/MaxMoskalenko/pis-course-work/pkg/config"
	"github.com/MaxMoskalenko/pis-course-work/pkg/ctx"
	"github.com/MaxMoskalenko/pis-course-work/pkg/db"
	"github.com/MaxMoskalenko/pis-course-work/pkg/hash"
	"github.com/google/uuid"
)

var (
	configFilePath = "./config/.env"
)

func main() {
	ctx, err := setupApp()

	if err != nil {
		fmt.Printf("Failed to initialize app: %+v\n", err)
		return
	}

	p := &profile.Profile{
		ID:           uuid.New(),
		Login:        uuid.NewString(),
		PasswordHash: hash.HashString("password"),
		ProfileRole:  profile.User,
	}

	err = ctx.ProfileInterface.Create(p)
	if err != nil {
		fmt.Printf("Failed to create profile: %+v\n", err)
		return
	}
	fmt.Printf("Profile created: %+v\n", p)

	err = ctx.ProfileInterface.UpdateByID(p.ID, &profile.Profile{ProfileRole: profile.SuperAdmin})
	if err != nil {
		fmt.Printf("Failed to update profile: %+v\n", err)
		return
	}

	p, err = ctx.ProfileInterface.GetByID(p.ID)
	if err != nil {
		fmt.Printf("Failed to get profile: %+v\n", err)
		return
	}
	fmt.Printf("Updated profile: %+v\n", p)
}

func setupApp() (*ctx.Ctx, error) {
	cnf, err := config.Load(configFilePath)

	if err != nil {
		return nil, err
	}

	db, err := db.SetupDB(&cnf.Database)
	if err != nil {
		return nil, err
	}

	ctx, err := handlers.NewApplication(cnf, db)
	if err != nil {
		return nil, err
	}

	return ctx, nil
}
