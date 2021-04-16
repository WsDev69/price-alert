package alert

import (
	"github.com/google/uuid"
)

type Alert struct {
	tableName struct{} `sql:"alerts" pg:",discard_unknown_columns"` // nolint

	ID         uuid.UUID `json:"id"           pg:"id,pk"`
	FromSymbol string    `json:"from_symbol"`
	ToSymbol   string    `json:"to_symbol"`
	Price      float64   `json:"price"`

	UserID uuid.UUID `json:"user_id" pg:"fk:user_id"`
}
