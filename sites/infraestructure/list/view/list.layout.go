package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/r0x16/PowerstoreOVMRestore/sites/domain/model"
	"github.com/r0x16/PowerstoreOVMRestore/sites/infraestructure/view"
)

type ListSiteView struct {
	container   *widget.Accordion
	SitesLayout []*fyne.Container
}

// Initializes the container view for listing the sites
func NewSiteListLayout() *ListSiteView {
	accordeon := widget.NewAccordion()
	listView := &ListSiteView{
		container: accordeon,
	}
	view.SitesContainer.RefreshSiteListLayout(container.NewVBox(listView.container))
	return listView
}

// Draw and references all sites layouts
func (v *ListSiteView) SetSites(sites []model.Site) {
	v.SitesLayout = make([]*fyne.Container, len(sites))
	for i, site := range sites {
		layout := v.siteLayout(&site)
		v.SitesLayout[i] = layout.Detail.(*fyne.Container)
		v.container.Append(layout)
	}
}

// Create a single site layout item
func (v *ListSiteView) siteLayout(site *model.Site) *widget.AccordionItem {
	return widget.NewAccordionItem(site.Name, container.NewVBox())
}

// Add a new widget to the site
func (v *ListSiteView) AddWidget(item fyne.CanvasObject, index int8) {
	v.SitesLayout[index].Add(item)
}
