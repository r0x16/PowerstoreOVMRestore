package ovm

import (
	"github.com/r0x16/PowerstoreOVMRestore/sites/infraestructure/sites"
)

type OvmWidget struct {
	sites.BaseSiteWidget
}

// Sites Widget implementation
var _ sites.SiteWidget = &OvmWidget{}

func (w *OvmWidget) Draw() {
}
