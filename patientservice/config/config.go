package config

import (
	"log"

	"github.com/caarlos0/env"
)

// Represents the configuration data for the service
type Config struct {
	PatientDetailServiceUrl string `env:"PATIENT_SERVICE_DETAILS_URL"`
}

func (c *Config) Bind() {
	err := env.Parse(c)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
}
