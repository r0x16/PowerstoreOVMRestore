package view

import (
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/gui/wizard"
)

func NewIntroductionStepLayout() *fyne.Container {
	title := getTitleLayout()
	content := getContentLayout()
	vb := container.NewVBox(title, widget.NewSeparator(), content)
	return container.NewPadded(vb)
}

func getTitleLayout() *canvas.Text {
	titleLabel := canvas.NewText(config.Lang.Module.CreateSitesWizard.IntroductionTitle, color.Black)
	titleLabel.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	titleLabel.TextSize = wizard.ContentTitleTextSize
	return titleLabel
}

func getContentLayout() *fyne.Container {

	content := "\n" + config.Lang.Module.CreateSitesWizard.IntroductionContent

	vb := container.NewVBox()
	for _, line := range strings.Split(content, "\n") {
		contentLabel := canvas.NewText(line, color.Black)
		contentLabel.TextSize = wizard.ContentTextSize
		vb.Add(contentLabel)
	}

	return vb
}
