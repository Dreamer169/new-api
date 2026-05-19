package system_setting

import "github.com/QuantumNous/new-api/setting/config"

type DiscordSettings struct {
	Enabled         bool   `json:"enabled"`
	ClientId        string `json:"client_id"`
	ClientSecret    string `json:"client_secret"`
	RequiredGuildId string `json:"required_guild_id"`
	RequiredRoleId  string `json:"required_role_id"`
}

var defaultDiscordSettings = DiscordSettings{}

func init() {
	config.GlobalConfig.Register("discord", &defaultDiscordSettings)
}

func GetDiscordSettings() *DiscordSettings {
	return &defaultDiscordSettings
}
