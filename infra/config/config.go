package config

import (
	"os"

	"github.com/Soreing/growler/domain/common"

	"github.com/Soreing/growler/infra/clients/discord"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type ConfigContext struct {
	lgr *zap.Logger
}

func NewConfigContext(
	lgr *zap.Logger,
) *ConfigContext {
	c := &ConfigContext{
		lgr: lgr,
	}
	c.LoadConfigCustom("./.env")
	return c
}

func (c *ConfigContext) LoadConfigCustom(loc string) {
	err := godotenv.Load(loc)
	if err != nil {
		c.lgr.Warn("failed to load .env file", zap.String("location", loc))
	}
}

type AppConfigs struct {
	Development bool
	PortNumber  string
}

func NewAppConfigs(c *ConfigContext) *AppConfigs {
	cfg := &AppConfigs{}

	env := os.Getenv("ENV")
	if env == "development" {
		cfg.Development = true
	}

	cfg.PortNumber = os.Getenv("PORT")
	if cfg.PortNumber == "" {
		cfg.PortNumber = common.DefaultPort
	}

	return cfg
}

func NewDiscordClientOptions(c *ConfigContext) *discord.Options {
	opt := &discord.Options{}

	opt.BaseUrl = os.Getenv("DiscordBaseUrl")
	if opt.BaseUrl == "" {
		panic("discord base url is missing")
	}
	opt.Token = os.Getenv("DiscordToken")
	if opt.Token == "" {
		panic("discord token is missing")
	}

	return opt
}
