package view

import (
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui"
)

var CreateSiteWizardWindow *gui.Window

func NewCreateSiteWizard() *gui.Window {
	w := gui.NewWindow("Create Site")
	w.SetSize(config.Global.WIZZARD_WIDTH, config.Global.WIZZARD_HEIGHT)
	CreateSiteWizardWindow = w
	return w
}
