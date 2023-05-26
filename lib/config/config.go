package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var Files map[string]map[interface{}]interface{}

func InitConfigFiles() (map[string]map[interface{}]interface{}, error) {
	var mappings = make(map[string]map[interface{}]interface{})
	configs := [...]string{"config.yml"}
	for _, config := range configs {
		configFile, err := os.ReadFile(fmt.Sprintf("%s/%s", "lib/config/yaml", config))
		if err != nil {
			return mappings, err
		}

		data := make(map[interface{}]interface{})
		err = yaml.Unmarshal(configFile, &data)
		if err != nil {
			return mappings, err
		}

		mappings[config] = data
	}
	return mappings, nil
}

func GetValue(key string) string {
	for _, v := range Files {
		if v[key] != nil {
			return v[key].(string)
		}
	}
	return ""
}
