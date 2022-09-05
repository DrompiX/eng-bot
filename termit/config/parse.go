package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

func Read(path string) []byte {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Could not open configuration file: %v", err)
	}
	return f
}

func Parse(data []byte) TermitConfig {
	conf := TermitConfig{}
	err := yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Fatalf("Could not parse configuration file: %s", err)
	}
	return conf
}
