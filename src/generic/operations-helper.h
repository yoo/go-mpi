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

#ifndef __OPERATIONS_HELPER_H__
#define __OPERATIONS_HELPER_H__

#include <mpi.h>

MPI_Op HELPER_MPI_Get_op_max( void ){return MPI_MAX;}
MPI_Op HELPER_MPI_Get_op_min( void ){return MPI_MIN;}
MPI_Op HELPER_MPI_Get_op_sum( void ){return MPI_SUM;}
MPI_Op HELPER_MPI_Get_op_prod( void ){return MPI_PROD;}
MPI_Op HELPER_MPI_Get_op_land( void ){return MPI_LAND;}
MPI_Op HELPER_MPI_Get_op_band( void ){return MPI_BAND;}
MPI_Op HELPER_MPI_Get_op_lor( void ){return MPI_LOR;}
MPI_Op HELPER_MPI_Get_op_bor( void ){return MPI_BOR;}
MPI_Op HELPER_MPI_Get_op_lxor( void ){return MPI_LXOR;}
MPI_Op HELPER_MPI_Get_op_bxor( void ){return MPI_BXOR;}
MPI_Op HELPER_MPI_Get_op_maxloc( void ){return MPI_MAXLOC;}
MPI_Op HELPER_MPI_Get_op_minloc( void ){return MPI_MINLOC;}

#endif
