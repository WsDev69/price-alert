package mail

import (
	"context"
	"github.com/wsdev69/price-alert/notififcation-service/v0.0.1/src/config"
	httperrors "github.com/wsdev69/price-alert/notififcation-service/v0.0.1/src/models/http-errors"

	"net/smtp"
	"sync"
)

type Service interface {
	Send(ctx context.Context,from, message string, to []string) httperrors.HTTPError
}

type unencryptedAuth struct {
	smtp.Auth
}

func (a unencryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	s := *server
	s.TLS = true
	return a.Auth.Start(&s)
}

type smtpService struct {
	auth unencryptedAuth
	host string
}

var (
	srv  Service
	once = &sync.Once{}
)

// New returns service instance.
func New(config config.SMTP) Service {
	once.Do(func() {
		auth := unencryptedAuth {
			smtp.PlainAuth(
				"",
				config.From,
				config.Password,
				"127.0.0.1",
			),
		}

		host := config.Host

		srv = smtpService{
			auth: auth,
			host: host,
		}
	})
	return srv
}
