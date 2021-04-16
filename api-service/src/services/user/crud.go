package user

import (
	"context"
	"fmt"
	"github.com/go-pg/pg"
	"github.com/google/uuid"
	httperrors "github.com/wsdev69/price-alert/api-service/v0.0.1/src/models/http-errors"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/models/user"
)

func (srv service) Get(ctx context.Context, ID uuid.UUID) (user.User, httperrors.HTTPError) {
	u, err := srv.pg.QueryContext(ctx).GetUser(ID)
	if err != nil {
		if err == pg.ErrNoRows {
			return user.User{}, httperrors.NewNotFoundError(err, fmt.Sprintf("alert ID %s", ID.String()))
		}
		return user.User{}, httperrors.NewInternalServerError(err)
	}

	return u, nil
}

func (srv service) Create(ctx context.Context, u user.User) (user.User, httperrors.HTTPError) {
	tx, err := srv.pg.NewTXContext(ctx)
	if err != nil {
		return user.User{}, httperrors.NewInternalServerError(err)
	}

	newU, errC := tx.SaveUser(&u)
	if errC != nil {
		return user.User{}, httperrors.NewInternalServerError(errC)
	}

	if err := tx.Commit(); err != nil {
		return user.User{}, httperrors.NewInternalServerError(err)
	}

	return newU, nil
}
