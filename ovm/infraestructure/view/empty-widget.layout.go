package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
)

// Get the button to setup the OVM
func GetEmptyWidgetContainer(linkOvm func()) *fyne.Container {
	return container.NewHBox(
		getEmptyMessage(),
		layout.NewSpacer(),
		widget.NewButtonWithIcon("", theme.SettingsIcon(), linkOvm),
	)
}

// Get the message indicating that the OVM is not set
func getEmptyMessage() *canvas.Text {
	message := canvas.NewText(config.Lang.Module.OvmWidget.OvmNotFound, theme.DisabledColor())
	message.TextStyle = fyne.TextStyle{Bold: true}
	message.TextSize = 10
	message.Alignment = fyne.TextAlignCenter
	return message
}
