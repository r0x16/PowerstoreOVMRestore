package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
)

type MainLayout struct {
	boxLayout *fyne.Container
	Header    *fyne.Container
	Body      *container.Split
	Footer    *fyne.Container
}

var MainContainer *MainLayout

func NewMainLayout() {
	ly := &MainLayout{}
	ly.boxLayout = ly.createBoxLayout()
	ly.Header = ly.createHeader()
	ly.Body = ly.createBody()
	ly.Body.SetOffset(0.3)
	ly.Footer = ly.createFooter()

	ly.boxLayout.Add(ly.Header)
	ly.boxLayout.Add(widget.NewSeparator())
	ly.boxLayout.Add(ly.Body)
	ly.boxLayout.Add(layout.NewSpacer())
	ly.boxLayout.Add(ly.Footer)

	MainContainer = ly
}

func (ly *MainLayout) GetContainer() *fyne.Container {
	return ly.boxLayout
}

func (ly *MainLayout) createBoxLayout() *fyne.Container {
	return container.NewVBox()
}

func (ly *MainLayout) createHeader() *fyne.Container {

	titleText := config.Global.APP_NAME
	titleColor := color.NRGBA{R: 70, G: 70, B: 70, A: 255}
	title := canvas.NewText(titleText, titleColor)
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.TextSize = 30
	title.Alignment = fyne.TextAlignCenter

	return container.NewCenter(title)

}

func (ly *MainLayout) createFooter() *fyne.Container {
	oraLogo := canvas.NewImageFromFile("assets/ora-logo.png")
	oraLogo.FillMode = canvas.ImageFillOriginal

	dellLogo := canvas.NewImageFromFile("assets/dell-logo.png")
	dellLogo.FillMode = canvas.ImageFillOriginal

	return container.NewHBox(layout.NewSpacer(), oraLogo, dellLogo)
}

func (ly *MainLayout) createBody() *container.Split {
	grey := &color.NRGBA{R: 140, G: 140, B: 140, A: 255}

	empty := canvas.NewText("Empty", grey)
	empty.TextSize = 10

	empty2 := canvas.NewText("Empty", grey)
	empty2.TextSize = 10

	return container.NewHSplit(empty, empty2)
}
