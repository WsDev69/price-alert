package config

import (
	cfg "github.com/wsdev69/price-alert/notififcation-service/v0.0.1/src/config/redact"

	"log"
)

var (
	// Config is a package variable, which is populated during init() execution and shared to whole application
	Config Configuration
)

// Load initialize config
func Load(configPath string) error {
	err := cfg.Load(&Config, configPath)
	log.Println(Config.String())
	return err
}

type (
	// Configuration options
	Configuration struct {
		AppName   string `json:"app_name"      envconfig:"APP_NAME"   default:"notification-service"`
		LogPreset string `json:"log_preset"    envconfig:"LOG_PRESET" default:"debug"`
		ListenURL string `json:"listen_url"    envconfig:"LISTEN_URL" default:":8085"`
		URLPrefix string `json:"url_prefix"    envconfig:"URL_PREFIX" default:"/api"`

		SMTP SMTP `json:"smtp"`
	}

	SMTP struct {
		Host     string `json:"host"     envconfig:"SMTP_HOST"`
		Password string `json:"password" envconfig:"SMTP_HOST"`
		From     string `json:"from"     envconfig:"SMTP_HOST"`
	}
)
