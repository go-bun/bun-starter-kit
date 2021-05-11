package app

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	AppDir  string
	Service string
	Env     string

	Debug     bool   `yaml:"debug"`
	SecretKey string `yaml:"secret_key"`

	PGX struct {
		DSN string `yaml:"dsn"`
	} `yaml:"pgx"`
}

func ReadConfig(service, env string) (*AppConfig, error) {
	appDir, err := findAppDir(env)
	if err != nil {
		return nil, fmt.Errorf("findAppDir failed: %w", err)
	}

	f, err := os.Open(configPath(appDir, env))
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	cfg := new(AppConfig)
	if err := yaml.Unmarshal(b, cfg); err != nil {
		return nil, err
	}

	cfg.AppDir = appDir
	cfg.Service = service
	cfg.Env = env

	return cfg, nil
}

func findAppDir(env string) (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	saved := dir

	for i := 0; i < 10; i++ {
		configPath := configPath(dir, env)
		_, err := os.Stat(configPath)
		if err == nil {
			return dir, nil
		}

		if dir == "." {
			break
		}
		dir = filepath.Dir(dir)
	}

	return saved, nil
}

func configPath(dir, env string) string {
	return filepath.Join(dir, "app", "config", env+".yaml")
}
