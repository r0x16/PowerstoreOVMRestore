package application

import (
	"github.com/r0x16/PowerstoreOVMRestore/ovm/application/services"
	"github.com/r0x16/PowerstoreOVMRestore/ovm/domain/model"
	"github.com/r0x16/PowerstoreOVMRestore/ovm/domain/repository"

	ms "github.com/r0x16/PowerstoreOVMRestore/sites/domain/model" // ms is for model site
)

// Represents the widget data provider
type OvmWidgetProvider struct {
	ovmService *services.OvmService
}

// Creates a new widget data provider and initializes service
func NewWidgetOvm(repository repository.OvmRepository) *OvmWidgetProvider {
	data := &OvmWidgetProvider{
		ovmService: services.NewOvmService(repository),
	}
	return data
}

// Returns the associated ovm for the given site
func (p *OvmWidgetProvider) GetOvmFromSite(site *ms.Site) (*model.Ovm, int64, error) {
	ovm, length, err := p.ovmService.GetBySite(site)
	if err != nil {
		return nil, 0, err
	}
	return ovm, length, nil
}
