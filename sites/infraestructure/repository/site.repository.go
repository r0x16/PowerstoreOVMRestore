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

var _ i.SiteRepository = &SiteRepositoryGorm{}

func NewSiteRepository() *SiteRepositoryGorm {
	repo := &SiteRepositoryGorm{
		db: database.Connection,
	}
	repo.db.AutoMigrate(&model.Site{})
	return repo
}

// GetAll implements repository.SiteRepository
func (s *SiteRepositoryGorm) GetAll() (*[]model.Site, int64, error) {
	var sites *[]model.Site
	result := s.db.Find(sites)
	return sites, result.RowsAffected, result.Error
}

// getByName implements repository.SiteRepository
func (s *SiteRepositoryGorm) GetByName(name string) (*model.Site, int64, error) {
	site := model.Site{}
	result := s.db.Where("name = ?", name).Limit(1).Find(&site)
	return &site, result.RowsAffected, result.Error
}

// Store implements repository.SiteRepository
func (s *SiteRepositoryGorm) Store(site *model.Site) (int64, error) {
	result := s.db.Create(site)
	return result.RowsAffected, result.Error
}
