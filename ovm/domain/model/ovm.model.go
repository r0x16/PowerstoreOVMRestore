package model

import (
	"github.com/r0x16/PowerstoreOVMRestore/shared/domain"
	msite "github.com/r0x16/PowerstoreOVMRestore/sites/domain/model"
)

type Ovm struct {
	domain.Model
	Hostname string
	Username string
	Password string
	SiteID   int
	Site     *msite.Site
}
