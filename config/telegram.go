package config

type TelegramConfig struct {
	BotToken string `env_name:"TELEGRAM_BOT_TOKEN" required:"true"`
}
