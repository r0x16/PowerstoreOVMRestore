package ovm

import (
	"github.com/r0x16/PowerstoreOVMRestore/ovm/infraestructure"
	"github.com/r0x16/PowerstoreOVMRestore/sites/infraestructure/sites"
)

type OvmWidget struct {
	sites.BaseSiteWidget
	Module *infraestructure.OvmManagerModule
}

// Sites Widget implementation
var _ sites.SiteWidget = &OvmWidget{}

// called by site module when widget is started (called as goroutine)
func (w *OvmWidget) Run() {
	w.Module = infraestructure.NewOvmManagerModule(&w.Site)
	w.Module.RunWidget(w.Drawer)
}
