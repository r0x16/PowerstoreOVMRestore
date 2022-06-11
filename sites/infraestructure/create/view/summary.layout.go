package view

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui/wizard"
	"github.com/r0x16/PowerstoreOVMRestore/sites/application"
)

// Create a table format container with the site current data
func NewSummaryStepLayout(data *application.SiteCreationData) *fyne.Container {
	table := container.NewVBox(
		getRow(
			config.Lang.Module.CreateSitesWizard.NameInput,
			data.Name,
		),
		getRow(
			config.Lang.Module.CreateSitesWizard.DescriptionInput,
			data.Description,
		),
	)
	return table
}

// Gets a label with the content key name
func getKeyLabel(keyname string) *canvas.Text {
	label := canvas.NewText(keyname, color.Black)
	label.TextSize = wizard.ContentTextSize - 1
	label.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	return label
}

// Gets a label with the content value
func getValueLabel(value string) *canvas.Text {
	label := canvas.NewText(value, color.Black)
	label.TextSize = wizard.ContentTextSize + 2
	return label
}

// Get a container representing a row for the summary table
func getRow(key, value string) *fyne.Container {
	keyLabel := getKeyLabel(key)
	valueLabel := getValueLabel(value)
	return container.NewVBox(
		keyLabel,
		valueLabel,
		widget.NewSeparator(),
		getRowSeparator(),
	)
}

// Get a simple label representing a vetical spacer for rows
func getRowSeparator() *canvas.Text {
	separator := canvas.NewText("", color.Transparent)
	separator.TextSize = 3
	return separator
}
