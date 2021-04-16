package config

import (
	cfg "github.com/wsdev69/price-alert/quote-service/v0.0.1/src/config/redact"

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
		AppName       string        `json:"app_name"      envconfig:"APP_NAME"   default:"quote-service"`
		LogPreset     string        `json:"log_preset"    envconfig:"LOG_PRESET" default:"debug"`
		ListenURL     string        `json:"listen_url"    envconfig:"LISTEN_URL" default:":8084"`
		URLPrefix     string        `json:"url_prefix"    envconfig:"URL_PREFIX" default:"/api"`
		Redis         Redis         `json:"redis"`
		Kafka         Kafka         `json:"kafka"`
		CryptoCompare CryptoCompare `json:"crypto_compare"`

		HTTPTimeout string `json:"http_timeout"     envconfig:"HTTP_TIMEOUT"     default:"10s"`
	}

	// Redis defines configs for Redis Cache
	Redis struct {
		Addrs    string `json:"addrs"     envconfig:"REDIS_ADDRS"`
		PoolSize int    `json:"pool_size" envconfig:"REDIS_POOL_SIZE" default:"10"`
		Password string `json:"password"  envconfig:"REDIS_PASSWORD"  default:""`
	}

	Kafka struct {
		BootstrapServer string `json:"bootstrap_server" envconfig:"BOOTSTRAP_SERVER"`

	}

	CryptoCompare struct {
		APIKey string `json:"api_key" envconfig:"API_KEY"`
		URL    string `json:"url" envconfig:"API_KEY"`
	}
)
