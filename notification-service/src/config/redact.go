package config

import (
	"encoding/json"
	"fmt"
	config "github.com/wsdev69/price-alert/notififcation-service/v0.0.1/src/config/redact"
	"reflect"


)

var (
	redactedFields = []string{"password"}
	redactor       = config.NewRedactor(redactedFields)
)

// aliases used to prevent recursion call of methods
type (
	aliasConfiguration Configuration
	aliasSMTP      SMTP
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
func (p *SMTP) MarshalJSON() ([]byte, error) {
	var tmp aliasSMTP
	if err := cloneRedact(&tmp, p); err != nil {
		return nil, err
	}
	return json.Marshal(tmp)
}

// String - %s, %v, %+v
func (p *SMTP) String() string {
	var tmp aliasSMTP
	if err := cloneRedact(&tmp, p); err != nil {
		return ""
	}
	return fmt.Sprintf("%+v", tmp)
}

// GoString - %#v
func (p *SMTP) GoString() string {
	var tmp aliasSMTP
	if err := cloneRedact(&tmp, p); err != nil {
		return ""
	}
	return fmt.Sprintf("%#v", tmp)
}
