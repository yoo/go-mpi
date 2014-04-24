// Copyright (c) 2013 Alexander Beifuß <7beifuss@informatik.uni-hamburg.de>
// Johann Weging <johann@weging.net>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

package MPI

import (
	"fmt"
	"reflect"
	"unsafe"
)

func Get_void_ptr(iface interface{}) unsafe.Pointer {

	if iface == nil {
		return nil
	}

	val := reflect.ValueOf(iface)

	return unsafe.Pointer(val.Pointer())
}

func Type_create_goslice(iface interface{}) (Datatype, int) {
	typ := reflect.TypeOf(iface)

	if typ.Kind() != reflect.Slice || typ.Kind() != reflect.Array {
		errFmt := "Type_create_goslice expects a slice, was:%s\n"
		errMsg := fmt.Sprintf(errFmt, typ.Name())
		class, err := Add_error_class()
		if err != SUCCESS {
			return INT, err
		}
		code, err := Add_error_code(class)
		if err != SUCCESS {
			return INT, err
		}
		err = Add_error_string(code, errMsg)
		if err != SUCCESS {
			return INT, err
		}
	}

	elm := typ.Elem()
	var dTyp Datatype
	var err int
	switch typ.Kind() {
	case reflect.Array, reflect.Slice:
		dTyp, err = Type_create_goslice(elm)
		if err != SUCCESS {
			return INT, err
		}
	case reflect.Struct:
		dTyp, err = Type_create_gostruct(elm)
		if err != SUCCESS {
			return INT, err
		}

	case reflect.Ptr, reflect.Map, reflect.UnsafePointer, reflect.Chan, reflect.Uintptr:
		errFmt := "Type_create_goslice can only handle continuous memory, was: %s\n"
		errMsg := fmt.Sprintf(errFmt, elm.Name())
		class, err := Add_error_class()
		if err != SUCCESS {
			return INT, err
		}
		code, err := Add_error_code(class)
		if err != SUCCESS {
			return INT, err
		}
		err = Add_error_string(code, errMsg)
		if err != SUCCESS {
			return INT, err
		}
	default:
		size := elm.Size()
		dTyp, err = Type_contiguous(int(size), BYTE)
		if err != SUCCESS {
			return INT, err
		}
	}

	newType, err := Type_contiguous(typ.Len(), dTyp)
	if err != SUCCESS {
		return INT, err
	}

	return newType, SUCCESS
}

func Type_create_gostruct(iface interface{}) (Datatype, int) {

	typ := reflect.TypeOf(iface)

	// dereference pointer, if it is one
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		errFmt := "Type_create_gostruct expects a struct, was: %s\n"
		errMsg := fmt.Sprintf(errFmt, typ.Name())
		class, err := Add_error_class()
		if err != SUCCESS {
			return INT, err
		}
		code, err := Add_error_code(class)
		if err != SUCCESS {
			return INT, err
		}
		err = Add_error_string(code, errMsg)
		if err != SUCCESS {
			return INT, err
		}

	}

	count := typ.NumField()

	for i := 0; i < count; i++ {
		f := typ.Field(i)
		switch f.Type.Kind() {
		case reflect.Slice, reflect.Map, reflect.String, reflect.Chan, reflect.Ptr,
			reflect.UnsafePointer, reflect.Uintptr:
			errFmt := "Type_create_gostruct can't handle referenced data types, was: %s at filed %d\n"
			errMsg := fmt.Sprintf(errFmt, f.Type.Name(), i)
			class, err := Add_error_class()
			if err != SUCCESS {
				return INT, err
			}
			code, err := Add_error_code(class)
			if err != SUCCESS {
				return INT, err
			}
			err = Add_error_string(code, errMsg)
			if err != SUCCESS {
				return INT, err
			}
		}
		if typ.Kind() == reflect.Struct {
			return Type_contiguous(int(typ.Size()), BYTE)

		}
	}
	size := typ.Size()

	return Type_contiguous(int(size), BYTE)

}
