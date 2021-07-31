package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Port      string `json:"port"`
	Templates string `json:"templates"`
}

func LoadConfig(path string) Config {
	handle, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	var config Config
	err = json.Unmarshal(handle, &config)
	if err != nil {
		log.Fatalln(err)
	}
	return config
}
