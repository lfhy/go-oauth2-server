package config

import (
	"encoding/json"
	"os"
)

type fileBackend struct {
}

func (b *fileBackend) InitConfigBackend() {
}

// 从文件加载配置信息
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
