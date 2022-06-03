package gui

import (
	"fyne.io/fyne/v2"
)

type Window struct {
	Title  string
	window fyne.Window
}

type Layout interface {
	GetContainer() *fyne.Container
}

func NewWindow(title string) *Window {
	w := &Window{}
	w.Title = title
	w.window = fyne.CurrentApp().NewWindow(w.Title)
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
