package ctx

import (
	"github.com/MaxMoskalenko/pis-course-work/models/profile"
	"github.com/MaxMoskalenko/pis-course-work/pkg/config"
	"gorm.io/gorm"
)

type Ctx struct {
	Config *config.Config
	DB     *gorm.DB

	ProfileInterface profile.Interface
}
