package wizard

import "fyne.io/fyne/v2"

type WizardConfig interface {
	GetTitle() string
	OnClose()
	OnFinish() *fyne.Container
}

type BaseWizardConfig struct {
	Title string
}

func (wc *BaseWizardConfig) GetTitle() string {
	return wc.Title
}
