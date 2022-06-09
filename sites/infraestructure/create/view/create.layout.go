package view

import (
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui/wizard"
)

func NewCreateSiteWizard(wizardConfig wizard.WizardConfig, steps []wizard.WizardStep) *gui.Window {
	w := gui.NewWindow(wizardConfig.GetTitle())
	w.SetSize(config.Global.WIZZARD_WIDTH, config.Global.WIZZARD_HEIGHT)

	w.SetContent(initializeWizard(wizardConfig, steps))

	return w
}

func initializeWizard(config wizard.WizardConfig, steps []wizard.WizardStep) *wizard.Wizard {

	return wizard.NewWizard(config, steps)
}
