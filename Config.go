package config

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
 * ConfigurationStruct holds some generic configuration information but is
 * primarily there to hang the related methods off of.
 */
type ConfigurationStruct struct {
	//Basic Site Configuration
	Host              string
	StaticPath        string
	TemplateRoot      string
	//Basic Security Configuration
	AdminToken        string
	MinimumNameLength int
	MinPasswordLength int
	EncKey            string
	AuthKey           string
	CookieName        string

	//Database Configuration
	DatabaseName         string
	DatabaseUserName     string
	DatabasePassword     string
	DatabaseServerURL    string
}


//LoadConfigs loads a configuration file and stuffs it into the interface named config.
func LoadConfigs(config interface{}, envVariable string) {
	configPath := getConfigPath(envVariable)

	file, err := os.Open(configPath)
	if err != nil {
		panic("Configuration file did not open!")
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)

	if err != nil {
		fmt.Println("Configuration File Malformed")
		panic(err)
	}

}


//getConfigPath() Gets the config path from the environmental variable.
func getConfigPath(envVariable string) string {
	//First we figure out where the configuration files should be.
	configpath, isEnv := os.LookupEnv(envVariable)
	if !isEnv {
		fmt.Println("Could not find the env variable for the configs.")
		panic("No env variable for configs.")
	}
	return configpath
}

