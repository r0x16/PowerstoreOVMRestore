package infraestructure

import (
	ms "github.com/r0x16/PowerstoreOVMRestore/sites/domain/model" // ms is for model site
)

type OvmManagerModule struct {
	Site *ms.Site
}

func NewOvmManagerModule(site *ms.Site) *OvmManagerModule {
	return &OvmManagerModule{
		Site: site,
	}
}

func (m *OvmManagerModule) RunWidget() {

}
