package interactions

import (
	"encoding/json"
	"fmt"
	"time"
)

type InteractionType int

const (
	_ InteractionType = iota
	Ping
	ApplicationCommand
)

type InteractionResponseType int

const (
	_ InteractionResponseType = iota
	Pong
	Acknowledge
	ChannelMessage
	ChannelMessageWithSource
	AcknowledgeWithSource
)

type InteractionResponseFlags int64

const Ephemeral InteractionResponseFlags = 1 << 6

type Data struct {
	Type   InteractionType `json:"type"`
	Token  string          `json:"token"`
	Member struct {
		User struct {
			ID            string `json:"id"`
			Username      string `json:"username"`
			Avatar        string `json:"avatar"`
			Discriminator string `json:"discriminator"`
			PublicFlags   int64  `json:"public_flags"`
		} `json:"user"`
		Roles        []string  `json:"roles"`
		PremiumSince time.Time `json:"premium_since"`
		Permissions  string    `json:"permissions"`
		Pending      bool      `json:"pending"`
		Nick         string    `json:"nick"`
		Mute         bool      `json:"mute"`
		JoinedAt     time.Time `json:"joined_at"`
		IsPending    bool      `json:"is_pending"`
		Deaf         bool      `json:"deaf"`
	} `json:"member"`
	ID      string `json:"id"`
	GuildID string `json:"guild_id"`
	Data    struct {
		Options []ApplicationCommandInteractionDataOption `json:"options"`
		Name    string                                    `json:"name"`
		ID      string                                    `json:"id"`
	} `json:"data"`
	ChannelID string `json:"channel_id"`
}

func (data *Data) ResponseURL() string {
	return fmt.Sprintf("https://discord.com/api/v8/interactions/%s/%s/callback", data.ID, data.Token)
}

type ApplicationCommandInteractionDataOption struct {
	Name    string                                    `json:"name"`
	Value   interface{}                               `json:"value,omitempty"`
	Options []ApplicationCommandInteractionDataOption `json:"options,omitempty"`
}

type InteractionResponse struct {
	Type InteractionResponseType                    `json:"type"`
	Data *InteractionApplicationCommandCallbackData `json:"data,omitempty"`
}

type InteractionApplicationCommandCallbackData struct {
	TTS             *bool            `json:"tts,omitempty"`
	Content         string           `json:"content"`
	Flags         int           	 `json:"flags,omitempty"`
	Embeds          []Embed `json:"embeds,omitempty"`
	AllowedMentions json.Unmarshaler `json:"allowed_mentions,omitempty"`
}


type Embed struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Color       int       `json:"color"`
	Footer      Footer    `json:"footer"`
	Thumbnail   Thumbnail `json:"thumbnail"`
	Fields      []Field   `json:"fields"`
}

type Footer struct {
	IconURL string `json:"icon_url"`
	Text    string `json:"text"`
}

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

type Thumbnail struct {
	URL string `json:"url"`
}
