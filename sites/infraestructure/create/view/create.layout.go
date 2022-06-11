package view

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui/wizard"
)

// Creates a window object and initialize their content to create a new site
func NewCreateSiteWizard(wizardConfig wizard.WizardConfig, steps []wizard.WizardStep) *gui.Window {
	w := gui.NewWindow(wizardConfig.GetTitle())
	w.SetSize(config.Global.WIZZARD_WIDTH, config.Global.WIZZARD_HEIGHT)

	w.SetContent(initializeWizard(wizardConfig, steps))

	return w
}

// Initialize and returns a new wizard pane
func initializeWizard(config wizard.WizardConfig, steps []wizard.WizardStep) *wizard.Wizard {
	return wizard.NewWizard(config, steps)
}

// Creates the error layout when and store error ocurred
func NewErrorSiteCreationLayout(err error) *fyne.Container {
	errorGrid := widget.NewTextGrid()
	errorGrid.SetText(err.Error())
	return container.NewVBox(
		canvas.NewText(config.Lang.Module.CreateSitesWizard.Errors.NameStoreError, color.RGBA{R: 255, A: 255}),
		errorGrid,
	)
}

// Creates the layout when store a site successfully
func NewSuccessSiteCreationLayout() *fyne.Container {
	msg := canvas.NewText(
		config.Lang.Module.CreateSitesWizard.OnFinishSuccess,
		color.RGBA{G: 200, A: 255},
	)
	msg.Alignment = fyne.TextAlignCenter
	return container.NewVBox(
		widget.NewLabel(""),
		msg,
		widget.NewIcon(theme.ConfirmIcon()),
	)
}
