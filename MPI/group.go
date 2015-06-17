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

// import (
// 	"fmt"
// )

//Group_compare
//Compares two groups.
func Group_compare(group1, group2 Group) (int, int) {

	var result C.int
	err := C.MPI_Group_compare(C.MPI_Group(group1), C.MPI_Group(group2), &result)
	return int(result), int(err)
}

//Group_difference
//Makes a group from the difference of two groups.
func Group_difference(group1, group2 Group) (Group, int) {

	var newgroup C.MPI_Group
	err := C.MPI_Group_difference(C.MPI_Group(group1), C.MPI_Group(group2), &newgroup)
	return Group(newgroup), int(err)
}

//Group_excl
//Produces a group by reordering an existing group and taking only unlisted members.
func Group_excl(group Group, ranks []int) (Group, int) {

	length := len(ranks)

	var newgroup C.MPI_Group
	CintArray := make([]C.int, length)

	// copy slices data into the array
	for i := 0; i < length; i++ {
		CintArray[i] = C.int(ranks[i])
	}

	err := C.MPI_Group_excl(C.MPI_Group(group), C.int(length), &CintArray[0], &newgroup)

	return Group(newgroup), int(err)
}

//Group_free
//Frees a group.
func Group_free(group *Group) int {

	err := C.MPI_Group_free((*C.MPI_Group)(group))
	return int(err)
}

//Group_incl
//Produces a group by reordering an existing group and taking only listed members.
func Group_incl(group Group, ranks []int) (Group, int) {

	length := len(ranks)

	var newgroup C.MPI_Group
	CintArray := make([]C.int, length)

	// copy slices data into the array
	for i := 0; i < length; i++ {
		CintArray[i] = C.int(ranks[i])
	}

	err := C.MPI_Group_incl(C.MPI_Group(group), C.int(length), &CintArray[0], &newgroup)

	return Group(newgroup), int(err)
}

//Group_intersection
//Produces a group at the intersection of two existing groups.
func Group_intersection(group1, group2 Group) (Group, int) {

	var newgroup C.MPI_Group
	err := C.MPI_Group_intersection(C.MPI_Group(group1), C.MPI_Group(group2), &newgroup)

	return Group(newgroup), int(err)
}

//Group_range_excl
//Produces a group by excluding ranges of processes from an existing group.
func Group_range_excl(group Group, ranks [][3]int) (Group, int) {

	length := len(ranks)

	var newgroup C.MPI_Group
	CintArray := make([][3]C.int, length)

	// copy slices data into the array
	for i := 0; i < length; i++ {
		CintArray[i][0] = C.int(ranks[i][0])
		CintArray[i][1] = C.int(ranks[i][1])
		CintArray[i][2] = C.int(ranks[i][2])
	}

	err := C.MPI_Group_range_excl(C.MPI_Group(group), C.int(length), &(CintArray[0]), &newgroup)

	return Group(newgroup), int(err)
}

//Group_range_incl
//Creates a new group from ranges of ranks in an existing group.
func Group_range_incl(group Group, ranks [][3]int) (Group, int) {

	length := len(ranks)

	var newgroup C.MPI_Group
	CintArray := make([][3]C.int, length)

	// copy slices data into the array
	for i := 0; i < length; i++ {
		CintArray[i][0] = C.int(ranks[i][0])
		CintArray[i][1] = C.int(ranks[i][1])
		CintArray[i][2] = C.int(ranks[i][2])
	}

	err := C.MPI_Group_range_incl(C.MPI_Group(group), C.int(length), &(CintArray[0]), &newgroup)

	return Group(newgroup), int(err)
}

//Group_rank
//Returns the rank of the calling process in the given group.
func Group_rank(group Group) (int, int) {

	var rank C.int

	err := C.MPI_Group_rank(C.MPI_Group(group), &rank)

	return int(rank), int(err)
}

//Group_size
//Returns the size of a group.
func Group_size(group Group) (int, int) {

	var size C.int

	err := C.MPI_Group_size(C.MPI_Group(group), &size)

	return int(size), int(err)
}

//Group_translate_ranks
//Translates the ranks of processes in one group to those in another group.
func Group_translate_ranks(group1 Group, ranks1 []int, group2 Group) ([]int, int) {
	length := len(ranks1)
	CintArray1 := make([]C.int, length)
	CintArray2 := make([]C.int, length)
	ranks2 := make([]int, length)

	// copy slices data into the array
	for i := 0; i < length; i++ {
		CintArray1[i] = C.int(ranks1[i])
	}

	err := C.MPI_Group_translate_ranks(C.MPI_Group(group1), C.int(length), &CintArray1[0], C.MPI_Group(group2), &CintArray2[0])

	for i := 0; i < length; i++ {
		ranks2[i] = int(CintArray2[i])
	}

	return ranks2, int(err)
}

//Group_union
//Produces a group by combining two groups.
func Group_union(group1 Group, group2 Group) (Group, int) {

	var newgroup C.MPI_Group

	err := C.MPI_Group_union(C.MPI_Group(group1), C.MPI_Group(group2), &newgroup)

	return Group(newgroup), int(err)
}
