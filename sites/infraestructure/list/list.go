package list

import (
	"time"

	"github.com/r0x16/PowerstoreOVMRestore/sites/application"
	"github.com/r0x16/PowerstoreOVMRestore/sites/domain/model"
	"github.com/r0x16/PowerstoreOVMRestore/sites/infraestructure/list/view"
	"github.com/r0x16/PowerstoreOVMRestore/sites/infraestructure/repository"
)

type ListSitesInstance struct {
	timer         int64
	Sites         []model.Site
	Size          int64
	sitesProvider *application.SiteListProvider
	view          *view.ListSiteView
}

var listInstance *ListSitesInstance

func ListSitesAction() {
	if !ready() {
		return
	}

	listInstance = &ListSitesInstance{
		timer: now(),
	}

	listInstance.sitesProvider = application.NewListSite(repository.NewSiteRepository())
	listInstance.refreshSites()

	listInstance.view = view.NewSiteListLayout()
	listInstance.view.SetSites(listInstance.Sites)

}

func ready() bool {
	if listInstance == nil {
		return true
	}

	return (now() - listInstance.timer) >= 5000

}

func now() int64 {
	return time.Now().UnixMicro()
}

func (l *ListSitesInstance) refreshSites() {
	var err error
	listInstance.Sites, listInstance.Size, err = listInstance.sitesProvider.GetAllSites()

	if err != nil {
		panic(err)
	}

}
