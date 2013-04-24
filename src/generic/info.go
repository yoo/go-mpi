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

/*
#include <mpi.h>

void* foo(){
	return ((void*)&(ompi_mpi_info_null));
};
*/
import "C"

import (
	"unsafe"
)

const (
	MAX_INFO_KEY = C.MPI_MAX_INFO_KEY
	MAX_INFO_VAL = C.MPI_MAX_INFO_VAL
)

func Info_create() (Info, int) {
	var info Info
	err := int(C.MPI_Info_create((*C.MPI_Info)(&info)))
	return info, err
}

func Info_delete(info Info, key string) int {
	return int(C.MPI_Info_delete(C.MPI_Info(info), C.CString(key)))
}

func Info_dup(info Info) (Info, int) {
	var newInfo Info
	err := int(C.MPI_Info_dup(C.MPI_Info(info), (*C.MPI_Info)(&newInfo)))
	return newInfo, err
}

func Info_free(info *Info) int {
	return int(C.MPI_Info_free((*C.MPI_Info)(info)))
}

func Info_get(info Info, key string) (string, int, int) {
	var flag C.int
	value := make([]byte, MAX_INFO_VAL+1)
	cValue := (*C.char)(unsafe.Pointer(&value[0]))
	err := int(C.MPI_Info_get(C.MPI_Info(info), C.CString(key),
		C.int(MAX_INFO_VAL+1), cValue, (*C.int)(&flag)))
	return string(value), int(flag), err
}

func Info_get_nkeys(info Info) (int, int) {
	var nkeys C.int
	err := int(C.MPI_Info_get_nkeys(C.MPI_Info(info), (*C.int)(&nkeys)))
	return int(nkeys), err
}

func Info_get_nthkey(info Info, n int) (string, int) {
	key := make([]byte, MAX_INFO_KEY+1)
	cKey := (*C.char)(unsafe.Pointer(&key[0]))
	err := int(C.MPI_Info_get_nthkey(C.MPI_Info(info), C.int(n), cKey))
	return string(key), err
}

func Info_get_valuelen(info Info, key string) (int, int, int) {
	var flag C.int
	var valuelen C.int
	err := int(C.MPI_Info_get_valuelen(C.MPI_Info(info), C.CString(key),
		&valuelen, &flag))
	return int(valuelen), int(flag), err
}

func Info_set(info Info, key string, value string) int {
	return int(C.MPI_Info_set(C.MPI_Info(info), C.CString(key), C.CString(value)))
}
