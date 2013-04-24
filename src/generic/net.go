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

//Close_port
//Close a port previously opened by MPI_Open_port.
func Close_port(portName string) int {

	cstrPortName := C.CString(portName)

	err := C.MPI_Close_port(cstrPortName)

	C.free(unsafe.Pointer(cstrPortName))

	return int(err)

}

//Open_port
//Establish an address that can be used to establish connections between groups of MPI processes.
func Open_port(info Info) (string, int) {

	var cstrPortName (*C.char)
	cstrPortName = (*C.char)(C.malloc(C.size_t(MAX_PORT_NAME)))

	err := C.MPI_Open_port(C.MPI_Info(info), cstrPortName)

	portName := C.GoString(cstrPortName)
	C.free(unsafe.Pointer(cstrPortName))

	return portName, int(err)
}
