package postgres

import (
	models "github.com/wsdev69/price-alert/api-service/v0.0.1/src/models/user"
	"reflect"
	"testing"
)

func TestDBQuery_SaveUser(t *testing.T) {
	type fields struct {
		DBModel   DBModel
		completed bool
	}
	type args struct {
		u *models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := DBQuery{
				DBModel:   tt.fields.DBModel,
				completed: tt.fields.completed,
			}
			got, err := q.SaveUser(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SaveUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
