package wizard

import (
	"fmt"

	"fyne.io/fyne/v2"
)

// Defines the minimum behavior to a step
type WizardStep interface {
	GetTitle() string
	GetCaption() string
	Load() *fyne.CanvasObject
	Validate() bool
}

// This method controls the back button behaviour
func (w *Wizard) Back() {
	fmt.Println("Previous")
}

// This method controls the next button behaviour
func (w *Wizard) Next() {
	fmt.Println("Next")
}

// This method controls the finish button behaviour
func (w *Wizard) Finish() {
	fmt.Println("Finish")
}
