package create

import (
	"github.com/r0x16/PowerstoreOVMRestore/ovm/infraestructure/create/view"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui/wizard"
)

// AccessDataStep is the step of the OVM creation wizard
// that allows the user to type the access data of the OVM Manager.
type AccessDataStep struct {
	wizard.BaseWizardStep
}

//	Creates and returns the ovm set access data step pane
func NewAccessDataStep(content *view.AccessDataStepLayout) *AccessDataStep {
	step := &AccessDataStep{}
	step.Content = content.GetContainer()
	step.Caption = config.Lang.Module.CreateOvmWizard.AccessDataCaption
	step.Title = config.Lang.Module.CreateOvmWizard.AccessDataSectionTitle
	return step
}
