package create

import (
	"fyne.io/fyne/v2"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui/wizard"
	"github.com/r0x16/PowerstoreOVMRestore/sites/application"
	"github.com/r0x16/PowerstoreOVMRestore/sites/infraestructure/create/view"
	"github.com/r0x16/PowerstoreOVMRestore/sites/infraestructure/repository"
)

// Defines all the items involved in a new site creation
type createSiteInstance struct {
	wizard *gui.Window
	site   *application.SiteCreationData
}

// Stores all wizard status for this instance
var createInstance *createSiteInstance

// Entrypoint called from main window
func CreateSiteAction() {

	if isOpened() {
		return
	}

	createInstance = &createSiteInstance{}
	createInstance.site = application.NewCreateSite(
		repository.NewSiteRepository(),
	)

	wc := &wizardConfig{}
	wc.Title = config.Lang.Module.CreateSitesWizard.WizardTitle

	createInstance.wizard = view.NewCreateSiteWizard(wc, []wizard.WizardStep{
		NewIntroductionStep(),
		NewNameStep(),
		NewSummaryStep(),
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
	_, err := createInstance.site.StoreSite()
	if err != nil {
		return view.NewErrorSiteCreationLayout(err)
	}
	return view.NewSuccessSiteCreationLayout()
}

// -- END of Wizard configuration struct
