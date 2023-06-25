package govalue_test

import (
	"testing"

	"github.com/podhmo/govalue"
)

type MyInt int

func TestToCode(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "nil", args: args{v: nil}, want: `nil`},
		{name: "bool", args: args{v: true}, want: `true`},
		{name: "bool-false", args: args{v: false}, want: `false`},
		{name: "int", args: args{v: 10}, want: `10`},
		{name: "int8", args: args{v: int8(10)}, want: `int8(10)`},
		{name: "uint", args: args{v: uint(10)}, want: `uint(10)`},
		{name: "uint16", args: args{v: uint16(10)}, want: `uint16(10)`},
		{name: "string", args: args{v: "foo"}, want: `"foo"`},
		{name: "string-with-quote", args: args{v: `"I'm lost"`}, want: `"\"I'm lost\""`},
		// advanced
		{name: "new-type", args: args{v: MyInt(10)}, want: `MyInt(10)`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := govalue.ToCode(tt.args.v); got != tt.want {
				t.Errorf("ToCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
