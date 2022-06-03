package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
)

type Window struct {
	Title  string
	window fyne.Window
}

type MainWindow struct {
	Window
	app fyne.App
}

type Layout interface {
	GetContainer() *fyne.Container
}

func NewMainWindow(title string) *MainWindow {
	gui := new(MainWindow)
	gui.app = app.New()
	gui.Window = *gui.NewWindow(title)
	gui.window.SetMaster()
	gui.window.Resize(fyne.NewSize(config.Global.WINDOW_WIDTH, config.Global.WINDOW_HEIGHT))
	return gui
}

func (g *MainWindow) NewWindow(title string) *Window {
	w := new(Window)
	w.Title = title
	w.window = g.app.NewWindow(w.Title)
	w.window.CenterOnScreen()
	return w
}

func (g *Window) Play() {
	g.window.ShowAndRun()
}

func (g *Window) Show() {
	g.window.Show()
}

func (g *MainWindow) Run() {
	g.app.Run()
}

func (g *Window) SetContent(ly Layout) {
	g.window.SetContent(ly.GetContainer())
}
