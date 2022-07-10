package sites

import (
	"github.com/r0x16/PowerstoreOVMRestore/sites/domain/model"
	"github.com/r0x16/PowerstoreOVMRestore/sites/infraestructure/list/view"
)

// A widget represent a submodule in the site list view.
type SiteWidget interface {
	SetSite(site *model.Site)
	SetDrawer(*view.WidgetDrawer)
	Run()
}

// A base implementation of the interface to general purpose widgets.
type BaseSiteWidget struct {
	Site   model.Site
	Drawer *view.WidgetDrawer
}

// SetSite sets the site to be used by the widget.
func (w *BaseSiteWidget) SetSite(site *model.Site) {
	w.Site = *site
}

// SetDrawer sets the drawer to be used by the widget.
func (w *BaseSiteWidget) SetDrawer(drawer *view.WidgetDrawer) {
	w.Drawer = drawer
}
