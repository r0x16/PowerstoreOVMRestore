package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type MainLayout struct {
	boxLayout *fyne.Container
	Header    *fyne.Container
	Body      *fyne.Container
	Footer    *fyne.Container
}

func NewMainLayout() *MainLayout {
	ly := new(MainLayout)
	ly.boxLayout = ly.createBoxLayout()
	ly.Header = ly.createHeader()
	ly.Footer = ly.createFooter()

	ly.boxLayout.Add(ly.Header)
	ly.boxLayout.Add(layout.NewSpacer())
	ly.boxLayout.Add(ly.Footer)
	ly.boxLayout.Refresh()
	return ly
}

func (ly *MainLayout) GetContainer() *fyne.Container {
	return ly.boxLayout
}

func (ly *MainLayout) createBoxLayout() *fyne.Container {
	return container.NewVBox()
}

func (ly *MainLayout) createHeader() *fyne.Container {

	titleText := "PowerStore OVM Restore"
	titleColor := color.NRGBA{R: 70, G: 70, B: 70, A: 255}
	title := canvas.NewText(titleText, titleColor)
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.TextSize = 30
	title.Alignment = fyne.TextAlignCenter

	return container.NewGridWithColumns(1, title)

}

func (ly *MainLayout) createFooter() *fyne.Container {
	oraLogo := canvas.NewImageFromFile("assets/ora-logo.png")
	oraLogo.FillMode = canvas.ImageFillOriginal

	dellLogo := canvas.NewImageFromFile("assets/dell-logo.png")
	dellLogo.FillMode = canvas.ImageFillOriginal

	return container.NewHBox(layout.NewSpacer(), oraLogo, dellLogo)
}
