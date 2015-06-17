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
#include "const-helper.h"
*/
import "C"

/*
 * Miscellaneous constants
 */
const (
	ANY_SOURCE         = C.MPI_ANY_SOURCE         /* match any source rank */
	PROC_NULL          = C.MPI_PROC_NULL          /* rank of null process */
	ROOT               = C.MPI_ROOT               /* */
	ANY_TAG            = C.MPI_ANY_TAG            /* match any message tag */
	MAX_PROCESSOR_NAME = C.MPI_MAX_PROCESSOR_NAME /* max proc. name length */
	MAX_ERROR_STRING   = C.MPI_MAX_ERROR_STRING   /* max error message length */
	MAX_OBJECT_NAME    = C.MPI_MAX_OBJECT_NAME    /* max object name length */
	UNDEFINED          = C.MPI_UNDEFINED          /* undefined stuff */
	CART               = C.MPI_CART               /* cartesian topology */
	GRAPH              = C.MPI_GRAPH              /* graph topology */
	KEYVAL_INVALID     = C.MPI_KEYVAL_INVALID     /* invalid key value */
)

/*
 * More constants
 */
const (
	//BOTTOM = C.MPI_BOTTOM /* base reference address */
	//IN_PLACE             = C.MPI_IN_PLACE             /* in place buffer */
	BSEND_OVERHEAD = C.MPI_BSEND_OVERHEAD /* size of bsend header + ptr */
	//ARGV_NULL      = C.MPI_ARGV_NULL      /* NULL argument vector */
	//ARGVS_NULL     = C.MPI_ARGVS_NULL     /* NULL argument vectors */
	//ERRCODES_IGNORE      = C.MPI_ERRCODES_IGNORE      /* don't return error codes */
	MAX_PORT_NAME        = C.MPI_MAX_PORT_NAME        /* max port name length */
	ORDER_C              = C.MPI_ORDER_C              /* C row major order */
	ORDER_FORTRAN        = C.MPI_ORDER_FORTRAN        /* Fortran column major order */
	DISTRIBUTE_BLOCK     = C.MPI_DISTRIBUTE_BLOCK     /* block distribution */
	DISTRIBUTE_CYCLIC    = C.MPI_DISTRIBUTE_CYCLIC    /* cyclic distribution */
	DISTRIBUTE_NONE      = C.MPI_DISTRIBUTE_NONE      /* not distributed */
	DISTRIBUTE_DFLT_DARG = C.MPI_DISTRIBUTE_DFLT_DARG /* default distribution arg */

)

/*
 * Since these values are arbitrary to Open MPI, we might as well make
 * them the same as ROMIO for ease of mapping.  These values taken
 * from ROMIO's mpio.h file.
 */
const (
	MODE_CREATE          = C.MPI_MODE_CREATE          /* ADIO_CREATE */
	MODE_RDONLY          = C.MPI_MODE_RDONLY          /* ADIO_RDONLY */
	MODE_WRONLY          = C.MPI_MODE_WRONLY          /* ADIO_WRONLY  */
	MODE_RDWR            = C.MPI_MODE_RDWR            /* ADIO_RDWR  */
	MODE_DELETE_ON_CLOSE = C.MPI_MODE_DELETE_ON_CLOSE /* ADIO_DELETE_ON_CLOSE */
	MODE_UNIQUE_OPEN     = C.MPI_MODE_UNIQUE_OPEN     /* ADIO_UNIQUE_OPEN */
	MODE_EXCL            = C.MPI_MODE_EXCL            /* ADIO_EXCL */
	MODE_APPEND          = C.MPI_MODE_APPEND          /* ADIO_APPEND */
	MODE_SEQUENTIAL      = C.MPI_MODE_SEQUENTIAL      /* ADIO_SEQUENTIAL */

	DISPLACEMENT_CURRENT = C.MPI_DISPLACEMENT_CURRENT

	SEEK_SET = C.MPI_SEEK_SET
	SEEK_CUR = C.MPI_SEEK_CUR
	SEEK_END = C.MPI_SEEK_END
)

/*
 * MPI-2 One-Sided Communications asserts
 */
const (
	MODE_NOCHECK   = C.MPI_MODE_NOCHECK
	MODE_NOPRECEDE = C.MPI_MODE_NOPRECEDE
	MODE_NOPUT     = C.MPI_MODE_NOPUT
	MODE_NOSTORE   = C.MPI_MODE_NOSTORE
	MODE_NOSUCCEED = C.MPI_MODE_NOSUCCEED

	LOCK_EXCLUSIVE = C.MPI_LOCK_EXCLUSIVE
	LOCK_SHARED    = C.MPI_LOCK_SHARED
)

// Comparison results (see mpi.h and mpif.h.in)
const (
	IDENT     = C.MPI_IDENT
	CONGRUENT = C.MPI_CONGRUENT
	SIMILAR   = C.MPI_SIMILAR
	UNEQUAL   = C.MPI_UNEQUAL
)

// MPI_Init_thread constants (see mpi.h and mpif.h.in)
const (
	THREAD_SINGLE     = C.MPI_THREAD_SINGLE
	THREAD_FUNNELED   = C.MPI_THREAD_FUNNELED
	THREAD_SERIALIZED = C.MPI_THREAD_SERIALIZED
	THREAD_MULTIPLE   = C.MPI_THREAD_MULTIPLE
)

// Datatype combiners (see mpi.h and mpif.h.in)
const (
	COMBINER_NAMED            = C.MPI_COMBINER_NAMED
	COMBINER_DUP              = C.MPI_COMBINER_DUP
	COMBINER_CONTIGUOUS       = C.MPI_COMBINER_CONTIGUOUS
	COMBINER_VECTOR           = C.MPI_COMBINER_VECTOR
	COMBINER_HVECTOR_INTEGER  = C.MPI_COMBINER_HVECTOR_INTEGER
	COMBINER_HVECTOR          = C.MPI_COMBINER_HVECTOR
	COMBINER_INDEXED          = C.MPI_COMBINER_INDEXED
	COMBINER_HINDEXED_INTEGER = C.MPI_COMBINER_HINDEXED_INTEGER
	COMBINER_HINDEXED         = C.MPI_COMBINER_HINDEXED
	COMBINER_INDEXED_BLOCK    = C.MPI_COMBINER_INDEXED_BLOCK
	COMBINER_STRUCT_INTEGER   = C.MPI_COMBINER_STRUCT_INTEGER
	COMBINER_STRUCT           = C.MPI_COMBINER_STRUCT
	COMBINER_SUBARRAY         = C.MPI_COMBINER_SUBARRAY
	COMBINER_DARRAY           = C.MPI_COMBINER_DARRAY
	COMBINER_F90_REAL         = C.MPI_COMBINER_F90_REAL
	COMBINER_F90_COMPLEX      = C.MPI_COMBINER_F90_COMPLEX
	COMBINER_F90_INTEGER      = C.MPI_COMBINER_F90_INTEGER
	COMBINER_RESIZED          = C.MPI_COMBINER_RESIZED
)

// Communicators
var (
	COMM_WORLD Comm = Comm(C.HELPER_MPI_Get_comm_world())
	COMM_SELF  Comm = Comm(C.HELPER_MPI_Get_comm_self())
	// Groups
	GROUP_EMPTY Group = Group(C.HELPER_MPI_Get_group_empty())
	//Null handles
	GROUP_NULL       Group      = Group(C.HELPER_MPI_Get_gorup_null())
	COMM_NULL        Comm       = Comm(C.HELPER_MPI_Get_comm_null())
	INFO_NULL        Info       = Info(C.HELPER_MPI_Get_info_null())
	REQUEST_NULL     Request    = Request(C.HELPER_MPI_Get_request_null())
	OP_NULL          Op         = Op(C.HELPER_MPI_Get_op_null())
	ERRHANDLER_NULL  Errhandler = Errhandler(C.HELPER_MPI_Get_errhandler_null())
	ERRORS_RETURN    Errhandler = Errhandler(C.HELPER_MPI_Get_errors_return())
	ERRORS_ARE_FATAL Errhandler = Errhandler(C.HELPER_MPI_Get_errors_are_fatal())
	WIN_NULL         Win        = Win(C.HELPER_MPI_Get_win_null())
	FILE_NULL        File       = File(C.HELPER_MPI_Get_file_null())
)
