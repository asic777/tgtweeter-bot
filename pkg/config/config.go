package config

import "github.com/spf13/viper"

type Messages struct {
	Responses
	Errors
}

type Responses struct {
	Start             string `mapstructure:"start"`
	AlreadyAuthorized string `mapstructure:"already_authorized"`
	UnknownCommand    string `mapstructure:"unknown_command"`
	LinkSaved         string `mapstructure:"link_saved"`
}

type Errors struct {
	Default      string `mapstructure:"default"`
	InvalidURL   string `mapstructure:"invalid_url"`
	UnableToSave string `mapstructure:"unable_to_save"`
}

type Config struct {
	TelegramToken     string
	PocketConsumerKey string
	AuthServerURL     string

	BotURL     string `mapstructure:"bot_url"`
	BoltDBFile string `mapstructure:"db_file"`

	Messages Messages
}

func Init() (*Config, error) {
	if err := setUpViper(); err != nil {
		return nil, err
	}

	var cfg Config
	err := unmarshal(&cfg)

	fromEnv(&cfg)

	return &cfg, err
}

func unmarshal(cfg *Config) error {
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("messages.response", &cfg.Messages.Responses); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("messages.error", &cfg.Messages.Errors); err != nil {
		return err
	}

	return nil
}

func fromEnv(cfg *Config) error {
	if err := viper.BindEnv("token"); err != nil {
		return err
	}
	cfg.TelegramToken = viper.GetString("token")

	if err := viper.BindEnv("consumer_key"); err != nil {
		return err
	}
	cfg.PocketConsumerKey = viper.GetString("consumer_key")

	return nil
}

func setUpViper() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	return viper.ReadInConfig()
}
