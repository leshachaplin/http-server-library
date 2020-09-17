package config

// if using go modules

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type config struct {
	Port             int    `env:"Port" envDefault:"1323"`
	GrpcPort         string `env:"GrpcPort" envDefault:"0.0.0.0:50051"`
}

func NewConfig() *config {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	return &cfg
}
