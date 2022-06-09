package create

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui/wizard"
	"github.com/r0x16/PowerstoreOVMRestore/sites/application"
	"github.com/r0x16/PowerstoreOVMRestore/sites/domain/model"
	"github.com/r0x16/PowerstoreOVMRestore/sites/infraestructure/create/view"
)

// Defines all the items involved in a new site creation
type createSiteInstance struct {
	wizard *gui.Window
	site   *model.Site
}

// Stores all wizard status for this instance
var createInstance *createSiteInstance

// Entrypoint called from main window
func CreateSiteAction() {

	if isOpened() {
		return
	}

	createInstance = &createSiteInstance{}

	createInstance.site = application.CreateSite()

	wc := &wizardConfig{}
	wc.Title = config.Lang.Module.CreateSitesWizard.WizardTitle

	createInstance.wizard = view.NewCreateSiteWizard(wc, []wizard.WizardStep{
		NewIntroductionStep(),
	})
	createInstance.wizard.SetOnClose(wc.OnClose)
	createInstance.wizard.Show()
}

// Verifies if a wizard create site is already opened, only one is allowed
func isOpened() bool {
	return createInstance != nil
}

// -- START of Wizard configuration struct --
type wizardConfig struct {
	wizard.BaseWizardConfig
}

// Callback called on 3 cases:
// 1. Cancel Button clicked
// 2. Close button clicked after Finished
// 3. Window instance closed
func (c *wizardConfig) OnClose() {
	createInstance.wizard.Close()
	createInstance = nil
}

// Stores new sites created and returns finish view
func (c *wizardConfig) OnFinish() *fyne.Container {
	return container.NewVBox(widget.NewLabel("Finished"))
}

// -- END of Wizard configuration struct
