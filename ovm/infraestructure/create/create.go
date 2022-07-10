package create

import (
	"fyne.io/fyne/v2"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui/wizard"

	ms "github.com/r0x16/PowerstoreOVMRestore/sites/domain/model"
)

type createOvmInstance struct {
	wizard *gui.Window
	Site   *ms.Site
}

var createInstance *createOvmInstance

func CreateOvmAction(site *ms.Site) {
	if isOpened() {
		return
	}
	createInstance = &createOvmInstance{
		Site: site,
	}

	wc := &createOvmWizardConfig{}
	wc.Title = config.Lang.Module.CreateOvmWizard.WizardTitle

	// Need almost 1 Step to run the wizard -- TODO: improve this
	//createInstance.wizard = view.NewCreateSiteWizard(wc, []wizard.WizardStep{})
	//createInstance.wizard.SetOnClose(wc.OnClose)
	//createInstance.wizard.Show()
}

func isOpened() bool {
	return createInstance != nil
}

type createOvmWizardConfig struct {
	wizard.BaseWizardConfig
}

// Callback called on 3 cases:
// 1. Cancel Button clicked
// 2. Close button clicked after Finished
// 3. Window instance closed
func (c *createOvmWizardConfig) OnClose() {
	createInstance.wizard.Close()
	createInstance = nil
}

// Stores new sites created and returns finish view
func (c *createOvmWizardConfig) OnFinish() *fyne.Container {
	return nil
}
