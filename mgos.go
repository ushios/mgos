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

// Scanner scan data
type Scanner interface {
	Scan(interface{}) error
}

// FromGetter make strunct from value
func FromGetter(getter Getter, dest interface{}) error {

	switch dest.(type) {
	default:
		return fmt.Errorf("dest is not struct")
	case interface{}:
		elem := reflect.ValueOf(dest).Elem()

		for i := 0; i < elem.NumField(); i++ {
			fieldValue := elem.Field(i)
			refType := elem.Type()
			typeField := refType.Field(i)
			name := typeField.Tag.Get("mgos")
			if name == "" {
				continue
			}

			uv := getter.Get(name)
			if uv == "" {
				continue
			}
			switch fieldValue.Kind() {
			case reflect.Int8, reflect.Int, reflect.Int64:
				if i, err := strconv.Atoi(uv); err == nil {
					fieldValue.SetInt(int64(i))
				}
			case reflect.Uint8, reflect.Uint, reflect.Uint64:
				if i, err := strconv.Atoi(uv); err == nil {
					fieldValue.SetUint(uint64(i))
				}
			case reflect.String:
				fieldValue.SetString(uv)
			case reflect.Bool:
				if uv != "" && uv != "0" {
					fieldValue.SetBool(true)
				}
			case reflect.Struct:
				if err := setToStruct(fieldValue.Interface(), uv, typeField.Name); err != nil {
					return err
				}
			case reflect.Ptr:
				if fieldValue.IsNil() {
					fieldValue.Set(reflect.New(typeField.Type.Elem()))
				}
				if err := setToStruct(fieldValue.Interface(), uv, typeField.Name); err != nil {
					return err
				}
			default:
				panic(fmt.Sprintf("kind elem(%s) field(%s) not supported", elem.Kind(), fieldValue.Kind()))
			}

		}

	}

	return nil
}

func setToStruct(i, v interface{}, name string) error {
	scanner, ok := i.(Scanner)
	if !ok {
		return fmt.Errorf("%s is not Scanner", name)
	}

	return scanner.Scan(v)
}
