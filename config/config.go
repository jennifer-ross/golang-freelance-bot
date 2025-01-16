package config

import (
	"context"
	"golang-freelance-bot/logger"
	"os"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigdotenv"
	"github.com/cristalhq/aconfig/aconfigyaml"
	"github.com/rs/zerolog"
)

type Config struct {
	Telegram TelegramConfig
	Redis    RedisConfig
	Postgres PostgresConfig
}

type contextKey string

const configKey contextKey = "config"

// Load loads configuration from files and environment variables
func Load(files []string) *Config {
	var (
		cfg Config
		log *zerolog.Logger
	)

	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		// feel free to skip some steps :)
		// SkipDefaults: true,
		// SkipFiles:    true,
		SkipEnv:            false,
		SkipFlags:          true,
		EnvPrefix:          "",
		FlagPrefix:         "",
		Files:              files,
		AllFieldRequired:   false,
		AllowUnknownFields: true,
		FileDecoders: map[string]aconfig.FileDecoder{
			".yml": aconfigyaml.New(),
			".env": aconfigdotenv.New(),
		},
	})

	log = logger.Get("app")
	err := loader.Load()

	if err != nil {
		log.Error().Err(err).Msg("Cannot load configuration")
		os.Exit(2)
	}

	return &cfg
}

// New create config and embed it into the context
func New(ctx context.Context, files []string) (context.Context, *Config) {
	cfg := Load(files)
	// Embed the configuration into the context
	return context.WithValue(ctx, configKey, &cfg), cfg
}

// RetrieveConfig extracts the configuration from the context
func RetrieveConfig(ctx context.Context) *Config {
	if cfg, ok := ctx.Value(configKey).(*Config); ok {
		return cfg
	}
	return nil
}
