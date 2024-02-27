package pkg

import (
	"os"

	"gopkg.in/yaml.v3"
)

func LoadFromYamlFile(path string, v interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, v)
}

// var mainConfig *model.Config

// func SetConfig(config *model.Config) {
// 	mainConfig = config
// }

// func GetConfig() *model.Config {
// 	return mainConfig
// }
