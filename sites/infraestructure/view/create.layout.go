package view

import "github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui"

var CreateSiteWizardWindow *gui.Window

func NewCreateSiteWizard() *gui.Window {
	w := gui.NewWindow("Create Site")
	w.SetSize(800, 600)
	CreateSiteWizardWindow = w
	return w
}
