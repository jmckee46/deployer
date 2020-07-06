package serializers

import (
	"fmt"
	"reflect"
)

// CopyAndStringify not sure what this is used for...
func CopyAndStringify(rawObject interface{}, newObject interface{}) {
	rawValue := reflect.ValueOf(rawObject)
	newValue := reflect.ValueOf(newObject)

	copyAndStringifyRecursively(rawValue, newValue)
}

func copyAndStringifyRecursively(rawValue, newValue reflect.Value) {
	switch rawValue.Kind() {

	case reflect.Ptr:
		rawValueElem := rawValue.Elem()
		newValueElem := newValue.Elem()

		if !rawValueElem.IsValid() {
			return
		}

		copyAndStringifyRecursively(rawValueElem, newValueElem)

	case reflect.Interface:
		rawValueElem := rawValue.Elem()
		newValueElem := newValue.Elem()
		copyAndStringifyRecursively(rawValueElem, newValueElem)

	case reflect.Struct:
		for i := 0; i < rawValue.NumField(); i++ {
			copyAndStringifyRecursively(rawValue.Field(i), newValue.Field(i))
		}

	case reflect.Slice:
		newValue.Set(reflect.MakeSlice(newValue.Type(), rawValue.Len(), rawValue.Cap()))
		for i := 0; i < rawValue.Len(); i++ {
			copyAndStringifyRecursively(rawValue.Index(i), newValue.Index(i))
		}

	case reflect.Map:
		newValue.Set(reflect.MakeMap(newValue.Type()))
		for _, key := range rawValue.MapKeys() {
			rawValueField := rawValue.MapIndex(key)
			newValueField := newValue.MapIndex(key)
			copyAndStringifyRecursively(rawValueField, newValueField)
		}

	default:
		stringValue := fmt.Sprintf("%v", rawValue.Interface())
		newValue.SetString(stringValue)
	}
}
