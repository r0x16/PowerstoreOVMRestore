package services

import (
	"github.com/r0x16/PowerstoreOVMRestore/ovm/domain/model"
	"github.com/r0x16/PowerstoreOVMRestore/ovm/domain/repository"

	ms "github.com/r0x16/PowerstoreOVMRestore/sites/domain/model"
)

type OvmService struct {
	repository repository.OvmRepository
}

// Instantiates OvmService
func NewOvmService(repository repository.OvmRepository) *OvmService {
	service := &OvmService{}
	service.repository = repository
	return service
}

// Get ovm from repository with the given site
func (s *OvmService) GetBySite(site *ms.Site) (*model.Ovm, int64, error) {
	return s.repository.GetBySite(site)
}

// Persist ovm data into repository
func (s *OvmService) Store(ovm *model.Ovm) (int64, error) {
	return s.repository.Store(ovm)
}
