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
	if err := writeCode(buf, v); err != nil {
		return fmt.Sprintf("<%s>", err.Error())
	}
	return buf.String()
}

func writeCode(buf *bytes.Buffer, v any) error {
	rt := reflect.TypeOf(v)
	switch rt.Kind() {
	case reflect.Invalid:
		if _, err := buf.WriteString("<invalid>"); err != nil {
			return err
		}
	case reflect.Bool:
		if v, ok := v.(bool); ok && v {
			if _, err := buf.WriteString("true"); err != nil {
				return err
			}
		} else if _, err := buf.WriteString("false"); err != nil {
			return err
		}
		return nil
	case reflect.Int:
		if rt.Name() == "int" {
			if _, err := buf.WriteString(strconv.Itoa(v.(int))); err != nil {
				return err
			}
		} else {
			if _, err := fmt.Fprintf(buf, "%s(%d)", rt.Name(), v); err != nil {
				return err
			}
		}
		return nil
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if _, err := fmt.Fprintf(buf, "%s(%d)", rt.Name(), v); err != nil {
			return err
		}
		return nil
	case reflect.Uintptr:
	case reflect.Float32:
	case reflect.Float64:
	case reflect.Complex64:
	case reflect.Complex128:
	case reflect.Array:
	case reflect.Chan:
	case reflect.Func:
	case reflect.Interface:
	case reflect.Map:
	case reflect.Pointer:
	case reflect.Slice:
	case reflect.String:
		if _, err := buf.WriteString(strconv.Quote(v.(string))); err != nil {
			return err
		}
		return nil
	case reflect.Struct:
	case reflect.UnsafePointer:
	default:
		rv := reflect.ValueOf(v)
		panic(fmt.Sprintf("unexpected kind: %s, %v", rt.Kind(), rv)) // we need panic instead of error
	}
	return fmt.Errorf("not implemented")
}
