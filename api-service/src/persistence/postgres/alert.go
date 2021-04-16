package postgres

import (
	"github.com/google/uuid"
	models "github.com/wsdev69/price-alert/api-service/v0.0.1/src/models/alert"
)

func (q DBQuery) SaveAlert(a *models.Alert) (models.Alert, error) {
	_, err := q.Model(a).
		Returning("*").
		Insert()

	if err != nil {
		return models.Alert{}, err
	}
	return *a, nil
}

func (q DBQuery) UpdateAlert(a *models.Alert) (models.Alert, error) {
	_, err := q.Model(a).
		Returning("*").
		WherePK().
		UpdateNotNull()

	if err != nil {
		return models.Alert{}, err
	}
	return *a, nil
}

func (q DBQuery) GetAlert(ID uuid.UUID) (a models.Alert, err error) {
	err = q.Model(&a).Where("id = ?", ID).Select("*")
	if err != nil {
		return models.Alert{}, err
	}
	return a, nil
}

func (q DBQuery) GetAlertByPriceAndCurrency(fromSymbol, toSymbol string, price float64) (a []*models.Alert, err error) {
	err = q.Model(&a).
		Where("from_symbol = ?", fromSymbol).
		Where("to_symbol = ?", toSymbol).
		Where("price >= ?", price).
		Select()

	if err != nil {
		return nil, err
	}

	return a, nil
}

func (q DBQuery) DeleteAlert(ID uuid.UUID) (a models.Alert, err error) {
	a.ID = ID
	_, err = q.Model(&a).
		Returning("*").
		WherePK().
		Delete()

	if err != nil {
		return models.Alert{}, err
	}
	return a, nil
}
