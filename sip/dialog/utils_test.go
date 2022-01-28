package dialog

import (
	"reflect"
	"testing"
)

func TestGenerateDialogId(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"test 1 Generate DialogID - returns int32", true},
		{"test 2 Generate DialogID - returns int32", true},
		{"test 3 Generate DialogID - returns int32", true},
		{"test 4 Generate DialogID - returns int32", true},
		{"test 5 Generate DialogID - returns int32", true},
		{"test 6 Generate DialogID - returns int32", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateDialogId()
			if reflect.TypeOf(got).Kind() != reflect.Int32 {
				t.Errorf("GenerateDialogId() = %v, want %v", got, tt.want)
			}
		})
	}
}