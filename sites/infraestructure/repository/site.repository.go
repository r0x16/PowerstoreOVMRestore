package repository

import (
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/database"
	"github.com/r0x16/PowerstoreOVMRestore/sites/domain/model"
	i "github.com/r0x16/PowerstoreOVMRestore/sites/domain/repository"
	"gorm.io/gorm"
)

type SiteRepositoryGorm struct {
	db *gorm.DB
}

var _ i.SiteRepository = SiteRepositoryGorm{}

func NewSiteRepository() *SiteRepositoryGorm {
	repo := &SiteRepositoryGorm{
		db: database.Connection,
	}
	repo.db.AutoMigrate(&model.Site{})
	return repo
}

// getByName implements repository.SiteRepository
func (s SiteRepositoryGorm) GetByName(name string) *model.Site {
	site := model.Site{}
	s.db.Where("name = ?", name).First(&site)
	return &site
}
