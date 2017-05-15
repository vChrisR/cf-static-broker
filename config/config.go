package config

import "github.com/kelseyhightower/envconfig"

type Env struct {
	BrokerUsername string `envconfig:"broker_username" required:"true"`
	BrokerPassword string `envconfig:"broker_password" required:"true"`
	LogLevel       string `envconfig:"log_level" default:"INFO"`
	Port           string `envconfig:"port" default:"3000"`
}

func LoadEnv() (Env, error) {
	var env Env
	err := envconfig.Process("", &env)
	if err != nil {
		return Env{}, err
	}
	return env, nil
}
