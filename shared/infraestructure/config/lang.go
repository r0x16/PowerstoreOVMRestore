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
	NameSectionTitle         string `json:"NameSectionTitle"`
	NameCaption              string `json:"NameCaption"`
	NameInput                string `json:"NameInput"`
	DescriptionInput         string `json:"DescriptionInput"`
	SummarySectionTitle      string `json:"SummarySectionTitle"`
	SummaryCaption           string `json:"SummaryCaption"`
	OnFinishSuccess          string `json:"OnFinishSuccess"`
	Errors                   struct {
		NameEmpty      string `json:"NameEmpty"`
		NameFormat     string `json:"NameFormat"`
		NameExists     string `json:"NameExists"`
		NameStoreError string `json:"NameStoreError"`
	} `json:"Errors"`
}

type OvmWidgetLang struct {
	OvmManager  string `json:"OvmManager,omitempty"`
	OvmNotFound string `json:"OvmNotFound,omitempty"`
}

type CreateOvmWizardLang struct {
	WizardTitle              string `json:"WizardTitle"`
	IntroductionSectionTitle string `json:"IntroductionSectionTitle"`
	IntroductionTitle        string `json:"IntroductionTitle"`
	IntroductionContent      string `json:"IntroductionContent"`
	AccessDataSectionTitle   string `json:"AccessDataSectionTitle"`
	AccessDataCaption        string `json:"AccessDataCaption"`
	HostnameInput            string `json:"HostnameInput"`
	UsernameInput            string `json:"UsernameInput"`
	PasswordInput            string `json:"PasswordInput"`
	PortInput                string `json:"PortInput"`
	SummarySectionTitle      string `json:"SummarySectionTitle"`
	SummaryCaption           string `json:"SummaryCaption"`
	OnFinishSuccess          string `json:"OnFinishSuccess"`
	Errors                   struct {
		HostnameEmpty  string `json:"HostnameEmpty"`
		HostnameFormat string `json:"HostnameFormat"`
		HostnameExists string `json:"HostnameExists"`
		UsernameEmpty  string `json:"UsernameEmpty"`
		UsernameFormat string `json:"UsernameFormat"`
		PasswordEmpty  string `json:"PasswordEmpty"`
		PortEmpty      string `json:"PortEmpty"`
		PortFormat     string `json:"PortFormat"`
		OvmStoreError  string `json:"OvmStoreError"`
	} `json:"Errors"`
}

type ModuleLang struct {
	CreateSitesWizard CreateSitesWizardLang `json:"CreateSitesWizard"`
	OvmWidget         OvmWidgetLang         `json:"OvmWidget"`
	CreateOvmWizard   CreateOvmWizardLang   `json:"CreateOvmWizard"`
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
