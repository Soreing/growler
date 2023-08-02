package usecases

import (
	"github.com/Soreing/growler/domain/common"
	"github.com/Soreing/growler/domain/general/cntxt"
)

func (u *UseCases) CreateContext() cntxt.IContext {
	return u.ctxf.Create(
		u.uidr.GetHexString(common.TraceIdDigits),
	)
}
