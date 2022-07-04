package infraestructure

import (
	"github.com/r0x16/PowerstoreOVMRestore/ovm/infraestructure/widget"

	ms "github.com/r0x16/PowerstoreOVMRestore/sites/domain/model"              // ms is for model site
	vs "github.com/r0x16/PowerstoreOVMRestore/sites/infraestructure/list/view" // vs is for view site
)

type OvmManagerModule struct {
	Site   *ms.Site
	action *widget.WidgetInstance
}

func NewOvmManagerModule(site *ms.Site) *OvmManagerModule {
	return &OvmManagerModule{
		Site: site,
	}
}

func (m *OvmManagerModule) RunWidget(drawer *vs.WidgetDrawer) {
	m.action = widget.OvmWidgetAction(m.Site)
	m.action.DrawWidget(drawer)
}
