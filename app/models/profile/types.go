package profile

import (
	"time"

	"github.com/google/uuid"
)

type ProfileRole string

const (
	User       ProfileRole = "user"
	Admin      ProfileRole = "admin"
	SuperAdmin ProfileRole = "superadmin"
)

type Profile struct {
	ID           uuid.UUID   `gorm:"column:id"`
	Login        string      `gorm:"column:login"`
	PasswordHash string      `gorm:"column:password_hash"`
	ProfileRole  ProfileRole `gorm:"column:profile_role"`
	CreatedAt    time.Time   `gorm:"column:created_at"`
	UpdatedAt    time.Time   `gorm:"column:updated_at"`
}
