package alert

import (
	"context"
	"fmt"
	"github.com/go-pg/pg"
	"github.com/google/uuid"
	models "github.com/wsdev69/price-alert/api-service/v0.0.1/src/models/alert"
	httperrors "github.com/wsdev69/price-alert/api-service/v0.0.1/src/models/http-errors"
	"github.com/wsdev69/price-alert/api-service/v0.0.1/src/models/quote"
)

func (srv service) Get(ctx context.Context, ID uuid.UUID) (models.Alert, httperrors.HTTPError) {
	a, err := srv.pg.QueryContext(ctx).GetAlert(ID)
	if err != nil {
		if err == pg.ErrNoRows {
			return models.Alert{}, httperrors.NewNotFoundError(err, fmt.Sprintf("alert ID %s", ID.String()))
		}
		return models.Alert{}, httperrors.NewInternalServerError(err)
	}

	return a, nil
}

func (srv service) Delete(ctx context.Context, ID uuid.UUID) (models.Alert, httperrors.HTTPError) {
	tx, err := srv.pg.NewTXContext(ctx)
	if err != nil {
		return models.Alert{}, httperrors.NewInternalServerError(err)
	}

	a, errD := tx.DeleteAlert(ID)
	if errD != nil {
		if err == pg.ErrNoRows {
			return models.Alert{}, httperrors.NewNotFoundError(errD, fmt.Sprintf("alert ID %s", ID.String()))
		}
		return models.Alert{}, httperrors.NewInternalServerError(errD)
	}

	if err := tx.Commit(); err != nil {
		return models.Alert{}, httperrors.NewInternalServerError(err)
	}

	return a, nil
}

func (srv service) Update(ctx context.Context, a models.Alert) (models.Alert, httperrors.HTTPError) {
	tx, err := srv.pg.NewTXContext(ctx)
	if err != nil {
		return models.Alert{}, httperrors.NewInternalServerError(err)
	}
	updA, errU := tx.UpdateAlert(&a)
	if errU != nil {
		if err == pg.ErrNoRows {
			return models.Alert{}, httperrors.NewNotFoundError(errU, fmt.Sprintf("alert ID %s", a.ID.String()))
		}
		return models.Alert{}, httperrors.NewInternalServerError(errU)
	}

	req := quote.AddNewCurrencyRequest{
		FromSymbol: updA.FromSymbol,
		ToSymbol:   updA.ToSymbol,
	}
	if err := srv.quote.UpdateCurrency(req); err != nil {
		return models.Alert{}, httperrors.NewInternalServerError(err)
	}

	if err := tx.Commit(); err != nil {
		return models.Alert{}, httperrors.NewInternalServerError(err)
	}

	return updA, nil
}

func (srv service) Create(ctx context.Context, userID uuid.UUID, a models.Alert) (models.Alert, httperrors.HTTPError) {
	tx, err := srv.pg.NewTXContext(ctx)
	if err != nil {
		return models.Alert{}, httperrors.NewInternalServerError(err)
	}
	isExist, err := tx.IsUserExist(userID)
	if err != nil {
		return models.Alert{}, httperrors.NewInternalServerError(err)
	}
	if !isExist {
		return models.Alert{}, httperrors.NewNotFoundError(fmt.Errorf("user %s doesn't exist", userID.String()), "userID")
	}
	a.UserID = userID
	newA, errC := tx.SaveAlert(&a)
	if errC != nil {
		return models.Alert{}, httperrors.NewInternalServerError(errC)
	}

	req := quote.AddNewCurrencyRequest{
		FromSymbol: newA.FromSymbol,
		ToSymbol:   newA.ToSymbol,
	}

	if err := tx.Commit(); err != nil {
		return models.Alert{}, httperrors.NewInternalServerError(err)
	}

	if err := srv.quote.UpdateCurrency(req); err != nil {
		return models.Alert{}, httperrors.NewInternalServerError(err)
	}

	return newA, nil
}

func (srv service) GetAllByUserID(ctx context.Context, limit, offset int, userID uuid.UUID) (
	[]models.Alert, httperrors.HTTPError) {

	return nil, nil
}
