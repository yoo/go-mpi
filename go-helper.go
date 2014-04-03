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
	"strings"
	"unsafe"
)

func Get_void_ptr(iface interface{}) unsafe.Pointer {

	if iface == nil {
		return nil
	}

	val := reflect.ValueOf(iface)

	return unsafe.Pointer(val.Pointer())
}

//Type_create_slice
//Makes MPI.Type_create_struct accessible for go-struct;
//At the moment only flat structs (not embedded) with max. 1d arrays (no slices) can be used!
//Hoepfully, this will change.
func Type_create_gostruct(iface interface{}) (Datatype, int) {

	typ := reflect.TypeOf(iface)

	// dereference pointer, if it is one
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() == reflect.Struct {

		count := typ.NumField()
		length := make([]int, count, count)
		types := make([]Datatype, count, count)
		offsets := make([]Aint, count, count)

		for i := 0; i < count; i++ {
			f := typ.Field(i)

			if !f.Anonymous {

				tmp := f.Type
				typStr := tmp.String()

				strbuf := typStr
				arr := strings.Split(strbuf, "]")

				if len(arr) == 2 {
					typStr = arr[1]
					length[i] = int(tmp.Size())
				} else if len(arr) == 1 {
					length[i] = 1
				} else {
					fmt.Println("ERROR")
				}

				offsets[i] = Aint(uint(f.Offset))

				switch {
				case typStr == "uint8":
					types[i] = Datatype(BYTE)
				case typStr == "int8":
					types[i] = Datatype(SIGNED_CHAR)
				case typStr == "int16":
					types[i] = Datatype(SHORT)
				case typStr == "uint16":
					types[i] = Datatype(UNSIGNED_SHORT)
				case typStr == "int":
					types[i] = Datatype(INT)
				case typStr == "int32":
					types[i] = Datatype(INT)
				case typStr == "uint32":
					types[i] = Datatype(UNSIGNED)
				case typStr == "int64":
					types[i] = Datatype(LONG_LONG_INT)
				case typStr == "uint64":
					types[i] = Datatype(UNSIGNED_LONG_LONG_INT)
				case typStr == "float32":
					types[i] = Datatype(FLOAT)
				case typStr == "float64":
					types[i] = Datatype(DOUBLE)
				case typStr == "complex64":
					types[i] = Datatype(COMPLEX)
				case typStr == "complex128":
					types[i] = Datatype(DOUBLE_COMPLEX)

				default:
					fmt.Printf("Type %s is not supported (at the moment).\n", f.Type.String())
					return BYTE, ERR_ARG
				}

			} else {
				fmt.Println("Anonymous fields are not allowed")
			}
		}

		// DBG MSG
		// fmt.Println(count)
		// fmt.Println(length)
		// fmt.Println(types)
		// fmt.Println(offsets)

		// var barDatatype C.MPI_Datatype
		// C.MPI_Type_struct(count, length, offsets, types, &barDatatype)

		return Type_create_struct(length, offsets, types)
	}

	fmt.Println("Provide a struct!")
	return Datatype(INT), ERR_ARG
}
