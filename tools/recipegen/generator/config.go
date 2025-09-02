package generator

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
)

var ConfigPattern = regexp.MustCompile(`.*\.recipegen\.json$`)

type Config struct {
	Template string `json:"template"`
	Source   string `json:"source"`
	Output   string `json:"output"`
}

func IsConfigFile(path string) bool {
	return ConfigPattern.MatchString(filepath.Base(path))
}

func LoadConfig(path string) (conf Config, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&conf)
	return
}

func LoadTemplate(path string) (template string, err error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return
	}

	template = string(f)
	return
}
