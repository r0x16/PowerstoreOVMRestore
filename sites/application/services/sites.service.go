package services

import (
	"regexp"

	"github.com/r0x16/PowerstoreOVMRestore/sites/domain/model"
	"github.com/r0x16/PowerstoreOVMRestore/sites/domain/repository"
)

type SiteService struct {
	repository repository.SiteRepository
}

// Instantiates SiteService
func NewSiteService(repository repository.SiteRepository) *SiteService {
	service := &SiteService{}
	service.repository = repository
	return service
}

// Validates the name of the site
func (s *SiteService) ValidateName(name string) bool {
	validName := regexp.MustCompile(`^[a-zA-Z]+$`)
	return validName.MatchString(name)
}

// Get all sites from repository
func (s *SiteService) GetAll() ([]model.Site, int64, error) {
	return s.repository.GetAll()
}

// Get a site from the repository with the given name
func (s *SiteService) GetByName(name string) (*model.Site, int64, error) {
	return s.repository.GetByName(name)
}

// Persist site data into repository
func (s *SiteService) Store(site *model.Site) (int64, error) {
	return s.repository.Store(site)
}
