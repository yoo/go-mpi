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
#include "operations-helper.h"
*/
import "C"

var (
	// MPI operations for reduce operatios.
	MAX    Op = Op(C.HELPER_MPI_Get_op_max())
	MIN    Op = Op(C.HELPER_MPI_Get_op_min())
	SUM    Op = Op(C.HELPER_MPI_Get_op_sum())
	PROD   Op = Op(C.HELPER_MPI_Get_op_prod())
	LAND   Op = Op(C.HELPER_MPI_Get_op_land())
	BAND   Op = Op(C.HELPER_MPI_Get_op_band())
	LOR    Op = Op(C.HELPER_MPI_Get_op_lor())
	BOR    Op = Op(C.HELPER_MPI_Get_op_bor())
	LXOR   Op = Op(C.HELPER_MPI_Get_op_lxor())
	BXOR   Op = Op(C.HELPER_MPI_Get_op_bxor())
	MAXLOC Op = Op(C.HELPER_MPI_Get_op_maxloc())
	MINLOC Op = Op(C.HELPER_MPI_Get_op_minloc())
)
