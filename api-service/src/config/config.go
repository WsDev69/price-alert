package config

import (
	"log"

	cfg "github.com/wsdev69/price-alert/api-service/v0.0.1/src/config/redact"
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
		AppName             string              `json:"app_name"      envconfig:"APP_NAME"   default:"api-service"`
		LogPreset           string              `json:"log_preset"    envconfig:"LOG_PRESET" default:"debug"`
		ListenURL           string              `json:"listen_url"    envconfig:"LISTEN_URL" default:":8080"`
		URLPrefix           string              `json:"url_prefix"    envconfig:"URL_PREFIX" default:"/api"`
		Postgres            Postgres            `json:"postgres"`
		PostgresTest        Postgres            `json:"postgres_test"`
		Kafka               Kafka               `json:"kafka"`
		QuoteService        QuoteService        `json:"quote_service"`
		NotificationService NotificationService `json:"notification_service"`

		PaginationLimit int    `json:"pagination_limit" envconfig:"PAGINATION_LIMIT" default:"100"`
		HTTPTimeout     string `json:"http_timeout"     envconfig:"HTTP_TIMEOUT"     default:"10s"`
	}

	// Postgres options
	Postgres struct {
		Host         string `json:"host"             envconfig:"POSTGRES_HOST"           default:"localhost"`
		Port         string `json:"port"             envconfig:"POSTGRES_PORT"           default:"5432"`
		DBName       string `json:"db_name"          envconfig:"POSTGRES_DB"             default:"orbis-api-db"`
		User         string `json:"user"             envconfig:"POSTGRES_USER"           default:"postgres"`
		Password     string `json:"password"         envconfig:"POSTGRES_PASSWORD"       default:"12345"`
		PoolSize     int    `json:"pool_size"        envconfig:"POSTGRES_POOL_SIZE"      default:"10"`
		MaxRetries   int    `json:"max_retries"      envconfig:"POSTGRES_MAX_RETRIES"    default:"5"`
		ReadTimeout  string `json:"read_timeout"     envconfig:"POSTGRES_READ_TIMEOUT"   default:"10s"`
		WriteTimeout string `json:"write_timeout"    envconfig:"POSTGRES_WRITE_TIMEOUT"  default:"10s"`
	}

	Kafka struct {
		BootstrapServer string `json:"bootstrap_server" envconfig:"KAFKA_BOOTSTRAP_SERVER"`
		GroupID         string `json:"group_id" envconfig:"KAFKA_GROUP_ID"`
	}

	QuoteService struct {
		URL string `json:"url" envconfig:"QUOTE_SERVICE_URL"`
	}
	NotificationService struct {
		URL string `json:"url" envconfig:"NOTIFICATION_SERVICE_URL"`
	}
)
