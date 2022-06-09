package application

import (
	"github.com/r0x16/PowerstoreOVMRestore/sites/application/services"
	"github.com/r0x16/PowerstoreOVMRestore/sites/domain/repository"
)

type SiteCreationData struct {
	Name        string
	Description string

	siteService *services.SiteService
}

const (
	ERROR_SITE_EMPTY      int = 1
	ERROR_SITE_NAMEFORMAT int = 2
	ERROR_SITE_EXISTS     int = 3
)

// Creates a new site model instance
func NewCreateSite(siteRepository repository.SiteRepository) *SiteCreationData {
	data := &SiteCreationData{
		siteService: services.NewSiteService(siteRepository),
	}
	return data
}

// This method checks if the site has a valid name or if already exists
func (d *SiteCreationData) ValidateName(name string) int {

	if len(name) == 0 {
		return ERROR_SITE_EMPTY
	}

	if !d.siteService.ValidateName(name) {
		return ERROR_SITE_NAMEFORMAT
	}

	_, count, _ := d.siteService.GetByName(name)

	if count != 0 {
		return ERROR_SITE_EXISTS
	}

	return 0
}
