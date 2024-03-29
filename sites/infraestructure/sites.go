package infraestructure

import (
	"github.com/r0x16/PowerstoreOVMRestore/sites/infraestructure/create"
	"github.com/r0x16/PowerstoreOVMRestore/sites/infraestructure/list"
	"github.com/r0x16/PowerstoreOVMRestore/sites/infraestructure/view"
)

type SitesModule struct {
}

func NewSitesModule() {
	view.NewSitesLayout()
	view.SitesContainer.NewSiteToolbarAction(
		create.CreateSiteAction,
		view.SitesContainer.GetAddIcon(),
	)
	view.SitesContainer.CreateTitle()
	go list.ListSitesAction()
}
