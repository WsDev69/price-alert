package user

import "github.com/google/uuid"

type User struct {
	tableName struct{} `sql:"users" pg:",discard_unknown_columns"` // nolint

	ID    uuid.UUID `json:"id" pg:"id,pk"`
	Email string    `json:"email"`
}
