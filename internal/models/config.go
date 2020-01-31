package models

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	SecretKey string `yaml:"secretKey"`
	Server    struct {
		IP   string `yaml:"ip"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	EmailServer struct {
		SMTP     string `yaml:"smtp"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"emailServer"`
	RootUser struct {
		Email    string `yaml:"email"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

func LoadConfigurationFromFile(filepath string) (*Configuration, error) {
	c := &Configuration{}
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("error while reading configuration from file")
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return nil, fmt.Errorf("error while parsing configuration as yaml file")
	}
	c.SecretKey = base64.StdEncoding.EncodeToString([]byte(c.SecretKey))
	return c, nil
}
