package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/config"
)

type Config struct {
	MetricsPort string `yaml:"metrics_port"`

	BotToken string `yaml:"bot_token"`

	DBConnConfig *DBConnConfig `yaml:"db_conn_conf"`
}

type DBConnConfig struct {
	User             string `yaml:"user"`
	Password         string `yaml:"password"`
	Host             string `yaml:"host"`
	Port             string `yaml:"port"`
	DBName           string `yaml:"db_name"`
	MigrationsEnable bool   `yaml:"migrations_enable"`
}

func InitConfig() (*Config, error) {
	cfgPath := os.Getenv("CONFIG_PATH")
	if cfgPath == "" {
		return nil, fmt.Errorf("config path is empty (help: set CONFIG_PATH=<path>)")
	}
	env := os.Getenv("ENV")
	if env == "" {
		return nil, fmt.Errorf("env is empty (help: set ENV=<env>)")
	}

	base, err := os.Open(fmt.Sprintf("%s/config.yaml", cfgPath))
	if err != nil {
		return nil, errors.Wrap(err, "didn't find base config")
	}

	override, err := os.Open(fmt.Sprintf("%s/config_%s.yaml", cfgPath, strings.ToLower(env)))
	if err != nil {
		return nil, errors.Wrap(err, "didn't find environment config file")
	}

	mergeCfg, err := config.NewYAML(config.Source(base), config.Source(override))
	if err != nil {
		return nil, errors.Wrap(err, "failed merge config")
	}

	var cfg Config
	if err := mergeCfg.Get("config").Populate(&cfg); err != nil {
		return nil, errors.Wrap(err, "marshal config")
	}

	return &cfg, nil
}
