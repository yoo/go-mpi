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
// "unsafe"
)

//Cart_coords
//Determines process coords in Cartesian topology given rank in group.
func Cart_coords(comm Comm, rank int, maxDims int) (int, int) {

	var cCoords C.int

	err := C.MPI_Cart_coords(C.MPI_Comm(comm),
		C.int(rank),
		C.int(maxDims),
		&cCoords)

	return int(cCoords), int(err)
}

//Cart_create
//Makes a new communicator to which Cartesian topology information has been attached.
func Cart_create(oldComm Comm, dims []int, periods []bool, reorder int) (Comm, int) {

	ndims := len(dims)
	cDims := make([]C.int, ndims)
	cPeriods := make([]C.int, ndims)

	for i := 0; i < ndims; i++ {
		cDims[i] = C.int(dims[i])
		if periods[i] {
			cPeriods[i] = C.int(1)
		} else {
			cPeriods[i] = C.int(0)
		}
	}

	var cCommCart C.MPI_Comm

	err := C.MPI_Cart_create(C.MPI_Comm(oldComm),
		C.int(ndims),
		&cDims[0],
		&cPeriods[0],
		C.int(reorder),
		&cCommCart)

	return Comm(cCommCart), int(err)
}

func Cart_get(comm Comm, maxDims int) ([]int, []bool, []int, int) {

	cDims := make([]C.int, maxDims)
	cPeriods := make([]C.int, maxDims)
	cCoords := make([]C.int, maxDims)

	arrayOfDims := make([]int, maxDims)
	arrayOfPeriods := make([]bool, maxDims)
	arrayOfCoords := make([]int, maxDims)

	err := C.MPI_Cart_get(C.MPI_Comm(comm),
		C.int(maxDims),
		&cDims[0],
		&cPeriods[0],
		&cCoords[0])

	for i := 0; i < maxDims; i++ {
		arrayOfDims[i] = int(cDims[i])
		if int(cPeriods[i]) == 0 {
			arrayOfPeriods[i] = false
		} else {
			arrayOfPeriods[i] = true
		}
		arrayOfCoords[i] = int(cCoords[i])
	}

	return arrayOfDims, arrayOfPeriods, arrayOfCoords, int(err)

}

//Cart_map
//Maps process to Cartesian topology information.
func Cart_map(comm Comm, dims []int, periods []bool) (int, int) {

	ndims := len(dims)
	cDims := make([]C.int, ndims)
	cPeriods := make([]C.int, ndims)
	for i := 0; i < ndims; i++ {
		cDims[i] = C.int(dims[i])
		if periods[i] {
			cPeriods[i] = C.int(1)
		} else {
			cPeriods[i] = C.int(0)
		}
	}
	var newRank C.int

	err := C.MPI_Cart_map(C.MPI_Comm(comm),
		C.int(ndims),
		&cDims[0],
		&cPeriods[0],
		&newRank)

	return int(newRank), int(err)

}

//Cart_rank
//Determines process rank in communicator given Cartesian location.
func Cart_rank(comm Comm, coords []int) (int, int) {

	ndims := len(coords)
	cCoords := make([]C.int, ndims)
	for i := 0; i < ndims; i++ {
		cCoords[i] = C.int(coords[i])
	}
	var cRank C.int

	err := C.MPI_Cart_rank(C.MPI_Comm(comm),
		&cCoords[0],
		&cRank)

	return int(cRank), int(err)
}

//Cart_shift
//Returns the shifted source and destination ranks, given a shift direction and amount.
func Cart_shift(comm Comm, direction int, displacement int) (int, int, int) {

	var rankSource C.int
	var rankDestination C.int

	err := C.MPI_Cart_shift(C.MPI_Comm(comm),
		C.int(direction),
		C.int(displacement),
		&rankSource,
		&rankDestination)

	return int(rankSource), int(rankDestination), int(err)
}

//Cart_sub
//Partitions a communicator into subgroups, which form lower-dimensional Cartesian subgrids.
func Cart_sub(comm Comm, remainDims []bool) (Comm, int) {

	count := len(remainDims)
	cRemainDims := make([]C.int, count)
	for i := 0; i < count; i++ {
		if remainDims[i] {
			cRemainDims[i] = C.int(1)
		} else {
			cRemainDims[i] = C.int(0)
		}
	}
	var newComm C.MPI_Comm

	err := C.MPI_Cart_sub(C.MPI_Comm(comm),
		&cRemainDims[0],
		&newComm)

	return Comm(newComm), int(err)

}

//Cartdim_get
//Retrieves Cartesian topology information associated with a communicator.
func Cartdim_get(comm Comm) (int, int) {

	var ndims C.int

	err := C.MPI_Cartdim_get(C.MPI_Comm(comm), &ndims)

	return int(ndims), int(err)

}

// next: int MPI_Cartdim_get(MPI_Comm comm, int *ndims);
