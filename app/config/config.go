package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type (
	ConfigMap struct {
		MongoDbAuth MongoDbAuthConfig
		HTTPData    HTTPConfig
		UserAuth    AuthConfig
	}

	MongoDbAuthConfig struct {
		URI      string `yaml:"uri"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"dbName"`
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
	defer ymlfile.Close()

	//Decoding opened config file
	ymlDecoding := yaml.NewDecoder(ymlfile)
	if err := ymlDecoding.Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
}
