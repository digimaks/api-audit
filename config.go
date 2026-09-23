// SPDX-License-Identifier: EUPL-1.2

package api

import (
	"azugo.io/azugo/config"
	"azugo.io/core/validation"
	"azugo.io/opentelemetry"
	"github.com/digimaks/go-idauth"
	"github.com/lx-lib/lx-go-jsondb"
	"github.com/spf13/viper"
)

// Configuration represents the configuration for the application.
type Configuration struct {
	*config.Configuration `mapstructure:",squash"`
	Telemetry             *opentelemetry.Configuration `mapstructure:"telemetry"`
	Auth                  *idauth.Configuration        `mapstructure:"idauth"`
	Postgres              *jsondb.Configuration        `mapstructure:"postgres"`
}

// NewConfiguration returns a new configuration.
func NewConfiguration() *Configuration {
	return &Configuration{
		Configuration: config.New(),
	}
}

// ServerCore returns the core configuration.
func (c *Configuration) ServerCore() *config.Configuration {
	return c.Configuration
}

// Bind configuration to viper.
func (c *Configuration) Bind(_ string, v *viper.Viper) {
	c.Configuration.Bind("", v)
	c.Auth = config.Bind(c.Auth, "idauth", v)
	c.Postgres = config.Bind(c.Postgres, "postgres", v)
	c.Telemetry = config.Bind(c.Telemetry, "telemetry", v)
}

// Validate application configuration.
func (c *Configuration) Validate(validate *validation.Validate) error {
	if err := c.Auth.Validate(validate); err != nil {
		return err
	}

	if err := c.Postgres.Validate(validate); err != nil {
		return err
	}

	if err := c.Telemetry.Validate(validate); err != nil {
		return err
	}

	return nil
}
