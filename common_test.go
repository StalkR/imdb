package imdb

import (
	"fmt"
	"reflect"
	"strings"
)

// diffStruct tests if two structs are equal and if not which fields are different.
func diffStruct(a1 interface{}, a2 interface{}) error {
	v1 := reflect.ValueOf(a1)
	v2 := reflect.ValueOf(a2)
	if v1.Kind() != reflect.Struct || v2.Kind() != reflect.Struct {
		return fmt.Errorf("%v and %v must be structs", a1, a2)
	}
	t1 := v1.Type()
	t2 := v2.Type()
	if t1 != t2 {
		return fmt.Errorf("%v and %v must be of the same type", a1, a2)
	}
	var errors []string
	for i := 0; i < v1.NumField(); i++ {
		f1 := v1.Field(i).Interface()
		f2 := v2.Field(i).Interface()
		if !reflect.DeepEqual(f1, f2) {
			errors = append(errors, fmt.Sprintf("%s.%s: %#v != %#v", t1.Name(), t1.Field(i).Name, f1, f2))
		}
	}
	if len(errors) == 0 {
		return nil
	}
	return fmt.Errorf("differences (%v): %v", len(errors), strings.Join(errors, "\n"))
}
