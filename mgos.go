package mgos

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

// FromURLValues make strunct from value
func FromURLValues(u url.Values, dest interface{}) error {

	switch dest.(type) {
	default:
		return fmt.Errorf("dest is not struct")
	case interface{}:
		vo := reflect.ValueOf(dest)
		// sv := reflect.Indirect(vo)
		val := vo.Elem()

		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			typeField := val.Type().Field(i)
			name := typeField.Tag.Get("mgos")

			uv := u.Get(name)
			switch field.Interface().(type) {
			case int8, uint8, int, uint, int64, uint64:
				if i, err := strconv.Atoi(uv); err == nil {
					vo.SetInt(int64(i))
				} else {
					vo.SetInt(0)
				}

			case string:
				vo.SetString(uv)
			default:
				panic("mgos not supported")
			}

		}

	}

	return nil
}
