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

const (
	SUCCESS                   = C.MPI_SUCCESS
	ERR_BUFFER                = C.MPI_ERR_BUFFER
	ERR_COUNT                 = C.MPI_ERR_COUNT
	ERR_TYPE                  = C.MPI_ERR_TYPE
	ERR_TAG                   = C.MPI_ERR_TAG
	ERR_COMM                  = C.MPI_ERR_COMM
	ERR_RANK                  = C.MPI_ERR_RANK
	ERR_REQUEST               = C.MPI_ERR_REQUEST
	ERR_ROOT                  = C.MPI_ERR_ROOT
	ERR_GROUP                 = C.MPI_ERR_GROUP
	ERR_OP                    = C.MPI_ERR_OP
	ERR_TOPOLOGY              = C.MPI_ERR_TOPOLOGY
	ERR_DIMS                  = C.MPI_ERR_DIMS
	ERR_ARG                   = C.MPI_ERR_ARG
	ERR_UNKNOWN               = C.MPI_ERR_UNKNOWN
	ERR_TRUNCATE              = C.MPI_ERR_TRUNCATE
	ERR_OTHER                 = C.MPI_ERR_OTHER
	ERR_INTERN                = C.MPI_ERR_INTERN
	ERR_IN_STATUS             = C.MPI_ERR_IN_STATUS
	ERR_PENDING               = C.MPI_ERR_PENDING
	ERR_ACCESS                = C.MPI_ERR_ACCESS
	ERR_AMODE                 = C.MPI_ERR_AMODE
	ERR_ASSERT                = C.MPI_ERR_ASSERT
	ERR_BAD_FILE              = C.MPI_ERR_BAD_FILE
	ERR_BASE                  = C.MPI_ERR_BASE
	ERR_CONVERSION            = C.MPI_ERR_CONVERSION
	ERR_DISP                  = C.MPI_ERR_DISP
	ERR_DUP_DATAREP           = C.MPI_ERR_DUP_DATAREP
	ERR_FILE_EXISTS           = C.MPI_ERR_FILE_EXISTS
	ERR_FILE_IN_USE           = C.MPI_ERR_FILE_IN_USE
	ERR_FILE                  = C.MPI_ERR_FILE
	ERR_INFO_KEY              = C.MPI_ERR_INFO_KEY
	ERR_INFO_NOKEY            = C.MPI_ERR_INFO_NOKEY
	ERR_INFO_VALUE            = C.MPI_ERR_INFO_VALUE
	ERR_INFO                  = C.MPI_ERR_INFO
	ERR_IO                    = C.MPI_ERR_IO
	ERR_KEYVAL                = C.MPI_ERR_KEYVAL
	ERR_LOCKTYPE              = C.MPI_ERR_LOCKTYPE
	ERR_NAME                  = C.MPI_ERR_NAME
	ERR_NO_MEM                = C.MPI_ERR_NO_MEM
	ERR_NOT_SAME              = C.MPI_ERR_NOT_SAME
	ERR_NO_SPACE              = C.MPI_ERR_NO_SPACE
	ERR_NO_SUCH_FILE          = C.MPI_ERR_NO_SUCH_FILE
	ERR_PORT                  = C.MPI_ERR_PORT
	ERR_QUOTA                 = C.MPI_ERR_QUOTA
	ERR_READ_ONLY             = C.MPI_ERR_READ_ONLY
	ERR_RMA_CONFLICT          = C.MPI_ERR_RMA_CONFLICT
	ERR_RMA_SYNC              = C.MPI_ERR_RMA_SYNC
	ERR_SERVICE               = C.MPI_ERR_SERVICE
	ERR_SIZE                  = C.MPI_ERR_SIZE
	ERR_SPAWN                 = C.MPI_ERR_SPAWN
	ERR_UNSUPPORTED_DATAREP   = C.MPI_ERR_UNSUPPORTED_DATAREP
	ERR_UNSUPPORTED_OPERATION = C.MPI_ERR_UNSUPPORTED_OPERATION
	ERR_WIN                   = C.MPI_ERR_WIN
	ERR_LASTCODE              = C.MPI_ERR_LASTCODE
)

//Error_string
//Returns a string for a given error code.
func Error_string(errorcode int) (string, int) {

	// temporary workaround
	var BUFSIZE_IN_BYTE C.size_t = 256

	var resultlen C.int
	var str unsafe.Pointer
	str = C.malloc(BUFSIZE_IN_BYTE)

	err := C.MPI_Error_string(C.int(errorcode), (*C.char)(str), &resultlen)
	resstr := C.GoStringN((*C.char)(str), resultlen)
	C.free(str)

	return resstr, int(err)
}
