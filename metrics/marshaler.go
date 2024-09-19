package metrics

import (
	"fmt"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

var (
	marshaler   pmetric.Marshaler
	unmarshaler pmetric.Unmarshaler
)

func InitMarshaler(encoding string) error {
	switch encoding {
	case "json":
		marshaler = &pmetric.JSONMarshaler{}
		unmarshaler = &pmetric.JSONUnmarshaler{}
	case "proto":
		marshaler = &pmetric.ProtoMarshaler{}
		unmarshaler = &pmetric.ProtoUnmarshaler{}
	default:
		return fmt.Errorf("unknown metrics encoding: %s", encoding)
	}
	return nil
}
