package go_config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var profilesFlag *string = nil
var profiles []string

func init() {
	profilesFlag = flag.String("profiles", "", "The configuration profiles to use")
}

// Load loads the configuration files.
// First configuration file will be config.json followed config-{profile}.json,
// for each profile in the order they are specified
func Load(target interface{}) error {
	if profilesFlag == nil {
		panic("go_config.init() needs to called")
	}

	if !flag.Parsed() {
		panic("flag.Parse() needs to be called first")
	}

	parseProfiles(*profilesFlag)
	log.Printf("Active profiles: %v", profiles)
	if target != nil {
		return loadConfigurationFiles(target)
	}

	return nil
}

// GetCurrentProfiles returns the currently active profiles
func GetCurrentProfiles() []string {
	return profiles
}

// IsProfileActive indicates whether the given profile is active or not
func IsProfileActive(profile string) bool {
	for _, p := range profiles {
		if p == profile {
			return true
		}
	}
	return false
}

func parseProfiles(profilesFlag string) {
	if profilesFlag == "" {
		return
	}

	profiles = strings.Split(profilesFlag, ",")
	for i, profile := range profiles {
		profiles[i] = strings.TrimSpace(profile)
	}
}

func loadConfigurationFiles(target interface{}) error {
	configFiles := make([]string, len(profiles)+1)
	configFiles[0] = "config.json"
	for i, profile := range profiles {
		configFiles[i+1] = fmt.Sprintf("config-%s.json", profile)
	}

	for _, configFile := range configFiles {
		contents, err := ioutil.ReadFile(configFile)
		if os.IsNotExist(err) {
			log.Printf("Config file %s does not exist. Skipping", configFile)
			continue
		} else if err != nil {
			return err
		}

		if err := json.Unmarshal(contents, target); err != nil {
			return err
		}
	}

	return nil
}
