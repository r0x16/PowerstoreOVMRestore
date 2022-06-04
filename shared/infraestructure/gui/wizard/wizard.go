package wizard

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Represents a wizard pane content with logic to navigate between steps
// with validation and completition logic
type Wizard struct {
	BorderLayout *fyne.Container
	containers   *WizardContainers
	currentStep  int
	caption      *canvas.Text

	// Step buttons
	BackButton   *widget.Button
	NextButton   *widget.Button
	FinishButton *widget.Button

	Steps []WizardStep
	Title string
}

// Instantiates a new wizard
func NewWizard(title string, steps []WizardStep) *Wizard {
	w := &Wizard{
		Title:       title,
		Steps:       steps,
		currentStep: 0,
		caption:     canvas.NewText(steps[0].GetCaption(), color.Gray{}),
	}
	w.containers = w.buildContainers()
	w.BorderLayout = w.buildBorderLayout()
	return w
}

// Build the main layout where all parts of wizard resides
func (w *Wizard) buildBorderLayout() *fyne.Container {
	return container.NewBorder(
		w.containers.titleContainer,
		w.containers.stepsButtonsCointainer,
		w.containers.stepsContainer,
		nil,
		*w.containers.taskContainer,
	)
}

// Implements gui.Layout, this makes possible to set the content on window
func (w *Wizard) GetContainer() *fyne.Container {
	return w.BorderLayout
}
