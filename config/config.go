package config

import (
	"os"
	"path/filepath"

	"github.com/kirsle/configdir"
	"gopkg.in/yaml.v3"
)

type ConfigS struct {
	OnlyWarnings                      bool
	Latitude                          float64
	Longitude                         float64
	IssueWarningAtEquivalentMagnitude float64 // 5.5 is a good default
	IssueWarningAtAnyMagnitude        bool
	IssueWarningSound                 bool
	IssuePopup                        bool
	IssueNotification                 bool
	OpenWebPages                      bool
	TestWarning                       bool
	RetryConnectionEveryXS            float64
}

var Config ConfigS

var DefaultConfig = ConfigS{
	OnlyWarnings:                      false,
	Latitude:                          35.6799833,
	Longitude:                         139.7655883,
	IssueWarningAtEquivalentMagnitude: 5.5,
	IssueWarningAtAnyMagnitude:        false,
	IssueWarningSound:                 false,
	IssuePopup:                        true,
	IssueNotification:                 true,
	OpenWebPages:                      false,
	TestWarning:                       false,
	RetryConnectionEveryXS:            30,
}

func Init() {
	configPath := configdir.LocalConfig("ontake", "hayai")
	err := configdir.MakePath(configPath) // Ensure it exists.
	if err != nil {
		panic(err)
	}

	configFile := filepath.Join(configPath, "config.yml")

	// Does the file not exist?
	if _, err = os.Stat(configFile); os.IsNotExist(err) {
		// Create the new config file.
		fh, err := os.Create(configFile)
		if err != nil {
			panic(err)
		}
		defer fh.Close()

		encoder := yaml.NewEncoder(fh)
		encoder.Encode(&DefaultConfig)
		Config = DefaultConfig
	} else {
		// Load the existing file.
		fh, err := os.Open(configFile)
		if err != nil {
			panic(err)
		}
		defer fh.Close()

		decoder := yaml.NewDecoder(fh)
		decoder.Decode(&Config)
	}
}
