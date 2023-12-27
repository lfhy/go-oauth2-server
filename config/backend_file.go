package config

import (
	"encoding/json"
	"os"
)

type fileBackend struct {
}

func (b *fileBackend) InitConfigBackend() {
}

// LoadConfig gets the JSON from ETCD and unmarshals it to the config object
func (b *fileBackend) LoadConfig() (*Config, error) {
	data, err := os.ReadFile(ConfigPath)
	if err != nil {
		return nil, err
	}
	newCnf := new(Config)
	if err := json.Unmarshal(data, newCnf); err != nil {
		return nil, err
	}

	return newCnf, nil
}

func (b *fileBackend) RefreshConfig(newCnf *Config) {
	*Cnf = *newCnf
}
