package view

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui/wizard"
)

// Creates the form layout requesting the NameStep Data
func NewNameStepLayout() (*fyne.Container, *NameStepFormData) {
	nameInput, nameLabel := getNameInput()
	descriptionInput, descriptionLabel := getDescriptionInput()
	formData := &NameStepFormData{
		name:        nameInput,
		description: descriptionInput,
	}
	return container.NewVBox(
		nameLabel,
		formData.name,
		descriptionLabel,
		formData.description,
	), formData
}

// --- START FormData: This struct defines the data requested in the NameStep
type NameStepFormData struct {
	name        *widget.Entry
	description *widget.Entry
}

// Get the name string from the input
func (f *NameStepFormData) GetName() string {
	return f.name.Text
}

// Get the description string from the input
func (f *NameStepFormData) GetDescription() string {
	return f.description.Text
}

// --- END FormData

// Returns the name label and input
func getNameInput() (*widget.Entry, *canvas.Text) {
	label := canvas.NewText(config.Lang.Module.CreateSitesWizard.NameInput, color.Black)
	label.TextSize = wizard.ContentTextSize

	return widget.NewEntry(), label
}

// Returns the description label and input
func getDescriptionInput() (*widget.Entry, *canvas.Text) {
	label := canvas.NewText(config.Lang.Module.CreateSitesWizard.DescriptionInput, color.Black)
	label.TextSize = wizard.ContentTextSize

	return widget.NewMultiLineEntry(), label
}
