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
			if name == "" {
				continue
			}

			uv := getter.Get(name)

			switch field.Kind() {
			case reflect.Int8, reflect.Int, reflect.Int64:
				if i, err := strconv.Atoi(uv); err == nil {
					field.SetInt(int64(i))
				}
			case reflect.Uint8, reflect.Uint, reflect.Uint64:
				if i, err := strconv.Atoi(uv); err == nil {
					field.SetUint(uint64(i))
				}
			case reflect.String:
				field.SetString(uv)
			default:
				panic(fmt.Sprintf("kind (%s) not supported", elem.Kind()))
			}

		}

	}

	return nil
}
