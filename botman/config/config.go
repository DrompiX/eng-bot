package config

type BotManConfig struct {
	BotConfig struct {
		Token string `yaml:"token"`
		Debug bool   `yaml:"debug"`
	} `yaml:"bot_config"`
	ServerConfig struct {
		URL string `yaml:"url"`
	} `yaml:"server_config"`
}
