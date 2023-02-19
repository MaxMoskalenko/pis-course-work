package profile

import "github.com/google/uuid"

type Interface interface {
	Create(p *Profile) error
	UpdateByID(id uuid.UUID, p *Profile) error
	GetByID(id uuid.UUID) (*Profile, error)
}
