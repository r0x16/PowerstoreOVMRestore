package create

import (
	"fyne.io/fyne/v2"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui/wizard"
	"github.com/r0x16/PowerstoreOVMRestore/sites/infraestructure/create/view"
)

// SummaryStep implements a wizard step based on default BasewizardStep
type SummaryStep struct {
	wizard.BaseWizardStep
}

// Create and returns the site creation summary step pane
func NewSummaryStep() *SummaryStep {
	step := &SummaryStep{}
	step.Caption = config.Lang.Module.CreateSitesWizard.SummaryCaption
	step.Title = config.Lang.Module.CreateSitesWizard.SummarySectionTitle
	return step
}

// Load the summary layout based on the actual site content
func (s *SummaryStep) OnLoad() *fyne.Container {
	s.Content = view.NewSummaryStepLayout(createInstance.site)
	return s.Content
}
