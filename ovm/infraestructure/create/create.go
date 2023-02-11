package create

import (
	"fyne.io/fyne/v2"
	"github.com/r0x16/PowerstoreOVMRestore/ovm/infraestructure/create/view"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui/wizard"

	ms "github.com/r0x16/PowerstoreOVMRestore/sites/domain/model"
)

// Stores the data for the wizard steps
type createOvmInstance struct {
	wizard *gui.Window
	Site   *ms.Site
}

var createInstance *createOvmInstance

// This action creates a new wizard window for link a new ovm manager with a site

func CreateOvmAction(site *ms.Site) {
	if isOpened() {
		return
	}
	createInstance = &createOvmInstance{
		Site: site,
	}

	wc := &createOvmWizardConfig{}
	wc.Title = config.Lang.Module.CreateOvmWizard.WizardTitle

	createInstance.wizard = view.NewCreateOvmWizard(wc, []wizard.WizardStep{
		NewIntroductionStep(view.NewIntroductionStepLayout(site.Name)),
		NewAccessDataStep(view.NewAccessDataStepLayout()),
	})
	createInstance.wizard.SetOnClose(wc.OnClose)
	createInstance.wizard.Show()
}

// Check if wizard window is opened
func isOpened() bool {
	return createInstance != nil
}

/*
	Holds all features needed for wizard to work
*/
type createOvmWizardConfig struct {
	wizard.BaseWizardConfig
}

/*
	Callback called on 3 cases:
		1. Cancel Button clicked
		2. Close button clicked after Finished
		3. Window instance closed
*/
func (c *createOvmWizardConfig) OnClose() {
	createInstance.wizard.Close()
	createInstance = nil
}

/*
	Stores new Ovm manager configured and returns finish view
*/
func (c *createOvmWizardConfig) OnFinish() *fyne.Container {
	return nil
}
