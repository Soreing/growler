package repos

import (
	"encoding/json"
	"time"

	"github.com/Soreing/growler/domain/common"
	discorddom "github.com/Soreing/growler/domain/domains/discord"
	"github.com/Soreing/growler/domain/general/cntxt"
	"github.com/Soreing/growler/infra/clients/discord"
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
)

type DiscordRepository struct {
	discl *discord.Client
	memch *cache.Cache
}

func NewDiscordRepository(
	discl *discord.Client,
	memch *cache.Cache,
) *DiscordRepository {
	return &DiscordRepository{
		discl: discl,
		memch: memch,
	}
}

var _ discorddom.IRepository = (*DiscordRepository)(nil)

func (r *DiscordRepository) GetDirectMessageChannel(
	ctx cntxt.IContext,
	userId string,
) (channelId string, err error) {
	lgr := ctx.Logger()
	lgr.Info(
		"creating dm channel",
		zap.String("userId", userId),
	)

	if val, ok := r.memch.Get(userId); ok {
		if str, ok := val.(string); ok {
			return str, nil
		}
	}

	lgr.Info("cache miss")

	dm, err := r.discl.CreateDMhannel(ctx, userId)
	if err != nil {
		lgr.Error("failed to create dm channel", zap.Error(err))
		return "", err
	}

	r.memch.Set(userId, dm.Id, time.Hour*24)

	return dm.Id, nil
}

func (r *DiscordRepository) CreateMessageBody(
	ctx cntxt.IContext,
	color int,
	format string,
	fields []common.StringPair,
) (body []byte, err error) {
	lgr := ctx.Logger()
	lgr.Info(
		"creating message body",
		zap.Int("color", color),
		zap.String("format", format),
		zap.Any("fields", fields),
	)

	embd := discord.RichEmbed{}
	if format != "" {
		err = json.Unmarshal([]byte(format), &embd)
		if err != nil {
			lgr.Warn(
				"failed to unmarshal embed format",
				zap.Error(err),
			)
		}
	}

	embd.Fields = make([]discord.EmbedField, 0, len(fields))
	for _, pair := range fields {
		fld := discord.EmbedField{
			Name:  pair.First,
			Value: pair.Second,
		}
		embd.Fields = append(embd.Fields, fld)
	}

	if embd.Color == nil {
		embd.Color = &color
	}

	msg := discord.CreateMessage{
		Embeds: []discord.RichEmbed{embd},
	}

	res, err := json.Marshal(msg)
	if err != nil {
		lgr.Error(
			"failed to marshal message into byte array",
			zap.Error(err),
		)
		return nil, err
	}

	return res, nil
}

func (r *DiscordRepository) SendMessage(
	ctx cntxt.IContext,
	channelId string,
	body []byte,
) (err error) {
	lgr := ctx.Logger()
	lgr.Info(
		"sending message",
		zap.String("chanelId", channelId),
		zap.ByteString("body", body),
	)

	err = r.discl.CreateMessage(ctx, channelId, body)
	if err != nil {
		lgr.Error("failed to send message", zap.Error(err))
		return err
	}

	return nil
}
