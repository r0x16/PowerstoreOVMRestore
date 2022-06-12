package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/r0x16/PowerstoreOVMRestore/sites/domain/model"
	"github.com/r0x16/PowerstoreOVMRestore/sites/infraestructure/view"
)

type ListSiteView struct {
	container   *fyne.Container
	SitesLayout []*fyne.Container
}

// Initializes the container view for listing the sites
func NewSiteListLayout() *ListSiteView {
	listView := &ListSiteView{
		container: container.NewVBox(),
	}
	view.SitesContainer.RefreshSiteListLayout(listView.container)
	return listView
}

// Draw and references all sites layouts
func (v *ListSiteView) SetSites(sites []model.Site) {
	v.SitesLayout = make([]*fyne.Container, len(sites))
	for i, site := range sites {
		layout := v.siteLayout(&site)
		v.SitesLayout[i] = layout
		v.container.Add(layout)
		v.container.Add(widget.NewSeparator())
	}
}

// Create a single site layout item
func (v *ListSiteView) siteLayout(site *model.Site) *fyne.Container {
	name := canvas.NewText(site.Name, theme.ForegroundColor())
	return container.NewVBox(name)
}
