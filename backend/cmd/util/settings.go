package util

import (
	"log"
	"os"

	"github.com/michaelpeterswa/shortlink2/backend/cmd/structs"
	"gopkg.in/yaml.v2"
)

var settings *structs.Settings

func init() {
	fileSettings, err := os.ReadFile("settings.yaml")
	if err != nil {
		log.Println("Error loading settings.yaml file")
	}

	err = yaml.Unmarshal(fileSettings, &settings)
	if err != nil {
		log.Println("Error unmarshalling settings")
	}
}

func GetSettings() *structs.Settings {
	return settings
}
