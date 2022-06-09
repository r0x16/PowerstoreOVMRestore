package create

import (
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui/wizard"
	"github.com/r0x16/PowerstoreOVMRestore/sites/infraestructure/create/view"
)

// All the data required for create a new site name descriptor
type NameStepData interface {
	GetName() string
	GetDescription() string
}

// NameStep implements a wizard step based on default BasewizardStep
type NameStep struct {
	wizard.BaseWizardStep
	Data NameStepData
}

// Creates and returns the set nameand description step pane
func NewNameStep() *NameStep {
	step := &NameStep{}
	step.Content, step.Data = view.NewNameStepLayout()
	step.Caption = config.Lang.Module.CreateSitesWizard.NameCaption
	step.Title = config.Lang.Module.CreateSitesWizard.NameSectionTitle
	return step
}

// Validates data before sending to the next step
func (s *NameStep) OnValidate() bool {
	return false
}
