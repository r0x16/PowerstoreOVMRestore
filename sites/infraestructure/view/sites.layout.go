package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui"
)

type SitesLayout struct {
	Toolbar   *widget.Toolbar
	boxLayout *fyne.Container
}

var SitesContainer *SitesLayout

func NewSitesLayout() {
	sc := &SitesLayout{}
	sc.boxLayout = sc.createBoxLayout()
	sc.Toolbar = sc.createToolbar()
	SitesContainer = sc

	gui.MainContainer.Body.Trailing = SitesContainer.boxLayout
	gui.MainContainer.Body.Refresh()
}

func (sc *SitesLayout) createBoxLayout() *fyne.Container {
	return container.NewVBox()
}

func (sc *SitesLayout) createToolbar() *widget.Toolbar {
	tb := widget.NewToolbar()
	sc.boxLayout.Add(tb)
	return tb
}

func (sc *SitesLayout) NewSiteToolbarAction(action func()) {
	sc.Toolbar.Append(
		widget.NewToolbarAction(theme.ContentAddIcon(), action),
	)
}
