package infraestructure

import (
	"github.com/r0x16/PowerstoreOVMRestore/sites/infraestructure/view"
)

type SitesModule struct {
}

func NewSitesModule() {
	view.NewSitesLayout()
	view.SitesContainer.NewSiteToolbarAction(newSiteAction)
}

func newSiteAction() {

}