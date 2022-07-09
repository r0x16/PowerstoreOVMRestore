package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
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
	accordeon := container.NewVBox()
	listView := &ListSiteView{
		container: accordeon,
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
	}
}

// Create a single site layout item
func (v *ListSiteView) siteLayout(site *model.Site) *fyne.Container {
	layout := container.NewVBox(
		widget.NewLabel(site.Name),
	)
	return layout
}

// Add a new widget to the site
func (v *ListSiteView) AddWidget(item fyne.CanvasObject, index int8) {
	v.SitesLayout[index].Add(item)
}
