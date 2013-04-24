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

#ifndef __DATATYPE_HELPER_H__
#define __DATATYPE_HELPER_H__

#include <mpi.h>

MPI_Datatype HELPER_MPI_Get_datatype_byte( void ){return MPI_BYTE;}
MPI_Datatype HELPER_MPI_Get_datatype_char( void ){return MPI_CHAR;}
MPI_Datatype HELPER_MPI_Get_datatype_character( void ){return MPI_CHARACTER;}
MPI_Datatype HELPER_MPI_Get_datatype_uchar( void ){return MPI_UNSIGNED_CHAR;}
MPI_Datatype HELPER_MPI_Get_datatype_signed_char( void ){return MPI_SIGNED_CHAR;}
MPI_Datatype HELPER_MPI_Get_datatype_short( void ){return MPI_SHORT;}
MPI_Datatype HELPER_MPI_Get_datatype_ushort( void ){return MPI_UNSIGNED_SHORT;}
MPI_Datatype HELPER_MPI_Get_datatype_int( void ){return MPI_INT;}
MPI_Datatype HELPER_MPI_Get_datatype_integer( void ){return MPI_INTEGER;}
MPI_Datatype HELPER_MPI_Get_datatype_uint( void ){return MPI_UNSIGNED;}
MPI_Datatype HELPER_MPI_Get_datatype_long( void ){return MPI_LONG;}
MPI_Datatype HELPER_MPI_Get_datatype_ulong( void ){return MPI_UNSIGNED_LONG;}
MPI_Datatype HELPER_MPI_Get_datatype_long_long_int( void ){return MPI_LONG_LONG_INT;}
MPI_Datatype HELPER_MPI_Get_datatype_ulong_long( void ){return MPI_UNSIGNED_LONG_LONG;}
MPI_Datatype HELPER_MPI_Get_datatype_float( void ){return MPI_FLOAT;}
MPI_Datatype HELPER_MPI_Get_datatype_real( void ){return MPI_REAL;}
MPI_Datatype HELPER_MPI_Get_datatype_double( void ){return MPI_DOUBLE;}
MPI_Datatype HELPER_MPI_Get_datatype_double_precision( void ){return MPI_DOUBLE_PRECISION;}
MPI_Datatype HELPER_MPI_Get_datatype_complex( void ){return MPI_COMPLEX;}
MPI_Datatype HELPER_MPI_Get_datatype_double_complex( void ){return MPI_DOUBLE_COMPLEX;}

#endif
