package ovm

import (
	"github.com/r0x16/PowerstoreOVMRestore/sites/domain/model"
	"github.com/r0x16/PowerstoreOVMRestore/sites/infraestructure/sites"
)

type OvmWidget struct {
	sites.BaseSiteWidget
}

// Sites Widget implementation
var _ sites.SiteWidget = &OvmWidget{}

func NewOvmWidget(site model.Site, index int8) *OvmWidget {
	widget := &OvmWidget{}
	widget.SetSite(site)
	widget.SetIndex(index)

	return widget
}

func (w *OvmWidget) Draw() {
}
