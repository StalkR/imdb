package imdb

import (
	"fmt"
	"reflect"
	"testing"
)

// diffStruct tests if two structs are equal and if not which fields are different.
func diffStruct(t *testing.T, s string, a1 interface{}, a2 interface{}, ignoreFields ...string) {
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

FieldLoop:
	for i := 0; i < v1.NumField(); i++ {
		fieldName := t1.Field(i).Name
		for _, name := range ignoreFields {
			if fieldName == name {
				continue FieldLoop
			}
		}
		f1 := v1.Field(i).Interface()
		f2 := v2.Field(i).Interface()
		if !reflect.DeepEqual(f1, f2) {
			errPrefix := fmt.Sprintf("%s: %s.%s", s, t1.Name(), t1.Field(i).Name)
			if t1.Field(i).Type.Kind() == reflect.Slice {
				// Try to compare the slices as sets (ignoring
				// order).
				diffSliceAsSet(t, errPrefix, v1.Field(i), v2.Field(i))
			} else {
				t.Errorf("%s: %#v != %#v", errPrefix, f1, f2)
			}
		}
	}
}

func diffSliceAsSet(t *testing.T, errPrefix string, s1, s2 reflect.Value) {
	// Compare in O(len(s1) * len(s2)).

	foundS2Indices := make(map[int]struct{})
	s1Len := s1.Len()
	s2Len := s2.Len()
S1Loop:
	for i := 0; i < s1Len; i++ {
		e1 := s1.Index(i).Interface()
		for j := 0; j < s2Len; j++ {
			if _, ok := foundS2Indices[j]; ok {
				continue
			}
			e2 := s2.Index(j).Interface()
			if reflect.DeepEqual(e1, e2) {
				foundS2Indices[j] = struct{}{}
				continue S1Loop
			}
		}
		t.Errorf("%s: left elem %#v (index %d) missing from right side", errPrefix, e1, i)
	}
	for j := 0; j < s2Len; j++ {
		if _, ok := foundS2Indices[j]; !ok {
			e2 := s2.Index(j).Interface()
			t.Errorf("%s: right elem %#v (index %d) missing from left side", errPrefix, e2, j)
		}
	}
}
