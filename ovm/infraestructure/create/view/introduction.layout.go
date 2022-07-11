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

type IntroductionStepLayout struct {
	container *fyne.Container
}

func NewIntroductionStepLayout(siteName string) *IntroductionStepLayout {
	title := getTitleLayout()
	content := getContentLayout(siteName)
	vb := container.NewVBox(title, widget.NewSeparator(), content)
	body := container.NewPadded(vb)
	return &IntroductionStepLayout{
		container: body,
	}
}

func (ly IntroductionStepLayout) GetContainer() *fyne.Container {
	return ly.container
}

func getTitleLayout() *canvas.Text {
	titleLabel := canvas.NewText(config.Lang.Module.CreateOvmWizard.IntroductionTitle, color.Black)
	titleLabel.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	titleLabel.TextSize = wizard.ContentTitleTextSize
	return titleLabel
}

func getContentLayout(siteName string) *fyne.Container {

	content := "\n" + config.Lang.Module.CreateOvmWizard.IntroductionContent
	content = strings.ReplaceAll(content, "{siteName}", siteName)

	vb := container.NewVBox()
	for _, line := range strings.Split(content, "\n") {
		contentLabel := canvas.NewText(line, color.Black)
		contentLabel.TextSize = wizard.ContentTextSize
		vb.Add(contentLabel)
	}

	return vb
}
