package alert

import (
	"context"
	"github.com/google/uuid"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/models/alert"
	httperrors "github.com/wsdev69/price-alert/api-service/v0.0.1/src/models/http-errors"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/persistence/postgres"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/services/quote"
	"sync"
)

type Service interface {
	Get(ctx context.Context, ID uuid.UUID) (alert.Alert, httperrors.HTTPError)
	GetAllByUserID(ctx context.Context, limit, offset int, userID uuid.UUID) ([]alert.Alert, httperrors.HTTPError)
	Delete(ctx context.Context, ID uuid.UUID) (alert.Alert, httperrors.HTTPError)
	Update(ctx context.Context, a alert.Alert) (alert.Alert, httperrors.HTTPError)
	Create(ctx context.Context, userID uuid.UUID, a alert.Alert) (alert.Alert, httperrors.HTTPError)
}

type service struct {
	pg    *postgres.Client
	quote quote.Service
}

var (
	srv  Service
	once = &sync.Once{}
)

// New returns service instance.
func New(pg *postgres.Client, quoteSrv quote.Service) Service {
	once.Do(func() {
		srv = service{
			pg:    pg,
			quote: quoteSrv,
		}
	})
	return srv
}
