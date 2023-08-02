package infra

import (
	discorddom "github.com/Soreing/growler/domain/domains/discord"
	icntxt "github.com/Soreing/growler/domain/general/cntxt"
	"github.com/Soreing/growler/domain/general/uids"
	"github.com/Soreing/growler/infra/clients/discord"
	"github.com/Soreing/growler/infra/cntxt"
	"github.com/Soreing/growler/infra/config"
	"github.com/Soreing/growler/infra/logger"
	"github.com/Soreing/growler/infra/memcache"
	"github.com/Soreing/growler/infra/repos"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var DependencySet = wire.NewSet(
	logger.NewLogger,
	config.NewConfigContext,
	config.NewAppConfigs,

	// Generals
	cntxt.NewFactory,
	wire.Bind(
		new(icntxt.IFactory),
		new(*cntxt.Factory),
	),

	// Repositories
	repos.NewUidsRepository,
	wire.Bind(
		new(uids.IRepository),
		new(*repos.UidsRepository),
	),
	repos.NewDiscordRepository,
	wire.Bind(
		new(discorddom.IRepository),
		new(*repos.DiscordRepository),
	),

	// Clients
	config.NewDiscordClientOptions,
	discord.NewClient,

	memcache.NewMemoryCache,
	NewInfrastructure,
)

type Infrastructure struct {
	lgr *zap.Logger
}

func NewInfrastructure(
	lgr *zap.Logger,
) *Infrastructure {
	return &Infrastructure{
		lgr: lgr,
	}
}

func (i *Infrastructure) Start() {
}

func (i *Infrastructure) Stop() {
	i.lgr.Sync()
}
