package mgos

import (
	"fmt"
	"reflect"
	"strconv"
)

// Getter like url.Values
type Getter interface {
	Get(string) string
}

// FromGetter make strunct from value
func FromGetter(getter Getter, dest interface{}) error {

	switch dest.(type) {
	default:
		return fmt.Errorf("dest is not struct")
	case interface{}:
		elem := reflect.ValueOf(dest).Elem()

		for i := 0; i < elem.NumField(); i++ {
			field := elem.Field(i)
			typeField := elem.Type().Field(i)
			name := typeField.Tag.Get("mgos")
			uv := getter.Get(name)

			switch field.Interface().(type) {
			case int8, uint8, int, uint, int64, uint64:
				// Insert int value when no errors.
				if i, err := strconv.Atoi(uv); err == nil {
					field.SetInt(int64(i))
				}
			case string:
				field.SetString(uv)
			default:
				panic("mgos not supported")
			}

		}

	}

	return nil
}
