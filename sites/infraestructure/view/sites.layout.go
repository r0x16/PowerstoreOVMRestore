package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui"
)

type SitesLayout struct {
	Toolbar       *widget.Toolbar
	SiteListIndex int
	boxLayout     *fyne.Container
}

var SitesContainer *SitesLayout

func NewSitesLayout() {
	sc := &SitesLayout{}
	sc.boxLayout = sc.createBoxLayout()
	sc.Toolbar = sc.createToolbar()
	sc.SiteListIndex = sc.createSiteListContainer()
	SitesContainer = sc

	gui.MainContainer.Body.Trailing = SitesContainer.boxLayout
	gui.MainContainer.Body.Refresh()
}

func (sc *SitesLayout) createBoxLayout() *fyne.Container {
	return container.NewVBox()
}

// START -- ListContainer
// Create a new object for contains the list of sites
// returns the index of that object to reference list container location
func (sc *SitesLayout) createSiteListContainer() int {
	index := len(sc.boxLayout.Objects) // container size is index for new added item

	sc.boxLayout.Add(container.NewVBox()) // <-- This item

	return index
}

// Overwrites list container to refresh sites list
func (sc *SitesLayout) RefreshSiteListLayout(list *fyne.Container) {
	sc.boxLayout.Objects[sc.SiteListIndex] = list
}

// -- END ListContainer

// START -- Toolbar
// Creates a new empty toolbar in the top of site container
func (sc *SitesLayout) createToolbar() *widget.Toolbar {
	tb := widget.NewToolbar()
	sc.boxLayout.Add(tb)
	return tb
}

// Adds a new item to the toolbar
func (sc *SitesLayout) NewSiteToolbarAction(action func(), icon fyne.Resource) {
	sc.Toolbar.Append(
		widget.NewToolbarAction(icon, action),
	)
}

// Returns the add icon for the toolbar
func (sc *SitesLayout) GetAddIcon() fyne.Resource {
	return theme.ContentAddIcon()
}

// -- END Toolbar
