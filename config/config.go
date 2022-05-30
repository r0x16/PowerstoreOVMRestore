package config

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
}

var Global Configuration = Configuration{}

func Initialize() error {
	configFile, err := getGlobalConfigFilename()

	if err != nil {
		return err
	}

	err = gonfig.GetConf(configFile, &Global)

	if err != nil {
		return err
	}

	return nil
}

func getGlobalConfigFilename() (string, error) {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "dev"
	}
	filename := []string{"config.", env, ".json"}
	_, dirname, _, ok := runtime.Caller(0)

	if !ok {
		return "", errors.New("Cannot get environment configuration file.")
	}

	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))

	return filePath, nil
}
