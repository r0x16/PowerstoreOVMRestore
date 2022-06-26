package view

import (
	"fyne.io/fyne/v2"
)

// WidgetDrawer is used to access the site view from the widget
type WidgetDrawer struct {
	listSiteView *ListSiteView
	index        int
}

// Creates a new widget drawer
func NewWidgetDrawer(lv *ListSiteView) *WidgetDrawer {
	return &WidgetDrawer{
		listSiteView: lv,
	}
}

// Draws the widget in the site container
func (w *WidgetDrawer) Draw(item fyne.CanvasObject) {
	w.listSiteView.AddWidget(item, int8(w.index))
}

// Sets the index of the widget for draw
func (w *WidgetDrawer) SetIndex(i int) {
	w.index = i
}
