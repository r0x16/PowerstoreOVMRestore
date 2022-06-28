package ovm

import (
	"github.com/r0x16/PowerstoreOVMRestore/sites/infraestructure/sites"
)

type OvmWidget struct {
	sites.BaseSiteWidget
}

// Sites Widget implementation
var _ sites.SiteWidget = &OvmWidget{}

// called by site module when widget is started (called as goroutine)
func (w *OvmWidget) Run() {

}

// called by site module when widget drawing is required (called as goroutine after Run())
func (w *OvmWidget) Draw() {

}
