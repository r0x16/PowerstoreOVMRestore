package repository

import "github.com/r0x16/PowerstoreOVMRestore/sites/domain/model"

type SiteRepository interface {
	GetByName(string) (*model.Site, int64, error)
}
