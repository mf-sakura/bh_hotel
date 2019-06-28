package config

import (
	"github.com/kelseyhightower/envconfig"
)

type DBConfig struct {
	Port     int    `envconfig:"MYSQL_PORT" default:"3306"`
	Password string `envconfig:"MYSQL_PASSWORD"`
	User     string `envconfig:"MYSQL_USER" default:"root"`
	Host     string `envconfig:"MYSQL_HOST" default:"localhost"`
}

func LoadConifg() (*DBConfig, error) {
	var dbConf DBConfig
	if err := envconfig.Process("", &dbConf); err != nil {
		return nil, err
	}

	return &dbConf, nil
}
