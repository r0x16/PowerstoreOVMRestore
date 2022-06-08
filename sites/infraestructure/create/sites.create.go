package create

import (
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui"
	"github.com/r0x16/PowerstoreOVMRestore/sites/application"
	"github.com/r0x16/PowerstoreOVMRestore/sites/infraestructure/create/view"
)

var mgui struct {
	wizard gui.Window
}

func CreateSiteAction() {

	if isOpened() {
		return
	}

	application.CreateSite()
	mgui.wizard = *view.NewCreateSiteWizard()
	mgui.wizard.Show()
}

// Verifies if a wizard create site is already opened
func isOpened() bool {
	return mgui.wizard != gui.Window{}
}
