package imdb

import (
	"reflect"
	"testing"
)

// diffStruct tests if two structs are equal and if not which fields are different.
func diffStruct(t *testing.T, s string, a1 interface{}, a2 interface{}) {
	v1 := reflect.ValueOf(a1)
	v2 := reflect.ValueOf(a2)
	if v1.Kind() != reflect.Struct || v2.Kind() != reflect.Struct {
		t.Fatalf("%s: %v and %v must be structs", s, a1, a2)
	}
	t1 := v1.Type()
	t2 := v2.Type()
	if t1 != t2 {
		t.Fatalf("%s: %v and %v must be of the same type", s, a1, a2)
	}
	for i := 0; i < v1.NumField(); i++ {
		f1 := v1.Field(i).Interface()
		f2 := v2.Field(i).Interface()
		if !reflect.DeepEqual(f1, f2) {
			t.Errorf("%s: %s.%s: %#v != %#v", s, t1.Name(), t1.Field(i).Name, f1, f2)
		}
	}
}
