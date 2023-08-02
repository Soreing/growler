package discord

import (
	"github.com/Soreing/growler/domain/common"
	"github.com/Soreing/growler/domain/general/cntxt"
)

type IRepository interface {
	GetDirectMessageChannel(
		ctx cntxt.IContext,
		userId string,
	) (channelId string, err error)

	CreateMessageBody(
		ctx cntxt.IContext,
		color int,
		format string,
		fields []common.StringPair,
	) (body []byte, err error)

	SendMessage(
		ctx cntxt.IContext,
		channelId string,
		body []byte,
	) (err error)
}
