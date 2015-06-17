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
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)

//Init
//Initializes the MPI execution environment
func Init(argv *[]string) int {

	argc := len(*argv)
	c_argc := C.int(argc)

	c_argv := make([]*C.char, argc)
	for index, value := range *argv {
		c_argv[index] = C.CString(value)
		defer C.free(unsafe.Pointer(c_argv[index]))
	}

	err := C.MPI_Init(&c_argc, (***C.char)(unsafe.Pointer(&c_argv)))
	goargs := make([]string, c_argc)
	for i := 0; i < int(c_argc); i++ {
		goargs[i] = C.GoString(c_argv[i])
	}
	*argv = goargs
	return int(err)
}

//Init
//Initialize MPI on different thread levels
func Init_thread(argv *[]string, required int) (int, int) {

	argc := len(*argv)
	c_argc := C.int(argc)

	c_argv := make([]*C.char, argc)
	for index, value := range *argv {
		c_argv[index] = C.CString(value)
		defer C.free(unsafe.Pointer(c_argv[index]))
	}

	var provided int
	err := C.MPI_Init_thread(&c_argc, (***C.char)(unsafe.Pointer(&c_argv)), C.int(required), (*C.int)(unsafe.Pointer(&provided)))
	goargs := make([]string, c_argc)
	for i := 0; i < int(c_argc); i++ {
		goargs[i] = C.GoString(c_argv[i])
	}
	*argv = goargs
	return provided, int(err)
}

//Initialized
//Indicates whether MPI_Init has been called.
func Initialized(flag *int) int {

	err := C.MPI_Initialized((*C.int)(unsafe.Pointer(flag)))

	return int(err)
}

//Finalize
//Terminates MPI execution environment.
func Finalize() int {

	err := C.MPI_Finalize()

	return int(err)
}

//Finalized
//Checks whether Finalize has completed.
func Finalized() (int, int) {

	var flag C.int

	err := C.MPI_Finalized(&flag)

	return int(flag), int(err)
}

//Abort
//Terminates MPI execution environment.
func Abort(comm Comm, errorcode int) int {

	err := C.MPI_Abort(C.MPI_Comm(comm), C.int(errorcode))

	return int(err)
}

//Cancel
//Cancels a communication request.
func Cancel(request Request) int {

	cRequest := C.MPI_Request(request)

	err := C.MPI_Cancel(&cRequest)

	return int(err)
}
