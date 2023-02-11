package view

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// FormData: This struct defines the data requested in the Access Data Step
type AccessDataStepFormData struct {
	hostname      *widget.Entry
	hostnameError *canvas.Text
	port          *widget.Entry
	portError     *canvas.Text
	username      *widget.Entry
	usernameError *canvas.Text
	password      *widget.Entry
}

// Returns the hostname input actual value
func (f *AccessDataStepFormData) GetHostname() string {
	return f.hostname.Text
}

// Set error message for the hostname input
func (f *AccessDataStepFormData) SetHostnameError(msg string) {
	f.hostnameError.Text = msg
	f.hostnameError.Refresh()
}

// Returns the port input actual value
func (f *AccessDataStepFormData) GetPort() string {
	return f.port.Text
}

// Set error message for the port input
func (f *AccessDataStepFormData) SetPortError(msg string) {
	f.portError.Text = msg
	f.portError.Refresh()
}

// Returns the username input actual value
func (f *AccessDataStepFormData) GetUsername() string {
	return f.username.Text
}

// Set error message for the username input
func (f *AccessDataStepFormData) SetUsernameError(msg string) {
	f.usernameError.Text = msg
	f.usernameError.Refresh()
}

// Returns the password input actual value
func (f *AccessDataStepFormData) GetPassword() string {
	return f.password.Text
}

// Clear all error messages shown in the form
func (f *AccessDataStepFormData) ClearErrors() {
	f.hostnameError.Text = ""
	f.hostnameError.Refresh()
	f.portError.Text = ""
	f.portError.Refresh()
	f.usernameError.Text = ""
	f.usernameError.Refresh()
}
