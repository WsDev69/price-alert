package postgres

import (
	"github.com/google/uuid"
	models "github.com/wsdev69/price-alert/api-service/v0.0.1/src/models/alert"
	"reflect"
	"testing"
)

func TestDBQuery_DeleteAlert(t *testing.T) {
	type fields struct {
		DBModel   DBModel
		completed bool
	}
	type args struct {
		ID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantA   models.Alert
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
			gotA, err := q.DeleteAlert(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteAlert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotA, tt.wantA) {
				t.Errorf("DeleteAlert() gotA = %v, want %v", gotA, tt.wantA)
			}
		})
	}
}

func TestDBQuery_GetAlert(t *testing.T) {
	type fields struct {
		DBModel   DBModel
		completed bool
	}
	type args struct {
		ID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantA   models.Alert
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
			gotA, err := q.GetAlert(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAlert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotA, tt.wantA) {
				t.Errorf("GetAlert() gotA = %v, want %v", gotA, tt.wantA)
			}
		})
	}
}

func TestDBQuery_SaveAlert(t *testing.T) {
	type fields struct {
		DBModel   DBModel
		completed bool
	}
	type args struct {
		a *models.Alert
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Alert
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
			got, err := q.SaveAlert(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveAlert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SaveAlert() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBQuery_UpdateAlert(t *testing.T) {
	type fields struct {
		DBModel   DBModel
		completed bool
	}
	type args struct {
		a *models.Alert
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Alert
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
			got, err := q.UpdateAlert(tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateAlert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateAlert() got = %v, want %v", got, tt.want)
			}
		})
	}
}
