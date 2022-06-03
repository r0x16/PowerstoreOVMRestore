package infraestructure

import (
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui"
	"github.com/r0x16/PowerstoreOVMRestore/sites/application"
	"github.com/r0x16/PowerstoreOVMRestore/sites/infraestructure/view"
)

var mgui struct {
	wizard gui.Window
}

func createSiteAction() {
	application.CreateSite()
	mgui.wizard = *view.NewCreateSiteWizard()
	mgui.wizard.Show()
}
