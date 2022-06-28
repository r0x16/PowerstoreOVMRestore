package repository

import (
	"github.com/r0x16/PowerstoreOVMRestore/ovm/domain/model"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/database"

	i "github.com/r0x16/PowerstoreOVMRestore/ovm/domain/repository" // i is for interface
	ms "github.com/r0x16/PowerstoreOVMRestore/sites/domain/model"   // ms is for model site

	"gorm.io/gorm"
)

type OvmRepositoryGorm struct {
	db *gorm.DB
}

var _ i.OvmRepository = &OvmRepositoryGorm{}

// Creates a new GORM Ovm Repository
func NewOvmRepository() *OvmRepositoryGorm {
	repo := &OvmRepositoryGorm{
		db: database.Connection,
	}
	repo.db.AutoMigrate(&model.Ovm{})
	return repo
}

// Get Ovm from database based on site
func (o *OvmRepositoryGorm) GetBySite(site *ms.Site) (*model.Ovm, int64, error) {
	ovm := &model.Ovm{}
	result := o.db.Where("site_id = ?", site.ID).Limit(1).Find(ovm)
	return ovm, result.RowsAffected, result.Error
}

// Store a new ovm in database
func (o *OvmRepositoryGorm) Store(ovm *model.Ovm) (int64, error) {
	result := o.db.Create(&ovm)
	return result.RowsAffected, result.Error
}
