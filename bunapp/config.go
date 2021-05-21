package bunapp

import (
	"io/fs"
	"path"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Service string
	Env     string

	Debug     bool   `yaml:"debug"`
	SecretKey string `yaml:"secret_key"`

	DB struct {
		DSN string `yaml:"dsn"`
	} `yaml:"db"`
}

func ReadConfig(fsys fs.FS, service, env string) (*AppConfig, error) {
	b, err := fs.ReadFile(fsys, path.Join("config", env+".yaml"))
	if err != nil {
		return nil, err
	}

	cfg := new(AppConfig)
	if err := yaml.Unmarshal(b, cfg); err != nil {
		return nil, err
	}

	cfg.Service = service
	cfg.Env = env

	return cfg, nil
}
