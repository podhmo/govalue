package govalue_test

import (
	"testing"

	"github.com/podhmo/govalue"
)

func TestToCode(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "bool", args: args{v: true}, want: `true`},
		{name: "bool-false", args: args{v: false}, want: `false`},
		{name: "int", args: args{v: 10}, want: `10`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := govalue.ToCode(tt.args.v); got != tt.want {
				t.Errorf("ToCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
