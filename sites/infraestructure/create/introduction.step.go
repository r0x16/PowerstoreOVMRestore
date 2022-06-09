package create

import (
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui/wizard"
	"github.com/r0x16/PowerstoreOVMRestore/sites/infraestructure/create/view"
)

// IntroductionStep implement a wizard step based on default BasewizardStep
type IntroductionStep struct {
	wizard.BaseWizardStep
}

// Creates and returns the instroduction step pane
func NewIntroductionStep() *IntroductionStep {
	step := &IntroductionStep{}
	step.Content = view.NewIntroductionStepLayout()
	step.Caption = ""
	step.Title = config.Lang.Module.CreateSitesWizard.IntroductionSectionTitle
	return step
}
