package config

import (
	"os"
	"path"
	"strings"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	APP_NAME string
	DB       string
	WINDOW_WIDTH,
	WINDOW_HEIGHT float32
}

var Global Configuration = Configuration{}

func Initialize() error {
	configFile := getGlobalConfigFilename()

	err := gonfig.GetConf(configFile, &Global)

	if err != nil {
		return err
	}

	return nil
}

func getGlobalConfigFilename() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "dev"
	}
	filename := []string{"config.", env, ".json"}

	filePath := path.Join("config", strings.Join(filename, ""))

	return filePath
}
