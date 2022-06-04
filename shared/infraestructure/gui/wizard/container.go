package wizard

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/r0x16/PowerstoreOVMRestore/shared/infraestructure/config"
)

// Represents all de parts for the border layout of the wizard
type WizardContainers struct {
	titleContainer         *fyne.Container
	stepsContainer         *fyne.Container
	taskContainer          *fyne.CanvasObject
	stepsButtonsCointainer *fyne.Container
}

// Instantiates a new container group
func (w *Wizard) buildContainers() *WizardContainers {
	return &WizardContainers{
		titleContainer:         w.buildTitleContainer(),
		stepsContainer:         w.buildStepsContainer(),
		taskContainer:          w.buildTaskContainer(),
		stepsButtonsCointainer: w.buildStepButtonsContainer(),
	}

}

// Builds the header container, with tittle and caption
func (w *Wizard) buildTitleContainer() *fyne.Container {
	title := canvas.NewText(w.Title, theme.TextColor())
	title.TextSize = 25
	title.Alignment = fyne.TextAlignCenter

	w.caption.TextSize = 10
	w.caption.Alignment = fyne.TextAlignCenter

	return container.NewVBox(
		title,
		w.caption,
		widget.NewSeparator(),
	)
}

// Builds steps lists for showing current task
func (w *Wizard) buildStepsContainer() *fyne.Container {
	vb := container.NewVBox()
	for _, step := range w.Steps {
		vb.Add(widget.NewLabel(step.GetTitle()))
	}
	return vb
}

// Defaults the task container with first task element
func (w *Wizard) buildTaskContainer() *fyne.CanvasObject {
	return w.Steps[0].Load()
}

// Builds all the button to navigate through the wizard
func (w *Wizard) buildStepButtonsContainer() *fyne.Container {
	w.BackButton, w.NextButton, w.FinishButton = w.buildStepButtons()

	return container.NewHBox(
		layout.NewSpacer(),
		w.BackButton,
		w.NextButton,
		w.FinishButton,
	)
}

func (w *Wizard) buildStepButtons() (*widget.Button, *widget.Button, *widget.Button) {
	backButton := widget.NewButton(
		config.Lang.Wizard.BackButton, func() {
			w.Back()
		})
	backButton.Disable()

	nextButton := widget.NewButton(
		config.Lang.Wizard.NextButton, func() {
			w.Next()
		})

	finishButton := widget.NewButton(
		config.Lang.Wizard.FinishButton, func() {
			w.Finish()
		})
	finishButton.Disable()

	return backButton, nextButton, finishButton
}
