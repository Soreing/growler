package usecases

import (
	"github.com/Soreing/growler/domain/domains/discord"
	"github.com/Soreing/growler/domain/general/cntxt"
	"github.com/Soreing/growler/domain/general/uids"
)

type UseCases struct {
	ctxf   cntxt.IFactory
	uidr   uids.IRepository
	dscrep discord.IRepository
}

func NewUseCases(
	ctxf cntxt.IFactory,
	uidr uids.IRepository,
	dscrep discord.IRepository,
) *UseCases {
	return &UseCases{
		ctxf:   ctxf,
		uidr:   uidr,
		dscrep: dscrep,
	}
}
