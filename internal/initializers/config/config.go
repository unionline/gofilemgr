package config

import (
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

// Setting ...
var Setting Config

const (
	ConfigPath = "config/config.toml"
)

// Init ...
func Init() {
	if err := LoadConfigs(ConfigPath); err != nil {
		panic(err)
	}
}

// LoadConfigs ...
func LoadConfigs(files ...string) error {
	for _, file := range files {
		if _, err := toml.DecodeFile(file, &Setting); err != nil {
			logrus.Error("decode error: ", err)
			return err
		}
	}

	return nil
}
