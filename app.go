// SPDX-License-Identifier: EUPL-1.2

package api

import (
	"azugo.io/azugo"
	"azugo.io/azugo/server"
	"azugo.io/core/instrumenter"
	"azugo.io/opentelemetry"
	"github.com/lx-lib/lx-go-jsondb"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// App is the application instance.
type App struct {
	*azugo.App

	config *Configuration
	store  jsondb.Store
}

// New returns a new application instance.
func New(cmd *cobra.Command, version string) (*App, error) {
	config := NewConfiguration()

	a, err := server.New(
		cmd,
		server.Options{
			AppName:       "Audit API",
			AppVer:        version,
			Configuration: config,
		})
	if err != nil {
		return nil, err
	}

	a.RouterOptions().CORS.SetHeaders("Accept", "Accept-Language", "Content-Type", "Authorization")

	store, _, err := jsondb.New(a.App, config.Postgres)
	if err != nil {
		return nil, err
	}

	tel, err := opentelemetry.Use(
		a,
		config.Telemetry,
		opentelemetry.InstrumentationRecorder("db", jsondb.Tracing, jsondb.InstrumentationExec),
	)
	if err != nil {
		return nil, err
	}

	if err := a.AddTask(tel); err != nil {
		return nil, err
	}

	app := &App{
		App:    a,
		config: config,
		store:  store,
	}

	app.Instrumentation(instrumenter.CombinedInstrumenter(app.Instrumenter(), app.storeInstrumenter()))

	return app, nil
}

func (a *App) Start() error {
	if err := a.Store().Start(a.BackgroundContext()); err != nil {
		a.Log().Warn("store failed to start", zap.Error(err))
	}

	return a.App.Start()
}

// Config returns application configuration.
// Panics if configuration is not loaded.
func (a *App) Config() *Configuration {
	if a.config == nil || !a.config.Ready() {
		panic("configuration is not loaded")
	}

	return a.config
}

// RemoteFileStore returns remote file store.
func (a *App) Store() jsondb.Store {
	return a.store
}
