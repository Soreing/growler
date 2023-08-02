package common

const (
	RedColor   = 13828096
	GreenColor = 3912003
	BlueColor  = 33022

	FiringStatus   = "firing"
	ResolvedStatus = "resolved"

	AlertNameKey          = "alertname"
	AlertFolderKey        = "grafana_folder"
	TargetsMetadataKey    = "discord.targets"
	EmbedMetadataKey      = "discord.embed"
	DiscordMetadataPrefix = "discord."

	ContextKey    = "cntxt"
	TraceIdDigits = 32
	DefaultPort   = "8080"
)
