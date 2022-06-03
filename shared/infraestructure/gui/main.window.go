package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
)

type MainWindow struct {
	*Window
	app fyne.App
}

var MainWindowContainer *MainWindow

func NewMainWindow(title string) {
	MainWindowContainer = &MainWindow{}
	MainWindowContainer.app = app.New()
	MainWindowContainer.Window = NewWindow(title)
	MainWindowContainer.window.SetMaster()
	MainWindowContainer.window.Resize(fyne.NewSize(config.Global.WINDOW_WIDTH, config.Global.WINDOW_HEIGHT))
}

func (w *MainWindow) SetMainContent() {
	NewMainLayout()
	w.SetContent(MainContainer)
}

func (w *MainWindow) PlayMainApplication() {
	w.Play()
}
