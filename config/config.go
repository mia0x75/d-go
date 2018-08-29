package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	Listen   string `yaml:"listen"`
	Port     int    `yaml:"port"`
	Salt     string `yaml:"salt"`
	Key      string `yaml:"key"`
	Secret   string `yaml:"secret"`
	LogLevel string `yaml:"log_level"`
	Debug    bool   `yaml:"debug"`

	Database struct {
		Host     string `yaml:"host"`
		Port     uint16 `yaml:"port"`
		Db       string `yaml:"db"`
		User     string `yaml:"user"`
		Password string `yaml:"pass"`
		Prefix   string `yaml:"prefix"`
	} `yaml:"database"`
}

func ParseConfigData(data []byte) (*AppConfig, error) {
	var cfg AppConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func ParseConfigFile(fileName string) (*AppConfig, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return ParseConfigData(data)
}
