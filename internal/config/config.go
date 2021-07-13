package config

import (
	"io/ioutil"

	"github.com/lorgioedtech/go-rest-dev/pkg/log"
	"github.com/qiangxue/go-env"
	"gopkg.in/yaml.v2"
)

// Config represents an application configuration.
type Config struct {
	// the data source name (DSN) for connecting to the database. required.
	DatabaseURI string `yaml:"database_uri" env:"DATABASE_URI,secret"`
}

// Load returns an application configuration which is populated from the given configuration file and environment variables.
func Load(file string, logger log.Logger) (*Config, error) {
	// default config
	c := Config{}

	// load from YAML config file
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(bytes, &c); err != nil {
		return nil, err
	}

	// load from environment variables prefixed with "APP_"
	if err = env.New("APP_", logger.Infof).Load(&c); err != nil {
		return nil, err
	}

	return &c, err
}
