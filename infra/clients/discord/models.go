package discord

type CreateMessage struct {
	Content *string     `json:"content,omitempty"`
	Embeds  []RichEmbed `json:"embeds,omitempty"`
}

type CreateDM struct {
	RecipientId string `json:"recipient_id"`
}

type DMChannel struct {
	Id string `json:"id"`
}

type EmbedFooter struct {
	Text         string  `json:"text"`
	IconUrl      *string `json:"icon_url,omitempty"`
	ProxyIconUrl *string `json:"proxy_icon_url,omitempty"`
}

type EmbedImage struct {
	Url      string  `json:"url"`
	ProxyUrl *string `json:"proxy_url,omitempty"`
	Height   *int    `json:"height,omitempty"`
	Width    *int    `json:"width,omitempty"`
}

type EmbedThumbnail struct {
	Url      string  `json:"url"`
	ProxyUrl *string `json:"proxy_url,omitempty"`
	Height   *int    `json:"height,omitempty"`
	Width    *int    `json:"width,omitempty"`
}

type EmbedVideo struct {
	Url      string  `json:"url"`
	ProxyUrl *string `json:"proxy_url,omitempty"`
	Height   *int    `json:"height,omitempty"`
	Width    *int    `json:"width,omitempty"`
}

type EmbedProvider struct {
	Name *string `json:"name,omitempty"`
	Url  *string `json:"url,omitempty"`
}

type EmbedAuthor struct {
	Name         string  `json:"name"`
	Url          *string `json:"url,omitempty"`
	IconUrl      *string `json:"icon_url,omitempty"`
	ProxyIconUrl *string `json:"proxy_icon_url,omitempty"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline *bool  `json:"inline,omitempty"`
}

type RichEmbed struct {
	Title       *string         `json:"title,omitempty"`
	Description *string         `json:"description,omitempty"`
	Url         *string         `json:"url,omitempty"`
	Color       *int            `json:"color,omitempty"`
	Footer      *EmbedFooter    `json:"footer,omitempty"`
	Image       *EmbedImage     `json:"image,omitempty"`
	Thumbnail   *EmbedThumbnail `json:"thumbnail,omitempty"`
	Video       *EmbedVideo     `json:"video,omitempty"`
	Provider    *EmbedProvider  `json:"provider,omitempty"`
	Author      *EmbedAuthor    `json:"author,omitempty"`
	Fields      []EmbedField    `json:"fields,omitempty"`
}
