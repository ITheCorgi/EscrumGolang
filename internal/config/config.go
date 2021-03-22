package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type (
	ConfigMap struct {
		Mongo    MongoConfig
		HTTPData HTTPConfig
		UserAuth AuthConfig
	}

	MongoConfig struct {
		URI      string `yaml:"uri"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	}

	HTTPConfig struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}

	AuthConfig struct {
		//FIXME: issue #3 at the github project workflow
	}
)

//getConfig function opens yaml config file, then gets settings into Config struct
func GetConfig(path string) (*ConfigMap, error) {
	config := &ConfigMap{}

	//Opening a config file
	ymlfile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	buffer, _ := ioutil.ReadAll(ymlfile)
	defer ymlfile.Close()

	err = yaml.Unmarshal(buffer, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
