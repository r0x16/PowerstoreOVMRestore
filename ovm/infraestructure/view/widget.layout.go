package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"

	vs "github.com/r0x16/PowerstoreOVMRestore/sites/infraestructure/list/view"
)

type OvmLayout struct {
	drawer *vs.WidgetDrawer
}

// Creates a new view layout
func NewOvmLayout(drawer *vs.WidgetDrawer) *OvmLayout {
	return &OvmLayout{
		drawer: drawer,
	}
}

// Draw the empty widget, with a message and a button to setup the OVM
func (ol *OvmLayout) DrawEmpty(linkOvm func()) {
	ol.drawer.Draw(
		container.NewVBox(
			getTitle(),
			GetEmptyWidgetContainer(linkOvm),
			widget.NewSeparator(),
		))
}

// Creates the title of the OVM manager widget
func getTitle() *fyne.Container {
	logo := getOvmLogo()
	title := getOvmTitle()
	c := container.NewHBox(logo, title)
	return c
}

// Creates the logo of the OVM manager widget
func getOvmLogo() *canvas.Image {
	logo := canvas.NewImageFromFile("assets/ora-logo_small.png")
	logo.FillMode = canvas.ImageFillOriginal
	logo.SetMinSize(fyne.NewSize(25, 25))
	return logo
}

// Creates the title text of the OVM manager widget
func getOvmTitle() *canvas.Text {
	title := canvas.NewText(" | "+config.Lang.Module.OvmWidget.OvmManager, theme.ForegroundColor())
	title.TextSize = 12
	return title
}
