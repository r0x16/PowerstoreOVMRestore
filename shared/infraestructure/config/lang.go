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
}

type ModuleLang struct {
	CreateSitesWizard string `json:"CreateSitesWizard"`
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
