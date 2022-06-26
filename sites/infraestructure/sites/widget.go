package sites

import "github.com/r0x16/PowerstoreOVMRestore/sites/domain/model"

type SiteWidget interface {
	SetSite(site model.Site)
	SetIndex(i int8)
	Draw()
}

type BaseSiteWidget struct {
	Site  model.Site
	Index int8
}

func (w *BaseSiteWidget) SetSite(site model.Site) {
	w.Site = site
}

func (w *BaseSiteWidget) SetIndex(i int8) {
	w.Index = i
}
