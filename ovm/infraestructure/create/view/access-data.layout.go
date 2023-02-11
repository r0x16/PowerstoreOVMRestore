package view

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
)

// Represents the form object
// contains view containers and View data access methods
type AccessDataStepLayout struct {
	content *fyne.Container
	Data    *AccessDataStepFormData
}

// NewAccessDataStepLayout creates a new layout for the access data step
func NewAccessDataStepLayout() *AccessDataStepLayout {
	hostnameInput, hostnameLabel, hostnameErrorLabel := getHostnameInput()
	usernameInput, usernameLabel, usernameErrorLabel := getUsernameInput()
	passwordInput, passwordLabel, passwordErrorLabel := getPasswordInput()
	portInput, portLabel, portErrorLabel := getPortNumberInput()
	formData := &AccessDataStepFormData{
		hostname:      hostnameInput,
		hostnameError: hostnameErrorLabel,
		username:      usernameInput,
		usernameError: usernameErrorLabel,
		password:      passwordInput,
		port:          portInput,
		portError:     portErrorLabel,
	}
	return &AccessDataStepLayout{
		content: container.NewVBox(
			container.NewHBox(hostnameLabel, hostnameErrorLabel),
			formData.hostname,
			container.NewHBox(portLabel, portErrorLabel),
			formData.port,
			container.NewHBox(usernameLabel, usernameErrorLabel),
			formData.username,
			container.NewHBox(passwordLabel, passwordErrorLabel),
			formData.password,
		),
		Data: formData,
	}
}

// Returns the content of the layout
func (l *AccessDataStepLayout) GetContainer() *fyne.Container {
	return l.content
}

// Returns the hostname input
func getHostnameInput() (*widget.Entry, *canvas.Text, *canvas.Text) {
	hostnameInput := widget.NewEntry()
	hostnameLabel := canvas.NewText(
		config.Lang.Module.CreateOvmWizard.HostnameInput,
		color.Black,
	)
	hostnameErrorLabel := canvas.NewText("", color.RGBA{R: 255, A: 255})
	return hostnameInput, hostnameLabel, hostnameErrorLabel
}

// Returns the username input
func getUsernameInput() (*widget.Entry, *canvas.Text, *canvas.Text) {
	usernameInput := widget.NewEntry()
	usernameLabel := canvas.NewText(
		config.Lang.Module.CreateOvmWizard.UsernameInput,
		color.Black,
	)
	usernameErrorLabel := canvas.NewText("", color.RGBA{R: 255, A: 255})
	return usernameInput, usernameLabel, usernameErrorLabel
}

// Returns the password input
func getPasswordInput() (*widget.Entry, *canvas.Text, *canvas.Text) {
	passwordInput := widget.NewPasswordEntry()
	passwordLabel := canvas.NewText(
		config.Lang.Module.CreateOvmWizard.PasswordInput,
		color.Black,
	)
	passwordErrorLabel := canvas.NewText("", color.RGBA{R: 255, A: 255})
	return passwordInput, passwordLabel, passwordErrorLabel
}

// Returns the port number input
func getPortNumberInput() (*widget.Entry, *canvas.Text, *canvas.Text) {
	portNumberInput := widget.NewEntry()
	portNumberLabel := canvas.NewText(
		config.Lang.Module.CreateOvmWizard.PortInput,
		color.Black,
	)
	portNumberErrorLabel := canvas.NewText("", color.RGBA{R: 255, A: 255})
	return portNumberInput, portNumberLabel, portNumberErrorLabel
}
