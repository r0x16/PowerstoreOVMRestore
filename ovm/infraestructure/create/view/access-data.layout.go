package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type AccessDataStepLayout struct {
	content *fyne.Container
}

// NewAccessDataStepLayout creates a new layout for the access data step
func NewAccessDataStepLayout() *AccessDataStepLayout {
	return &AccessDataStepLayout{
		content: container.NewVBox(),
	}
}

// Returns the content of the layout
func (l *AccessDataStepLayout) GetContainer() *fyne.Container {
	return l.content
}
