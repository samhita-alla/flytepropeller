// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package webhook

import (
	"encoding/json"
	"reflect"

	"fmt"

	"github.com/spf13/pflag"
)

// If v is a pointer, it will get its element value or the zero value of the element type.
// If v is not a pointer, it will return it as is.
func (Config) elemValueOrNil(v interface{}) interface{} {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		if reflect.ValueOf(v).IsNil() {
			return reflect.Zero(t.Elem()).Interface()
		} else {
			return reflect.ValueOf(v).Interface()
		}
	} else if v == nil {
		return reflect.Zero(t).Interface()
	}

	return v
}

func (Config) mustMarshalJSON(v json.Marshaler) string {
	raw, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return string(raw)
}

// GetPFlagSet will return strongly types pflags for all fields in Config and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg Config) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("Config", pflag.ExitOnError)
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "metrics-prefix"), defaultConfig.MetricsPrefix, "An optional prefix for all published metrics.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "certDir"), defaultConfig.CertDir, "Certificate directory to use to write generated certs. Defaults to /etc/webhook/certs/")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "listenPort"), defaultConfig.ListenPort, "The port to use to listen to webhook calls. Defaults to 9443")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "serviceName"), defaultConfig.ServiceName, "The name of the webhook service.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "secretName"), defaultConfig.SecretName, "Secret name to write generated certs to.")
	return cmdFlags
}
