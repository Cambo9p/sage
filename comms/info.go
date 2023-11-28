package comms

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// TODO dependancy injection this for
type config struct {
	port int    `yaml:"port"`
	host string `yaml:"host"`
}

var configFile = "config.yml"

func getConfig() (config, error) {
	var config config

	file, err := os.ReadFile(configFile)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to open file %s", configFile))
		return config, err
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to unmarshal contents from %s", configFile))
		return config, err
	}

	return config, nil
}
