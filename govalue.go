package govalue

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"sync"
)

var pool = &sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func ToCode(v any) string {
	if v == nil {
		return "nil"
	}

	buf := pool.Get().(*bytes.Buffer)
	buf.Reset()
	defer pool.Put(buf)
	if err := writeCode(buf, reflect.TypeOf(v), reflect.ValueOf(v), false); err != nil {
		return fmt.Sprintf("<%s>", err.Error())
	}
	return buf.String()
}

func writeCode(buf *bytes.Buffer, rt reflect.Type, rv reflect.Value, inslice bool) error {
	switch rt.Kind() {
	case reflect.Invalid:
		if _, err := buf.WriteString("<invalid>"); err != nil {
			return err
		}
	case reflect.Bool:
		if rt == rbool || inslice {
			if _, err := fmt.Fprintf(buf, "%t", rv.Bool()); err != nil {
				return err
			}
		} else {
			if _, err := fmt.Fprintf(buf, "%s(%t)", rt.Name(), rv.Bool()); err != nil {
				return err
			}
		}
		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if rt == rint || inslice {
			if _, err := fmt.Fprintf(buf, "%d", rv.Int()); err != nil {
				return err
			}
		} else {
			if _, err := fmt.Fprintf(buf, "%s(%d)", rt.Name(), rv.Int()); err != nil {
				return err
			}
		}
		return nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if inslice {
			if _, err := fmt.Fprintf(buf, "%d", rv.Int()); err != nil {
				return err
			}
		} else {
			if _, err := fmt.Fprintf(buf, "%s(%d)", rt.Name(), rv.Uint()); err != nil {
				return err
			}
		}
		return nil
	case reflect.Uintptr:
	case reflect.Float32, reflect.Float64:
		if rt == rfloat64 || inslice {
			if _, err := fmt.Fprintf(buf, "%f", rv.Float()); err != nil {
				return err
			}
		} else {
			if _, err := fmt.Fprintf(buf, "%s(%f)", rt.Name(), rv.Float()); err != nil {
				return err
			}
		}
		return nil
	case reflect.Complex64:
	case reflect.Complex128:
	case reflect.Chan:
	case reflect.Func:
	case reflect.Interface:
	case reflect.Map:
	case reflect.Pointer:
	case reflect.Slice, reflect.Array:
		if err := writeType(buf, rt); err != nil {
			return err
		}
		buf.WriteString("{")
		st := rt.Elem()
		inslice := true
		for i, n := 0, rv.Len(); i < n; i++ {
			sv := rv.Index(i)
			if err := writeCode(buf, st, sv, inslice); err != nil {
				return err
			}
			if i < n-1 {
				buf.WriteString(", ")
			}
		}
		buf.WriteString("}")
		return nil
	case reflect.String:
		if rt == rstring {
			if _, err := buf.WriteString(strconv.Quote(rv.String())); err != nil {
				return err
			}
		} else {
			if _, err := fmt.Fprintf(buf, "%s(%q)", rt.Name(), rv.String()); err != nil {
				return err
			}
		}
		return nil
	case reflect.Struct:
	case reflect.UnsafePointer:
	default:
		panic(fmt.Sprintf("unexpected kind: %s, %v", rt.Kind(), rv)) // we need panic instead of error
	}
	return fmt.Errorf("not implemented")
}

var (
	rbool    = reflect.TypeOf(bool(false))
	rint     = reflect.TypeOf(int(0))
	rfloat64 = reflect.TypeOf(float64(0.0))
	rstring  = reflect.TypeOf("")
)

func writeType(buf *bytes.Buffer, rt reflect.Type) error {
	switch rt.Kind() {
	case reflect.Slice:
		buf.WriteString("[]")
		return writeType(buf, rt.Elem())
	default:
		buf.WriteString(rt.Name()) // now, supporting basic type only
		return nil
	}
}
