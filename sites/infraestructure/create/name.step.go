package create

import (
	"strings"

	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui/wizard"
	"github.com/r0x16/PowerstoreOVMRestore/sites/application"
	"github.com/r0x16/PowerstoreOVMRestore/sites/infraestructure/create/view"
)

// All the data required for create a new site name descriptor
type NameStepData interface {
	GetName() string
	SetNameError(string)

	GetDescription() string

	ClearErrors()
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
	s.Data.ClearErrors()
	err := createInstance.site.ValidateName(s.Data.GetName())
	if err > 0 {
		switch err {
		case application.ERROR_SITE_EMPTY:
			s.Data.SetNameError(config.Lang.Module.CreateSitesWizard.Errors.NameEmpty)
		case application.ERROR_SITE_NAMEFORMAT:
			s.Data.SetNameError(config.Lang.Module.CreateSitesWizard.Errors.NameFormat)
		case application.ERROR_SITE_EXISTS:
			msg := strings.ReplaceAll(config.Lang.Module.CreateSitesWizard.Errors.NameExists, "{name}", s.Data.GetName())
			s.Data.SetNameError(msg)
		}
		return false
	}
	return true
}

// Set data on globals before leave the step
func (s *NameStep) OnLeave() {
	createInstance.site.Name = s.Data.GetName()
	createInstance.site.Description = s.Data.GetDescription()
}
