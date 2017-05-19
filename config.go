package main

import "github.com/kelseyhightower/envconfig"

type Config struct {
	BrokerUsername string `envconfig:"broker_username" required:"true"`
	BrokerPassword string `envconfig:"broker_password" required:"true"`
	LogLevel       string `envconfig:"log_level" default:"INFO"`
	Port           string `envconfig:"port" default:"3000"`
}

func ConfigLoad() (Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
