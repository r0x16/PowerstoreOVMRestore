package widget

import (
	"github.com/r0x16/PowerstoreOVMRestore/ovm/application"
	"github.com/r0x16/PowerstoreOVMRestore/ovm/domain/model"
	"github.com/r0x16/PowerstoreOVMRestore/ovm/infraestructure/repository"

	ms "github.com/r0x16/PowerstoreOVMRestore/sites/domain/model"
	vs "github.com/r0x16/PowerstoreOVMRestore/sites/infraestructure/list/view" // vs is for View Site
)

type WidgetInstance struct {
	Ovm            *model.Ovm
	site           *ms.Site
	widgetProvider *application.OvmWidgetProvider
}

// Process the widget
func OvmWidgetAction(site *ms.Site) *WidgetInstance {
	widgetInstance := &WidgetInstance{
		widgetProvider: application.NewWidgetOvm(repository.NewOvmRepository()),
		site:           site,
	}
	widgetInstance.initOvm()
	return widgetInstance
}

// Set the Ovm if exists, nil otherwise
func (w *WidgetInstance) initOvm() {
	var err error
	ovm, size, err := w.widgetProvider.GetOvmFromSite(w.site)

	if err != nil {
		panic(err)
	}

	if size == 1 {
		w.Ovm = ovm
	}
}

// Draws the widget in the site list based on Ovm value
func (w *WidgetInstance) DrawWidget(drawer *vs.WidgetDrawer) {
	if w.Ovm != nil {
		// Draws current Ovm
	} else {
		// Draws Empty
	}
}
