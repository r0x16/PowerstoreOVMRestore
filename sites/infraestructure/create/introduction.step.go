package create

import (
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui/wizard"
	"github.com/r0x16/PowerstoreOVMRestore/sites/infraestructure/create/view"
)

type IntroductionStep struct {
	wizard.BaseWizardStep
}

func NewIntroductionStep() *IntroductionStep {
	step := &IntroductionStep{}
	step.Content = view.NewIntroductionStepLayout()
	step.Caption = ""
	step.Title = config.Lang.Module.CreateSitesWizard.IntroductionSectionTitle
	return step
}
