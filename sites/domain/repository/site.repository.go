package repository

import "github.com/r0x16/PowerstoreOVMRestore/sites/domain/model"

type SiteRepository interface {
	GetAll() (*[]model.Site, int64, error)
	GetByName(string) (*model.Site, int64, error)
	Store(*model.Site) (int64, error)
}
