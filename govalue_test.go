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
		{name: "float32", args: args{v: float32(3.14)}, want: `float32(3.140000)`},
		{name: "float64", args: args{v: 3.14}, want: `3.140000`},
		// composite
		{name: "[]string-0", args: args{v: []string{}}, want: `[]string{}`},
		{name: "[]string-1", args: args{v: []string{"foo"}}, want: `[]string{"foo"}`},
		{name: "[]string-2", args: args{v: []string{"foo", "bar"}}, want: `[]string{"foo", "bar"}`},
		{name: "[]int3", args: args{v: []int{1, 2, 3}}, want: `[]int{1, 2, 3}`},
		// advanced
		{name: "new-type", args: args{v: MyInt(10)}, want: `MyInt(10)`},
		{name: "new-type[]", args: args{v: []MyInt{1, 2, 3}}, want: `[]MyInt{1, 2, 3}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := govalue.ToCode(tt.args.v); got != tt.want {
				t.Errorf("ToCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
