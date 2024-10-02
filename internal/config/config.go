package config

type Config struct {
	BotUsername string    `yaml:"botUsername"`
	Port        int       `yaml:"port"`
	Webhooks    []Webhook `yaml:"webhooks"`
}

type Webhook struct {
	SecretToken     string  `yaml:"secretToken,omitempty"`
	ListenPath      string  `yaml:"listenPath"`
	SlackWebhookUrl *string `yaml:"slackWebhookUrl,omitempty"`
}
