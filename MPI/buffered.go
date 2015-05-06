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
#include <stdlib.h>
#include <mpi.h>
*/
import "C"

import (
	"unsafe"
)

func Buffer_attach(buffer interface{}, size int) int {

	bufferVoidPointer := Get_void_ptr(buffer)

	err := C.MPI_Buffer_attach(bufferVoidPointer, C.int(size))

	return int(err)
}

func Buffer_detach() (uintptr, int, int) {

	var buffer uintptr
	var size C.int
	err := C.MPI_Buffer_detach(unsafe.Pointer(buffer), &size)

	return buffer, int(size), int(err)
}
