package repository

import (
	"github.com/r0x16/PowerstoreOVMRestore/ovm/domain/model"
	msite "github.com/r0x16/PowerstoreOVMRestore/sites/domain/model"
)

type SiteRepository interface {
	GetBySite(msite.Site) (model.Ovm, int64, error)
	Store(model.Ovm) (int64, error)
}
