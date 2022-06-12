package application

import (
	"github.com/r0x16/PowerstoreOVMRestore/sites/application/services"
	"github.com/r0x16/PowerstoreOVMRestore/sites/domain/model"
	"github.com/r0x16/PowerstoreOVMRestore/sites/domain/repository"
)

type SiteListProvider struct {
	siteService *services.SiteService
}

// Creates a new site list provider and initializes service
func NewListSite(siteRepository repository.SiteRepository) *SiteListProvider {
	data := &SiteListProvider{
		siteService: services.NewSiteService(siteRepository),
	}
	return data
}

// Returns all the sites stored on the repository
func (p *SiteListProvider) GetAllSites() ([]model.Site, int64, error) {
	sites, length, err := p.siteService.GetAll()
	if err != nil {
		return nil, 0, err
	}
	return sites, length, nil
}
