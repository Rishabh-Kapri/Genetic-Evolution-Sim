package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// default config file is conf.json
type Config struct {
	SizeX           int
	SizeY           int
	MinGenomeLength int
	MaxGenomeLength int
	Sensor          int
	Action          int
	Neuron          int
	Population      int
}

func LoadConfig(configFilePath string) Config {
	file, _ := os.Open(configFilePath)
	defer file.Close()
	decorder := json.NewDecoder(file)
	conf := Config{}
	err := decorder.Decode(&conf)
	if err != nil {
		fmt.Println(err)
	}
	return conf
}
