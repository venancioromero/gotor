package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

const (
	NOT_PROVIDED_PATH           = "Please, provide a valid path to config file"
	CONFIG_FILE_DOES_NOT_EXISTS = "Config file doesn't exists"
	CONFIG_FILE_PARSING_ERROR   = "Config file parsing error"
)

type Logging struct {
	LogLevel string `toml:"log_level"`
}

type TorProxy struct {
	Url string `toml:"url"`
}

type Config struct {
	Logging  Logging  `toml:"logging"`
	TorProxy TorProxy `toml:"torProxy"`
}

func LoadConfig(configFilePath string) Config {
	if configFilePath == "" {
		panic(NOT_PROVIDED_PATH)
	}

	if !fileExists(configFilePath) {
		panic(CONFIG_FILE_DOES_NOT_EXISTS)
	}

	var conf Config
	if _, err := toml.DecodeFile(configFilePath, &conf); err != nil {
		panic(CONFIG_FILE_PARSING_ERROR)
	}

	return conf
}

func fileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
