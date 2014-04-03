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

#ifndef __CONST_HELPER_H__
#define __CONST_HELPER_H__

#include <mpi.h>

MPI_File HELPER_MPI_Get_file_null( void ){
  return MPI_FILE_NULL;
}

MPI_Win HELPER_MPI_Get_win_null( void ){
  return MPI_WIN_NULL;
}

MPI_Group HELPER_MPI_Get_gorup_null( void ){
    return MPI_GROUP_NULL;
}

MPI_Comm HELPER_MPI_Get_comm_self( void ){
    return MPI_COMM_SELF;
}

MPI_Errhandler HELPER_MPI_Get_errhandler_null( void ){
  return MPI_ERRHANDLER_NULL;
}

MPI_Request HELPER_MPI_Get_request_null( void ){
  return MPI_REQUEST_NULL;
}

MPI_Comm HELPER_MPI_Get_comm_null( void ){
    return MPI_COMM_NULL;
}

MPI_Info HELPER_MPI_Get_info_null( void ){
  return MPI_INFO_NULL;
}

MPI_Comm HELPER_MPI_Get_comm_world( void ){
    return MPI_COMM_WORLD;
}

MPI_Op HELPER_MPI_Get_op_null( void ){
  return MPI_OP_NULL;
}

MPI_Group HELPER_MPI_Get_group_empty( void ){
    return MPI_GROUP_EMPTY;
}

#endif
