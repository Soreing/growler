package usecases

import (
	"strings"

	"github.com/Soreing/growler/domain/common"
	"github.com/Soreing/growler/domain/general/cntxt"
	"go.uber.org/zap"
)

func (u *UseCases) PublishAlert(
	ctx cntxt.IContext,
	status string,
	labels map[string]string,
	annotations map[string]string,
) (err error) {
	lgr := ctx.Logger()
	lgr.Info(
		"publishing alert",
		zap.String("status", status),
		zap.Any("labels", labels),
		zap.Any("annotations", annotations),
	)

	// Targets
	targets, ok := labels[common.TargetsMetadataKey]
	if !ok || len(targets) == 0 {
		lgr.Warn("no targets specified")
		return
	}

	// Format
	format := annotations[common.EmbedMetadataKey]

	// Color
	var color int
	switch status {
	case common.FiringStatus:
		color = common.RedColor
	case common.ResolvedStatus:
		color = common.GreenColor
	default:
		color = common.BlueColor
	}

	// Fields
	fields := make([]common.StringPair, 0, len(labels))
	pfx, pln := common.DiscordMetadataPrefix, len(common.DiscordMetadataPrefix)
	for k, v := range labels {
		if (len(k) >= pln && k[:pln] == pfx) ||
			k == "alertname" || k == "grafana_folder" {
			continue
		}
		fields = append(fields, common.StringPair{
			First:  k,
			Second: v,
		})
	}

	// Create Message
	msg, err := u.dscrep.CreateMessageBody(
		ctx, color, format, fields,
	)
	if err != nil {
		lgr.Error("failed to create alert message", zap.Error(err))
		return err
	}

	targetIds := strings.Split(targets, ",")
	for _, targetId := range targetIds {
		chanid := ""
		if strings.HasPrefix(targetId, "c:") {
			chanid = strings.TrimPrefix(targetId, "c:")
		} else if strings.HasPrefix(targetId, "u:") {
			usrid := strings.TrimPrefix(targetId, "u:")
			chanid, err = u.dscrep.GetDirectMessageChannel(ctx, usrid)
			if err != nil {
				lgr.Error("failed to get DM channel id", zap.Error(err))
				continue
			}
		} else {
			lgr.Warn("invalid target", zap.String("target", targetId))
		}

		if err := u.dscrep.SendMessage(ctx, chanid, msg); err != nil {
			lgr.Error("failed to send message", zap.Error(err))
		}
	}

	return nil
}
