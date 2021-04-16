package user

import (
	"context"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/models/user"

	"github.com/google/uuid"
	httperrors "github.com/wsdev69/price-alert/api-service/v0.0.1/src/models/http-errors"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/persistence/postgres"
	"sync"
)

type Service interface {
	Get(ctx context.Context, ID uuid.UUID) (user.User, httperrors.HTTPError)
	Create(ctx context.Context, a user.User) (user.User, httperrors.HTTPError)
}

type service struct {
	pg *postgres.Client
}

var (
	srv  Service
	once = &sync.Once{}
)

// New returns service instance.
func New(pg *postgres.Client) Service {
	once.Do(func() {
		srv = service{
			pg: pg,
		}
	})
	return srv
}
