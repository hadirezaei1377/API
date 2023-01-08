package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// decoding config to GO code

// we want to put all config to a struct
type Config struct {
	Server struct {
		// `` is using for convert yaml to GO code
		Port string `yaml:"port"`
	} `yaml:"server"`
}

var AppConfig Config

// decode yml
// set config for running server at first of start by this function
func GetConfig() error {
	file, err := os.Open("config.yml")
	if err != nil {
		return err
	}
	defer file.Close()
	// decode yml and put it to our struct
	// it need file (config file)
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		return err
	}
	return nil
}
