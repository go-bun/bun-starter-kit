package bunapp

import (
	"embed"
	"io/fs"
	"os"
	"path"
	"sync"

	"gopkg.in/yaml.v3"
)

var (
	//go:embed embed
	embedFS      embed.FS
	unwrapFSOnce sync.Once
	unwrappedFS  fs.FS
)

func FS() fs.FS {
	unwrapFSOnce.Do(func() {
		fsys, err := fs.Sub(embedFS, "embed")
		if err != nil {
			panic(err)
		}
		unwrappedFS = fsys
	})
	return unwrappedFS
}

type AppConfig struct {
	Service string
	Env     string

	Debug     bool   `yaml:"debug"`
	SecretKey string `yaml:"secret_key"`

	DB struct {
		DSN    string `yaml:"dsn"`
		Driver string `yaml:"driver"`
	} `yaml:"db"`
}

func ReadConfig(fsys fs.FS, service, env string) (*AppConfig, error) {
	b, err := fs.ReadFile(fsys, path.Join("config", env+".yaml"))
	if err != nil {
		return nil, err
	}

	confContent := []byte(os.ExpandEnv(string(b)))
	cfg := new(AppConfig)
	if err := yaml.Unmarshal(confContent, cfg); err != nil {
		panic(err)
	}

	cfg.Service = service
	cfg.Env = env

	return cfg, nil
}
