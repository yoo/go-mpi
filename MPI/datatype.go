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
# include "datatype-helper.h"
*/
import "C"

import (
// "fmt"
// "reflect"
// "strings"
// "unsafe"
)

//experimental (for test cases onlu)
type GOMPI_INT C.int

type Aint C.MPI_Aint
type Comm C.MPI_Comm
type Datatype C.MPI_Datatype
type Errhandler C.MPI_Errhandler
type Group C.MPI_Group
type File C.MPI_File
type Info C.MPI_Info
type Offset C.MPI_Offset
type Op C.MPI_Op
type Request C.MPI_Request
type Status C.MPI_Status
type Win C.MPI_Win

// Datatypes
// Sice there is a lot of macro magic, we need to involve c helper functions.
// advantage: transparent to changes in openmpi.
// other possibilities: look into xxx.h and define it statically.
var (
	BYTE                   Datatype = Datatype(C.HELPER_MPI_Get_datatype_byte())
	CHAR                   Datatype = Datatype(C.HELPER_MPI_Get_datatype_char())
	CHARACTER              Datatype = Datatype(C.HELPER_MPI_Get_datatype_character())
	UNSIGNED_CHAR          Datatype = Datatype(C.HELPER_MPI_Get_datatype_uchar())
	SIGNED_CHAR            Datatype = Datatype(C.HELPER_MPI_Get_datatype_signed_char())
	SHORT                  Datatype = Datatype(C.HELPER_MPI_Get_datatype_short())
	UNSIGNED_SHORT         Datatype = Datatype(C.HELPER_MPI_Get_datatype_ushort())
	INT                    Datatype = Datatype(C.HELPER_MPI_Get_datatype_int())
	INTEGER                Datatype = Datatype(C.HELPER_MPI_Get_datatype_integer())
	UNSIGNED               Datatype = Datatype(C.HELPER_MPI_Get_datatype_uint())
	LONG                   Datatype = Datatype(C.HELPER_MPI_Get_datatype_long())
	UNSIGNED_LONG          Datatype = Datatype(C.HELPER_MPI_Get_datatype_ulong())
	LONG_LONG_INT          Datatype = Datatype(C.HELPER_MPI_Get_datatype_long_long_int())
	UNSIGNED_LONG_LONG_INT Datatype = Datatype(C.HELPER_MPI_Get_datatype_ulong_long())
	FLOAT                  Datatype = Datatype(C.HELPER_MPI_Get_datatype_float())
	REAL                   Datatype = Datatype(C.HELPER_MPI_Get_datatype_real())
	DOUBLE                 Datatype = Datatype(C.HELPER_MPI_Get_datatype_double())
	DOUBLE_PRECISION       Datatype = Datatype(C.HELPER_MPI_Get_datatype_double_precision())
	COMPLEX                Datatype = Datatype(C.HELPER_MPI_Get_datatype_complex())
	DOUBLE_COMPLEX         Datatype = Datatype(C.HELPER_MPI_Get_datatype_double_complex())
)
