package config

import (
	configs "go-clean-arch-temp/pkg/config"
)

type (
	Config struct {
		configs.App      `yaml:"app"`
		configs.MariaDB  `yaml:"mariadb"`
		configs.RabbitMQ `yaml:"rabbitmq"`
	}
)
