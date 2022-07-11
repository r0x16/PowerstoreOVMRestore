package create

import (
	"github.com/r0x16/PowerstoreOVMRestore/ovm/infraestructure/create/view"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui/wizard"
)

// IntroductionStep implement a wizard step based on default BasewizardStep
type IntroductionStep struct {
	wizard.BaseWizardStep
}

// Creates and returns the introduction step pane
func NewIntroductionStep(content *view.IntroductionStepLayout) *IntroductionStep {
	step := &IntroductionStep{}
	step.Content = content.GetContainer()
	step.Caption = ""
	step.Title = config.Lang.Module.CreateOvmWizard.IntroductionSectionTitle
	return step
}
