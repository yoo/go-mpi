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
#include<mpi.h>
*/
import "C"

import (
	"fmt"
	"os"
	"unsafe"
)

//Type_commit
//Commits a data type.
func Type_commit(datatype *Datatype) int {

	err := C.MPI_Type_commit((*C.MPI_Datatype)(unsafe.Pointer(datatype)))

	return int(err)
}

//Type_contiguous
//Creates a contiguous datatype.
func Type_contiguous(count int, oldtype Datatype) (Datatype, int) {

	var newtype C.MPI_Datatype

	err := C.MPI_Type_contiguous(C.int(count),
		C.MPI_Datatype(oldtype),
		(*C.MPI_Datatype)(unsafe.Pointer(&newtype)))

	return Datatype(newtype), int(err)
}

//Type_create_darray
//Create a datatype representing a distributed array
func Type_create_darray(size int, rank int, ndims int,
	arrayOfGSizes []int,
	arrayOfDistribs []int,
	arrayOfDArgs []int,
	arrayOfPSizes []int,
	order int,
	oldType Datatype) (Datatype, int) {

	cArrayOfGSizes := make([]C.int, len(arrayOfGSizes))
	cArrayOfDistribs := make([]C.int, len(arrayOfDistribs))
	cArrayOfDArgs := make([]C.int, len(arrayOfDArgs))
	cArrayOfPSizes := make([]C.int, len(arrayOfPSizes))

	for i := 0; i < len(arrayOfGSizes); i++ {
		cArrayOfGSizes[i] = C.int(arrayOfGSizes[i])
	}

	for i := 0; i < len(arrayOfDistribs); i++ {
		cArrayOfDistribs[i] = C.int(arrayOfDistribs[i])
	}

	for i := 0; i < len(arrayOfDArgs); i++ {
		cArrayOfDArgs[i] = C.int(arrayOfDArgs[i])
	}

	for i := 0; i < len(arrayOfPSizes); i++ {
		cArrayOfPSizes[i] = C.int(arrayOfPSizes[i])
	}

	var newtype C.MPI_Datatype

	err := C.MPI_Type_create_darray(C.int(size), C.int(rank), C.int(ndims),
		&cArrayOfGSizes[0],
		&cArrayOfDistribs[0],
		&cArrayOfDArgs[0],
		&cArrayOfPSizes[0],
		C.int(order),
		C.MPI_Datatype(oldType),
		&newtype)

	return Datatype(newtype), int(err)

}

//Type_hvector
//type_hvector.
func Type_hvector(count int, blockLength int, stride Aint, oldType Datatype) (Datatype, int) {

	var cNewType C.MPI_Datatype

	err := C.MPI_Type_hvector(C.int(count),
		C.int(blockLength),
		C.MPI_Aint(stride),
		C.MPI_Datatype(oldType),
		&cNewType)

	return Datatype(cNewType), int(err)
}

//Type_create_hindexed
//Create a datatype for an indexed datatype with displacements in bytes.
func Type_create_hindexed(blockLengths []int, displacements []Aint, oldType Datatype) (Datatype, int) {

	if len(blockLengths) != len(displacements) {
		fmt.Println("[ERROR] In Type_get_indexed: Length of blocklLengths differs from length of displacements.")
		os.Exit(1)
	}

	count := len(blockLengths)

	cArrayOfBlockLengths := make([]C.int, count)
	cArrayOfDisplacements := make([]C.MPI_Aint, count)

	for i := 0; i < count; i++ {
		cArrayOfBlockLengths[i] = C.int(blockLengths[i])
		cArrayOfDisplacements[i] = C.MPI_Aint(displacements[i])
	}

	var cNewType C.MPI_Datatype

	err := C.MPI_Type_create_hindexed(C.int(count),
		&cArrayOfBlockLengths[0],
		&cArrayOfDisplacements[0],
		C.MPI_Datatype(oldType),
		&cNewType)

	return Datatype(cNewType), int(err)
}

// int MPI_Type_create_keyval(MPI_Type_copy_attr_function *type_copy_attr_fn, MPI_Type_delete_attr_function *type_delete_attr_fn, int *type_keyval, void *extra_state);

//Type_create_indexed_block
//Create an indexed datatype with constant-sized blocks.
func Type_create_indexed_block(blockLength int, displacements []int, oldType Datatype) (Datatype, int) {

	count := len(displacements)
	cArrayOfDisplacements := make([]C.int, count)

	for i := 0; i < count; i++ {
		cArrayOfDisplacements[i] = C.int(displacements[i])
	}

	var cNewType C.MPI_Datatype

	err := C.MPI_Type_create_indexed_block(C.int(count),
		C.int(blockLength),
		&cArrayOfDisplacements[0],
		C.MPI_Datatype(oldType),
		&cNewType)

	return Datatype(cNewType), int(err)
}

//Type_create_struct
//Creates a structured data type.
func Type_create_struct(array_of_blocklengths []int,
	array_of_displacements []Aint,
	array_of_types []Datatype) (Datatype, int) {

	var newtype C.MPI_Datatype
	count := len(array_of_types)
	if count == len(array_of_displacements) && count == len(array_of_blocklengths) {

		err := C.MPI_Type_create_struct(C.int(count),
			(*C.int)(unsafe.Pointer(&array_of_blocklengths[0])),
			(*C.MPI_Aint)(unsafe.Pointer(&array_of_displacements[0])),
			(*C.MPI_Datatype)(unsafe.Pointer(&array_of_types[0])),
			(*C.MPI_Datatype)(unsafe.Pointer(&newtype)))

		return Datatype(newtype), int(err)

	}
	return Datatype(newtype), ERR_DIMS
}

//Type_create_subarray
//Create a datatype for a subarray of a regular, multidimensional array.
func Type_create_subarray(ndims int, arrayOfSizes []int, arrayOfSubsizes []int, arrayOfStarts []int,
	order int, oldType Datatype) (Datatype, int) {

	cArrayOfSizes := make([]C.int, len(arrayOfSizes))
	cArrayOfSubsizes := make([]C.int, len(arrayOfSubsizes))
	cArrayOfStarts := make([]C.int, len(arrayOfStarts))

	for i := 0; i < len(arrayOfSizes); i++ {
		cArrayOfSizes[i] = C.int(arrayOfSizes[i])
	}

	for i := 0; i < len(arrayOfSubsizes); i++ {
		cArrayOfSubsizes[i] = C.int(arrayOfSubsizes[i])
	}

	for i := 0; i < len(arrayOfStarts); i++ {
		cArrayOfStarts[i] = C.int(arrayOfStarts[i])
	}

	var newtype C.MPI_Datatype

	err := C.MPI_Type_create_subarray(C.int(ndims),
		&cArrayOfSizes[0],
		&cArrayOfSubsizes[0],
		&cArrayOfStarts[0],
		C.int(order),
		C.MPI_Datatype(oldType),
		&newtype)

	return Datatype(newtype), int(err)
}

//Type_create_resized
//Create a datatype with a new lower bound and extent from an existing datatype.
func Type_create_resized(oldType Datatype, lowerBound Aint, extent Aint) (Datatype, int) {

	var newtype C.MPI_Datatype

	err := C.MPI_Type_create_resized(C.MPI_Datatype(oldType),
		C.MPI_Aint(lowerBound),
		C.MPI_Aint(extent),
		&newtype)

	return Datatype(newtype), int(err)
}

//Type_delete_attr
//Deletes an attribute value associated with a key on a datatype.
func Type_delete_attr(datatype Datatype, datatypeKeyval int) int {
	return int(C.MPI_Type_delete_attr(C.MPI_Datatype(datatype), C.int(datatypeKeyval)))
}

//Type_dup
//Duplicate a datatype.
func Type_dup(datatype Datatype) (Datatype, int) {

	var newtype C.MPI_Datatype
	err := C.MPI_Type_dup(C.MPI_Datatype(datatype), &newtype)

	return Datatype(newtype), int(err)
}

//Type_free_keyval
//Frees an attribute key for datatypes.
func Type_free_keyval(typeKeyval int) (int, int) {

	cTypeKeyval := C.int(typeKeyval)
	err := C.MPI_Type_free_keyval(&cTypeKeyval)

	return int(cTypeKeyval), int(err)
}

//Type_get_attr
//Retrieves attribute value by key.
func Type_get_attr(datatype Datatype, datatypeKeyval int) (uintptr, bool, int) {

	flag := false
	var cFlag C.int
	var attributeVal uintptr
	err := C.MPI_Type_get_attr(C.MPI_Datatype(datatype),
		C.int(datatypeKeyval),
		unsafe.Pointer(&attributeVal),
		&cFlag)

	if 0 != C.int(cFlag) {
		flag = true
	}

	return attributeVal, flag, int(err)
}

//Type_get_contents
//get type contents.
func Type_get_contents(datatype Datatype,
	maxIntegers int,
	maxAddresses int,
	maxDatatypes int) ([]int, []Aint, []Datatype, int) {

	cArrayOfIntegers := make([]C.int, maxIntegers)
	arrayOfIntegers := make([]int, maxIntegers)
	cArrayOfAddresses := make([]C.MPI_Aint, maxAddresses)
	arrayOfAddresses := make([]Aint, maxAddresses)
	cArrayOfDatatypes := make([]C.MPI_Datatype, maxDatatypes)
	arrayOfDatatypes := make([]Datatype, maxDatatypes)

	err := C.MPI_Type_get_contents(C.MPI_Datatype(datatype),
		C.int(maxIntegers),
		C.int(maxAddresses),
		C.int(maxDatatypes),
		&cArrayOfIntegers[0],
		&cArrayOfAddresses[0],
		&cArrayOfDatatypes[0])

	for i := 0; i < maxIntegers; i++ {
		arrayOfIntegers[i] = int(cArrayOfIntegers[i])
	}
	for i := 0; i < maxAddresses; i++ {
		arrayOfAddresses[i] = Aint(cArrayOfAddresses[i])
	}
	for i := 0; i < maxDatatypes; i++ {
		arrayOfDatatypes[i] = Datatype(cArrayOfDatatypes[i])
	}

	return arrayOfIntegers, arrayOfAddresses, arrayOfDatatypes, int(err)

}

//Type_get_envelope
//get type envelope.
func Type_get_envelope(datatype Datatype) (int, int, int, int, int) {
	var numIntegers C.int
	var numAddresses C.int
	var numDatatypes C.int
	var combiner C.int

	err := C.MPI_Type_get_envelope(C.MPI_Datatype(datatype),
		&numIntegers,
		&numAddresses,
		&numDatatypes,
		&combiner)

	return int(numIntegers), int(numAddresses), int(numDatatypes), int(combiner), int(err)
}

//Type_get_extent
//Get the lower bound and extent for a Datatype.
func Type_get_extent(datatype Datatype) (Aint, Aint, int) {

	var lb C.MPI_Aint
	var extent C.MPI_Aint
	err := C.MPI_Type_get_extent(C.MPI_Datatype(datatype), &lb, &extent)

	return Aint(lb), Aint(extent), int(err)
}

//Type_get_name
//Get the print name for a datatype.
func Type_get_name(datatype Datatype) (string, int) {

	var cLength C.int
	var cTypeName (*C.char)

	cTypeName = (*C.char)(C.malloc(C.size_t(MAX_OBJECT_NAME)))

	err := C.MPI_Type_get_name(C.MPI_Datatype(datatype),
		cTypeName,
		&cLength)

	typeName := C.GoStringN(cTypeName, cLength)

	C.free(unsafe.Pointer(cTypeName))

	return typeName, int(err)

}

//Type_get_true_extent
//Get the true lower bound and extent for a datatype.
func Type_get_true_extent(datatype Datatype) (Aint, Aint, int) {

	var trueLB C.MPI_Aint
	var trueExtent C.MPI_Aint
	err := C.MPI_Type_get_true_extent(C.MPI_Datatype(datatype),
		&trueLB,
		&trueExtent)

	return Aint(trueLB), Aint(trueExtent), int(err)
}

//Type_indexed
//Creates an indexed datatype.
func Type_indexed(blocklLengths []int, indices []int, oldType Datatype) (Datatype, int) {

	if len(blocklLengths) != len(indices) {
		fmt.Println("[ERROR] In Type_get_indexed: Length of blocklLengths differs from length of indices.")
		os.Exit(1)
	}
	count := len(indices)
	var newType C.MPI_Datatype

	cArrayOfBlockLengths := make([]C.int, count)
	cArrayOfIndices := make([]C.int, count)

	for i := 0; i < count; i++ {
		cArrayOfBlockLengths[i] = C.int(blocklLengths[i])
		cArrayOfIndices[i] = C.int(indices[i])
	}

	err := C.MPI_Type_indexed(C.int(count),
		&cArrayOfBlockLengths[0],
		&cArrayOfIndices[0],
		C.MPI_Datatype(oldType),
		&newType)

	return Datatype(newType), int(err)
}

//Type_match_size
//Find an MPI datatype matching a specified size.
func Type_match_size(typeClass int, size int) (Datatype, int) {

	var datatype C.MPI_Datatype
	err := C.MPI_Type_match_size(C.int(typeClass),
		C.int(size),
		&datatype)

	return Datatype(datatype), int(err)
}

//Type_set_attr
//Stores attribute value associated with a key.
func Type_set_attr(datatype Datatype, typeKeyval int, attributeVal interface{}) int {

	attributeValVoidPointer := Get_void_ptr(attributeVal)
	err := C.MPI_Type_set_attr(C.MPI_Datatype(datatype),
		C.int(typeKeyval),
		attributeValVoidPointer)

	return int(err)
}

//Type_set_name
//set datatype name.
func Type_set_name(datatype Datatype, typeName string) int {

	cstrTypeName := C.CString(typeName)
	err := C.MPI_Type_set_name(C.MPI_Datatype(datatype),
		cstrTypeName)

	C.free(unsafe.Pointer(cstrTypeName))

	return int(err)
}

func Type_size(datatype Datatype) (int, int) {

	var cSize C.int

	err := C.MPI_Type_size(C.MPI_Datatype(datatype),
		&cSize)

	return int(cSize), int(err)
}

//Type_vector
//Creates a vector (strided) datatype
func Type_vector(count int, blockLength int, stride int, oldType Datatype) (Datatype, int) {

	var newtype C.MPI_Datatype

	err := C.MPI_Type_vector(C.int(count),
		C.int(blockLength),
		C.int(stride),
		C.MPI_Datatype(oldType),
		&newtype)

	return Datatype(newtype), int(err)
}

//Type_free
//Frees the datatype
func Type_free(datatype *Datatype) int {

	err := C.MPI_Type_free((*C.MPI_Datatype)(unsafe.Pointer(datatype)))

	return int(err)
}
