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
	nameInput, nameLabel, nameErrorLabel := getNameInput()
	descriptionInput, descriptionLabel := getDescriptionInput()
	formData := &NameStepFormData{
		name:        nameInput,
		nameError:   nameErrorLabel,
		description: descriptionInput,
	}
	return container.NewVBox(
		container.NewHBox(nameLabel, formData.nameError),
		formData.name,
		descriptionLabel,
		formData.description,
	), formData
}

// --- START FormData: This struct defines the data requested in the NameStep
type NameStepFormData struct {
	name        *widget.Entry
	nameError   *canvas.Text
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

// Set the error for the name input
func (f *NameStepFormData) SetNameError(msg string) {
	f.nameError.Text = msg
	f.nameError.Refresh()
}

// Clear all error messages
func (f *NameStepFormData) ClearErrors() {
	f.nameError.Text = ""
	f.nameError.Refresh()
}

// --- END FormData

// Returns the name label and input
func getNameInput() (*widget.Entry, *canvas.Text, *canvas.Text) {
	label := canvas.NewText(config.Lang.Module.CreateSitesWizard.NameInput, color.Black)
	label.TextSize = wizard.ContentTextSize

	errorLabel := canvas.NewText("", color.RGBA{R: 255, A: 255})
	errorLabel.TextSize = wizard.ContentTextSize

	return widget.NewEntry(), label, errorLabel
}

// Returns the description label and input
func getDescriptionInput() (*widget.Entry, *canvas.Text) {
	label := canvas.NewText(config.Lang.Module.CreateSitesWizard.DescriptionInput, color.Black)
	label.TextSize = wizard.ContentTextSize

	return widget.NewMultiLineEntry(), label
}
