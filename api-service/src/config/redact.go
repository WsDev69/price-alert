package config

import (
	"encoding/json"
	"fmt"
	"reflect"

	config "github.com/wsdev69/price-alert/api-service/v0.0.1/src/config/redact"
)

var (
	redactedFields = []string{"password"}
	redactor       = config.NewRedactor(redactedFields)
)

// aliases used to prevent recursion call of methods
type (
	aliasConfiguration Configuration
	aliasPostgres      Postgres
)

func cloneRedact(dst, src interface{}) error {
	err := config.Clone(dst, src)
	redactor.Redact(reflect.ValueOf(dst))
	return err
}

// MarshalJSON redact sensitive data
func (c *Configuration) MarshalJSON() ([]byte, error) {
	var tmp aliasConfiguration
	if err := cloneRedact(&tmp, c); err != nil {
		return nil, err
	}
	return json.Marshal(tmp)
}

// String - %s, %v, %+v
func (c *Configuration) String() string {
	var tmp aliasConfiguration
	if err := cloneRedact(&tmp, c); err != nil {
		return ""
	}
	return fmt.Sprintf("%+v", tmp)
}

// GoString - %#v
func (c *Configuration) GoString() string {
	var tmp aliasConfiguration
	if err := cloneRedact(&tmp, c); err != nil {
		return ""
	}
	return fmt.Sprintf("%#v", tmp)
}

// MarshalJSON redact sensitive data
func (p *Postgres) MarshalJSON() ([]byte, error) {
	var tmp aliasPostgres
	if err := cloneRedact(&tmp, p); err != nil {
		return nil, err
	}
	return json.Marshal(tmp)
}

// String - %s, %v, %+v
func (p *Postgres) String() string {
	var tmp aliasPostgres
	if err := cloneRedact(&tmp, p); err != nil {
		return ""
	}
	return fmt.Sprintf("%+v", tmp)
}

// GoString - %#v
func (p *Postgres) GoString() string {
	var tmp aliasPostgres
	if err := cloneRedact(&tmp, p); err != nil {
		return ""
	}
	return fmt.Sprintf("%#v", tmp)
}
