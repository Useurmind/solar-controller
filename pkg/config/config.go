package config

import (
	"os"

	"github.com/Useurmind/solar-controller/pkg/log"
	"gopkg.in/yaml.v3"
)

func GetConfig(path string) (*Config, error) {
	log.LogInfo("read config from %s", path)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil, err
	}

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := Config{}
	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		return nil, err
	}

	log.LogTrace("config is:\n%s\n", string(b))

	return &cfg, nil
}
