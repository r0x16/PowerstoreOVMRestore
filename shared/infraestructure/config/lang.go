package config

import (
	"path"
	"strings"

	"github.com/tkanos/gonfig"
)

// Language Configuration struct from json
// ----------------

type LangConfiguration struct {
	Module ModuleLang `json:"Module"`
	Wizard WizardLang `json:"Wizard"`
}

type CreateSitesWizardLang struct {
	WizardTitle              string `json:"WizardTitle"`
	IntroductionSectionTitle string `json:"IntroductionSectionTitle"`
	IntroductionTitle        string `json:"IntroductionTitle"`
	IntroductionContent      string `json:"IntroductionContent"`
}
type ModuleLang struct {
	CreateSitesWizard CreateSitesWizardLang `json:"CreateSitesWizard"`
}

type WizardLang struct {
	BackButton   string `json:"backButton"`
	NextButton   string `json:"nextButton"`
	FinishButton string `json:"finishButton"`
	CancelButton string `json:"cancelButton"`
	CloseButton  string `json:"closeButton"`
}

// ----------------

// Public language access
var Lang LangConfiguration = LangConfiguration{}

// Loads the language configuration file and outputs into language object
func InitializeLang() error {
	err := gonfig.GetConf(getLangConfigFilename(), &Lang)

	if err != nil {
		return err
	}

	return nil
}

// Calculates language filename's path
func getLangConfigFilename() string {
	// Current supports english only
	filename := []string{"en", ".json"}

	filePath := path.Join("config/lang", strings.Join(filename, ""))

	return filePath
}
