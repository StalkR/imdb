package imdb

import (
	"fmt"
	"reflect"
	"strings"
)

// diffStruct tests if two structs are equal and if not which fields are different.
func diffStruct(got interface{}, want interface{}) error {
	gotv := reflect.ValueOf(got)
	wantv := reflect.ValueOf(want)
	if gotv.Kind() != reflect.Struct || wantv.Kind() != reflect.Struct {
		return fmt.Errorf("%v and %v must be structs", got, want)
	}
	tgot := gotv.Type()
	twant := wantv.Type()
	if tgot != twant {
		return fmt.Errorf("%v and %v must be of the same type", got, want)
	}
	var errors []string
	if gotv.NumField() != wantv.NumField() {
		errors = append(errors, fmt.Sprintf("different lengths: got %v, want %v", gotv.NumField(), wantv.NumField()))
	}
	min := gotv.NumField()
	if wantv.NumField() < min {
		min = wantv.NumField()
	}
	for i := 0; i < min; i++ {
		gotf := gotv.Field(i).Interface()
		wantf := wantv.Field(i).Interface()
		if !reflect.DeepEqual(gotf, wantf) {
			errors = append(errors, fmt.Sprintf("%s.%s:\n\t got: %#v\n\twant: %#v", tgot.Name(), tgot.Field(i).Name, gotf, wantf))
		}
	}
	if len(errors) == 0 {
		return nil
	}
	return fmt.Errorf("differences (%v):\n%v", len(errors), strings.Join(errors, "\n"))
}
