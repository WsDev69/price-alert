package postgres

import (
	"github.com/google/uuid"
	models "github.com/wsdev69/price-alert/api-service/v0.0.1/src/models/user"
)

func (q DBQuery) GetUser(ID uuid.UUID) (models.User, error) {
	u := models.User{ID: ID}
	err := q.Model(&u).WherePK().Select()
	if err != nil {
		return models.User{}, err
	}

	return u, nil
}

func (q DBQuery) SaveUser(u *models.User) (models.User, error) {
	_, err := q.Model(u).
		Returning("*").
		Insert()

	if err != nil {
		return models.User{}, err
	}

	return *u, nil
}

func (q DBQuery) IsUserExist(ID uuid.UUID) (bool, error) {
	return q.Model(&models.User{
		ID: ID,
	}).Exists()
}
