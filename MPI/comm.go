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

//Accumulate
//Combines the contents of the origin buffer with that of a target buffer.
func Accumulate(originAddr interface{},
	originCount int,
	originDatatype Datatype,
	targetRank int,
	targetDisp Aint,
	targetCount int,
	targetDatatype Datatype,
	op Op,
	win Win) int {

	originAddrVoidPointer := Get_void_ptr(originAddr)
	return int(
		C.MPI_Accumulate(originAddrVoidPointer,
			C.int(originCount),
			C.MPI_Datatype(originDatatype),
			C.int(targetRank),
			C.MPI_Aint(targetDisp),
			C.int(targetCount),
			C.MPI_Datatype(targetDatatype),
			C.MPI_Op(op),
			C.MPI_Win(win)))
}

//Add_error_class
//Creates a new error class and returns its value
func Add_error_class() (int, int) {

	var errorClass C.int
	err := C.MPI_Add_error_class(&errorClass)

	return int(errorClass), int(err)
}

//Add_error_code
//Creates a new error code associated with errorclass
func Add_error_code(errorClass int) (int, int) {

	var errorCode C.int
	err := C.MPI_Add_error_code(C.int(errorClass), &errorCode)

	return int(errorCode), int(err)
}

//Add_error_string
//Associates a string with an error code or class
func Add_error_string(errorCode int, errorString string) int {

	err := C.MPI_Add_error_string(C.int(errorCode), C.CString(errorString))

	return int(err)
}

//Allgather
//Gathers data from all processes and distributes it to all processes.
func Allgather(sendBuffer interface{},
	sendCount int,
	sendType Datatype,
	recvBuffer interface{},
	recvCount int,
	recvType Datatype,
	comm Comm) int {

	sendBufferVoidPointer := Get_void_ptr(sendBuffer)
	recvBufferVoidPointer := Get_void_ptr(recvBuffer)
	err := C.MPI_Allgather(sendBufferVoidPointer,
		C.int(sendCount),
		C.MPI_Datatype(sendType),
		recvBufferVoidPointer,
		C.int(recvCount),
		C.MPI_Datatype(recvType),
		C.MPI_Comm(comm))

	return int(err)
}

//Allgatherv
//Gathers data from all processes and delivers it to all.
//Each process may contribute a different amount of data.
func Allgatherv(sendBuffer interface{},
	sendCount int,
	sendType Datatype,
	recvBuffer interface{},
	recvCount []int,
	displacements []int,
	recvType Datatype,
	comm Comm) int {

	sendBufferVoidPointer := Get_void_ptr(sendBuffer)
	recvBufferVoidPointer := Get_void_ptr(recvBuffer)

	err := C.MPI_Allgatherv(sendBufferVoidPointer,
		C.int(sendCount),
		C.MPI_Datatype(sendType),
		recvBufferVoidPointer,
		(*C.int)(unsafe.Pointer(&recvCount[0])),
		(*C.int)(unsafe.Pointer(&displacements[0])),
		C.MPI_Datatype(recvType),
		C.MPI_Comm(comm))

	return int(err)
}

//Gather
//Gathers together values from a group of processes
func Gather(sendBuffer interface{},
	sendCount int,
	sendType Datatype,
	recvBuffer interface{},
	recvCount int,
	recvType Datatype,
	rootRank int,
	comm Comm) int {

	sendBufferVoidPointer := Get_void_ptr(sendBuffer)
	recvBufferVoidPointer := Get_void_ptr(recvBuffer)

	err := C.MPI_Gather(sendBufferVoidPointer,
		C.int(sendCount),
		C.MPI_Datatype(sendType),
		recvBufferVoidPointer,
		C.int(recvCount),
		C.MPI_Datatype(recvType),
		C.int(rootRank),
		C.MPI_Comm(comm))

	return int(err)
}

//Gatherv
//Gathers into specified locations from all processes in a group
func Gatherv(sendBuffer interface{},
	sendCount int,
	sendType Datatype,
	recvBuffer interface{},
	recvCount []int,
	displacements []int,
	recvType Datatype,
	rootRank int,
	comm Comm) int {

	sendBufferVoidPointer := Get_void_ptr(sendBuffer)
	recvBufferVoidPointer := Get_void_ptr(recvBuffer)

	err := C.MPI_Gatherv(sendBufferVoidPointer,
		C.int(sendCount),
		C.MPI_Datatype(sendType),
		recvBufferVoidPointer,
		(*C.int)(unsafe.Pointer(&recvCount[0])),
		(*C.int)(unsafe.Pointer(&displacements[0])),
		C.MPI_Datatype(recvType),
		C.int(rootRank),
		C.MPI_Comm(comm))

	return int(err)
}

//Alltoallv
//All processes send different amount of data to,
//and receive different amount of data from, all processes
func Alltoallv(sendBuffer interface{},
	sendCounts []int,
	sendDisplacements []int,
	sendType Datatype,
	recvBuffer interface{},
	recvCounts []int,
	recvDisplacements []int,
	recvType Datatype,
	comm Comm) int {

	sendBufferVoidPointer := Get_void_ptr(sendBuffer)
	recvBufferVoidPointer := Get_void_ptr(recvBuffer)

	err := C.MPI_Alltoallv(sendBufferVoidPointer,
		(*C.int)(unsafe.Pointer(&sendCounts[0])),
		(*C.int)(unsafe.Pointer(&sendDisplacements[0])),
		C.MPI_Datatype(sendType),
		recvBufferVoidPointer,
		(*C.int)(unsafe.Pointer(&recvCounts[0])),
		(*C.int)(unsafe.Pointer(&recvDisplacements[0])),
		C.MPI_Datatype(recvType),
		C.MPI_Comm(comm))

	return int(err)
}

//Alltoallw
//All processes send data of different types to,
//and receive data of different types from, all processes
func Alltoallw(sendBuffer interface{},
	sendCounts []int,
	sendDisplacements []int,
	sendTypes []Datatype,
	recvBuffer interface{},
	recvCounts []int, recvDisplacements []int, recvTypes []Datatype, comm Comm) int {

	sendBufferVoidPointer := Get_void_ptr(sendBuffer)
	recvBufferVoidPointer := Get_void_ptr(recvBuffer)

	err := C.MPI_Alltoallw(sendBufferVoidPointer,
		(*C.int)(unsafe.Pointer(&sendCounts[0])),
		(*C.int)(unsafe.Pointer(&sendDisplacements[0])),
		(*C.MPI_Datatype)(unsafe.Pointer(&sendTypes[0])),
		recvBufferVoidPointer,
		(*C.int)(unsafe.Pointer(&recvCounts[0])),
		(*C.int)(unsafe.Pointer(&recvDisplacements[0])),
		(*C.MPI_Datatype)(unsafe.Pointer(&recvTypes[0])),
		C.MPI_Comm(comm))

	return int(err)
}

//Bsend_init
//Builds a handle for a buffered send.
func MPI_Bsend_init(sendBuffer interface{},
	count int,
	dataType Datatype,
	dest int,
	tag int,
	comm Comm) (Request, int) {

	var request C.MPI_Request
	sendBufferVoidPointer := Get_void_ptr(sendBuffer)

	err := C.MPI_Bsend_init(sendBufferVoidPointer,
		C.int(count),
		C.MPI_Datatype(dataType),
		C.int(dest),
		C.int(tag),
		C.MPI_Comm(comm),
		&request)

	return Request(request), int(err)
}

//Comm_group
//Returns the group associated with a communicator.
func Comm_group(comm Comm) (Group, int) {

	var group C.MPI_Group
	err := C.MPI_Comm_group(C.MPI_Comm(comm), &group)

	return Group(group), int(err)

}

//Comm_join
//Establishes communication between MPI jobs
func Comm_join(fd int) (Comm, int) {

	var interComm C.MPI_Comm
	err := C.MPI_Comm_join(C.int(fd), &interComm)

	return Comm(interComm), int(err)
}

//Comm_accept
//Establishes communication with a client.
func Comm_accept(portName string, info Info, root int, comm Comm) (Comm, int) {

	var newComm C.MPI_Comm
	err := C.MPI_Comm_accept(C.CString(portName),
		C.MPI_Info(info),
		C.int(root),
		C.MPI_Comm(comm),
		&newComm)

	return Comm(newComm), int(err)
}

//Comm_call_errhandler
//Passes the supplied error code to the
func Comm_call_errhandler(comm Comm, errorCode int) int {

	err := C.MPI_Comm_call_errhandler(C.MPI_Comm(comm), C.int(errorCode))

	return int(err)
}

//Comm_compare
//Compares two communicators.
func Comm_compare(comm1 Comm, comm2 Comm) (int, int) {

	var result C.int
	err := C.MPI_Comm_compare(C.MPI_Comm(comm1), C.MPI_Comm(comm2), &result)

	return int(result), int(err)
}

//Comm_connect
//Establishes communication with a server.
func Comm_connect(portName string, info Info, root int, comm Comm) (Comm, int) {

	var newComm C.MPI_Comm
	err := C.MPI_Comm_connect(C.CString(portName),
		C.MPI_Info(info),
		C.int(root),
		C.MPI_Comm(comm),
		&newComm)

	return Comm(newComm), int(err)
}

//Comm_create
//Creates a new communicator.
func Comm_create(comm Comm, group Group) (Comm, int) {

	var newComm C.MPI_Comm
	err := C.MPI_Comm_create(C.MPI_Comm(comm), C.MPI_Group(group), &newComm)

	return Comm(newComm), int(err)
}

//Comm_disconnect
//Deallocates communicator object and sets handle to MPI_COMM_NULL.
func Comm_disconnect(comm *Comm) int {

	err := C.MPI_Comm_disconnect((*C.MPI_Comm)(unsafe.Pointer(comm)))

	return int(err)
}

//Comm_free
//Mark a communicator object for deallocation.
func Comm_free(comm *Comm) int {

	err := C.MPI_Comm_free((*C.MPI_Comm)(unsafe.Pointer(comm)))

	return int(err)
}

//Comm_rank
//Determines the rank of the calling process in the communicator.
func Comm_rank(comm Comm) (int, int) {

	var rank C.int
	err := C.MPI_Comm_rank(C.MPI_Comm(comm), &rank)

	return int(rank), int(err)
}

//Comm_size
//Returns the size of the group associated with a communicator.
func Comm_size(comm Comm) (int, int) {

	var size C.int
	err := C.MPI_Comm_size(C.MPI_Comm(comm), &size)

	return int(size), int(err)
}

//Comm_split
//Creates new communicators based on colors and keys.
func Comm_split(comm Comm, color, key int) (Comm, int) {

	var newComm C.MPI_Comm
	err := C.MPI_Comm_split(C.MPI_Comm(comm), C.int(color), C.int(key), &newComm)

	return Comm(newComm), int(err)
}

//Comm_dup
//Duplicates an existing communicator with all its cached information.
func Comm_dup(comm Comm) (Comm, int) {

	var newComm C.MPI_Comm
	err := C.MPI_Comm_dup(C.MPI_Comm(comm), &newComm)

	return Comm(newComm), int(err)

}

//Send
//Performs a standard-mode blocking send.
func Send(buffer interface{},
	count int,
	dataType Datatype,
	dest int,
	tag int,
	comm Comm) int {

	bufferVoidPointer := Get_void_ptr(buffer)

	err := C.MPI_Send(bufferVoidPointer,
		C.int(count),
		C.MPI_Datatype(dataType),
		C.int(dest),
		C.int(tag),
		C.MPI_Comm(comm))

	return int(err)
}

//Recv
//Performs a standard-mode blocking receive.
func Recv(buffer interface{},
	count int,
	dataType Datatype,
	source int,
	tag int,
	comm Comm) (Status, int) {

	var status C.MPI_Status
	bufferVoidPointer := Get_void_ptr(buffer)

	err := C.MPI_Recv(bufferVoidPointer,
		C.int(count),
		C.MPI_Datatype(dataType),
		C.int(source),
		C.int(tag),
		C.MPI_Comm(comm),
		&status)

	return Status(status), int(err)
}

//Barrier
//Blocks until all processes have reached this routine.
func Barrier(comm Comm) int {

	err := C.MPI_Barrier(C.MPI_Comm(comm))

	return int(err)
}

//Bcast
//Broadcasts a message from the process with rank root to all other processes of the group.
func Bcast(buffer interface{},
	count int,
	dataType Datatype,
	root int,
	comm Comm) int {

	bufferVoidPointer := Get_void_ptr(buffer)

	err := C.MPI_Bcast(bufferVoidPointer,
		C.int(count),
		C.MPI_Datatype(dataType),
		C.int(root),
		C.MPI_Comm(comm))

	return int(err)
}

//Bsend
//Basic send with user-specified buffering.
func Bsend(buffer interface{},
	count int,
	dataType Datatype,
	dest int,
	tag int,
	comm Comm) int {

	bufferVoidPointer := Get_void_ptr(buffer)

	err := C.MPI_Bsend(bufferVoidPointer,
		C.int(count),
		C.MPI_Datatype(dataType),
		C.int(dest),
		C.int(tag),
		C.MPI_Comm(comm))

	return int(err)
}

//MPI_Probe
//Blocking test for a message.
func Probe(source int, tag int, comm Comm) (Status, int) {

	var status C.MPI_Status
	err := C.MPI_Probe(
		C.int(source),
		C.int(tag),
		C.MPI_Comm(comm),
		&status)

	return Status(status), int(err)
}

//Get_count
//Gets the number of top-level elements received.
func Get_count(status Status, dataType Datatype) (int, int) {

	var count C.int
	var cStatus C.MPI_Status
	cStatus = C.MPI_Status(status)

	err := C.MPI_Get_count(&cStatus, C.MPI_Datatype(dataType), &count)

	return int(count), int(err)
}

//Allreduce
//Combines values from all processes and distributes the result back to all processes.
func Allreduce(sendBuffer interface{},
	recvBuffer interface{},
	count int,
	dataType Datatype,
	op Op,
	comm Comm) int {

	sendBufferVoidPointer := Get_void_ptr(sendBuffer)
	recvBufferVoidPointer := Get_void_ptr(recvBuffer)

	err := C.MPI_Allreduce(sendBufferVoidPointer,
		recvBufferVoidPointer,
		C.int(count),
		C.MPI_Datatype(dataType),
		C.MPI_Op(op),
		C.MPI_Comm(comm))

	return int(err)
}

//Alltoall
//All processes send data to all processes
func Alltoall(sendBuffer interface{},
	sendCount int,
	sendType Datatype,
	recvBuffer interface{},
	recvCount int,
	recvType Datatype,
	comm Comm) int {

	sendBufferVoidPointer := Get_void_ptr(sendBuffer)
	recvBufferVoidPointer := Get_void_ptr(recvBuffer)

	err := C.MPI_Alltoall(sendBufferVoidPointer,
		C.int(sendCount),
		C.MPI_Datatype(sendType),
		recvBufferVoidPointer,
		C.int(recvCount),
		C.MPI_Datatype(recvType),
		C.MPI_Comm(comm))

	return int(err)
}

//Isend
//Starts a standard-mode, nonblocking send.
func Isend(buffer interface{},
	sendCount int,
	dataType Datatype,
	dest int,
	tag int,
	comm Comm,
	request *Request) int {

	bufferVoidPointer := Get_void_ptr(buffer)
	cRequestPointer := (*C.MPI_Request)(request)

	err := C.MPI_Isend(bufferVoidPointer,
		C.int(sendCount),
		C.MPI_Datatype(dataType),
		C.int(dest),
		C.int(tag),
		C.MPI_Comm(comm),
		cRequestPointer)

	return int(err)
}

// Irecv
// Begins a nonblocking receive
func Irecv(buffer interface{},
	count int,
	datatype Datatype,
	source int,
	tag int,
	comm Comm,
	request *Request) int {

	recvBufferVoidPointer := Get_void_ptr(buffer)

	err := C.MPI_Irecv(recvBufferVoidPointer,
		C.int(count),
		C.MPI_Datatype(datatype),
		C.int(source),
		C.int(tag),
		C.MPI_Comm(comm),
		(*C.MPI_Request)(request))
	return int(err)
}

//Wait
//Waits for an MPI send or receive to complete.
func Wait(request *Request) (Status, int) {

	var status C.MPI_Status
	var cRequest *C.MPI_Request
	cRequest = (*C.MPI_Request)(request)

	err := C.MPI_Wait(cRequest, &status)

	return Status(status), int(err)
}

//Waitall
//Waits for all given communications to complete.
func Waitall(sliceOfRequests []Request) ([]Status, int) {

	length := len(sliceOfRequests)
	cSliceOfRequests := (*C.MPI_Request)(&sliceOfRequests[0])
	sliceOfStatuses := make([]Status, length)
	cSliceOfStatuses := make([]C.MPI_Status, length)

	err := C.MPI_Waitall(C.int(length), cSliceOfRequests, &cSliceOfStatuses[0])

	for i := 0; i < length; i++ {
		sliceOfStatuses[i] = Status(cSliceOfStatuses[i])
	}

	return sliceOfStatuses, int(err)
}

//MPI_Waitany
//Waits for any specified send or receive to complete.
func Waitany(sliceOfRequests []Request) (int, Status, int) {

	length := len(sliceOfRequests)
	cSliceOfRequests := (*C.MPI_Request)(&sliceOfRequests[0])
	var index C.int
	var status C.MPI_Status

	err := C.MPI_Waitany(C.int(length), cSliceOfRequests, &index, &status)

	return int(index), Status(status), int(err)

}

//MPI_Waitsome
//Waits for some given communications to complete.
func Waitsome(sliceOfRequests []Request) (int, []int, []Status, int) {

	length := len(sliceOfRequests)
	cSliceOfRequests := (*C.MPI_Request)(&sliceOfRequests[0])
	cSliceOfIndices := make([]C.int, length)
	cSliceOfStatuses := make([]C.MPI_Status, length)
	sliceOfIndices := make([]int, length)
	sliceOfStatuses := make([]Status, length)
	var count C.int

	err := C.MPI_Waitsome(C.int(length), cSliceOfRequests, &count, &cSliceOfIndices[0], &cSliceOfStatuses[0])

	for i := 0; i < length; i++ {
		sliceOfIndices[i] = int(cSliceOfIndices[i])
		sliceOfStatuses[i] = Status(cSliceOfStatuses[i])
	}

	return int(count), sliceOfIndices, sliceOfStatuses, int(err)
}

// int MPI_Comm_create_keyval(MPI_Comm_copy_attr_function *comm_copy_attr_fn, MPI_Comm_delete_attr_function *comm_delete_attr_fn, int *comm_keyval, void *extra_state);

// Comm_delete_attr
// Deletes an attribute value associated with a key on a communicator
func Comm_delete_attr(comm Comm, commKeyval int) int {

	return int(C.MPI_Comm_delete_attr(C.MPI_Comm(comm), C.int(commKeyval)))
}

// Comm_free_keyval
// Frees an attribute key for communicators
func Comm_free_keyval(commKeyval int) (int, int) {

	var cCommKeyval C.int
	cCommKeyval = C.int(commKeyval)
	err := C.MPI_Comm_free_keyval(&cCommKeyval)
	return int(cCommKeyval), int(err)
}

//Comm_get_attr
//Retrieves attribute value by key
func Comm_get_attr(comm Comm, commKeyval int) (uintptr, bool, int) {

	var attributeVal uintptr
	var cflag C.int
	flag := false
	err := C.MPI_Comm_get_attr(C.MPI_Comm(comm),
		C.int(commKeyval),
		unsafe.Pointer(attributeVal),
		&cflag)

	if int(cflag) != 0 {
		flag = true
	}

	return attributeVal, flag, int(err)
}

//Comm_get_errhandler
//Get the error handler attached to a communicator
func Comm_get_errhandler(comm Comm) (Errhandler, int) {

	var errhandler C.MPI_Errhandler
	err := C.MPI_Comm_get_errhandler(C.MPI_Comm(comm), &errhandler)

	return Errhandler(errhandler), int(err)
}

//Comm_get_name
//Return the print name from the communicator
func Comm_get_name(comm Comm) (string, int) {

	// var cstr *C.char
	var strLength C.int
	var commName string
	cstr := C.malloc(C.size_t(MAX_BUFFER_SIZE))

	err := C.MPI_Comm_get_name(C.MPI_Comm(comm), (*C.char)(cstr), &strLength)

	commName = C.GoStringN((*C.char)(cstr), strLength)
	C.free(cstr)
	return commName, int(err)
}

//Comm_get_parent
//Return the parent communicator for this process
func Comm_get_parent() (Comm, int) {

	var comm C.MPI_Comm
	err := C.MPI_Comm_get_parent(&comm)

	return Comm(comm), int(err)
}

//Comm_remote_group
//Accesses the remote group associated with the given inter-communicator
func Comm_remote_group(comm Comm) (Group, int) {

	var group C.MPI_Group
	err := C.MPI_Comm_remote_group(C.MPI_Comm(comm), &group)

	return Group(group), int(err)
}

//Comm_remote_size
//Determines the size of the remote group associated with an inter-communictor
func Comm_remote_size(comm Comm) (int, int) {

	var size C.int
	err := C.MPI_Comm_remote_size(C.MPI_Comm(comm), &size)

	return int(size), int(err)
}

//Comm_set_attr
//Stores attribute value associated with a key
func Comm_set_attr(comm Comm, commKeyval int, attributeVal interface{}) int {

	attributeValVoidPointer := Get_void_ptr(attributeVal)

	return int(C.MPI_Comm_set_attr(C.MPI_Comm(comm),
		C.int(commKeyval),
		attributeValVoidPointer))
}

//Comm_set_errhandler
//Set the error handler for a communicator
func Comm_set_errhandler(comm Comm, errhandler Errhandler) (Comm, int) {

	cComm := C.MPI_Comm(comm)
	cErrhandler := C.MPI_Errhandler(errhandler)
	err := C.MPI_Comm_set_errhandler(cComm, cErrhandler)

	return Comm(cComm), int(err)
}

//Comm_set_name
//Sets the print name for a communicator
func Comm_set_name(comm Comm, commName string) int {

	cCommName := C.CString(commName)

	err := C.MPI_Comm_set_name(C.MPI_Comm(comm),
		cCommName)

	C.free(unsafe.Pointer(cCommName))
	return int(err)
}

//Comm_spawn
//Spawn up to maxprocs instances of a single MPI application
func Comm_spawn(command string, arguments []string, maxProcs int,
	info Info, root int, comm Comm) (Comm, []int, int) {

	cCommand := C.CString(command)
	defer C.free(unsafe.Pointer(cCommand))

	numberOfArguments := len(arguments)
	var cArguments [](*C.char)

	var intercomm C.MPI_Comm
	cArrayOfErrorcodes := make([]C.int, maxProcs)
	arrayOfErrorcodes := make([]int, maxProcs)

	for i := 0; i < numberOfArguments; i++ {
		cArguments[i] = C.CString(arguments[i])
	}

	err := C.MPI_Comm_spawn(cCommand,
		&cArguments[0],
		C.int(maxProcs),
		C.MPI_Info(info),
		C.int(root),
		C.MPI_Comm(comm),
		&intercomm,
		(*C.int)(&cArrayOfErrorcodes[0]))

	for i := 0; i < maxProcs; i++ {
		arrayOfErrorcodes[i] = int(cArrayOfErrorcodes[i])
	}

	for i := 0; i < numberOfArguments; i++ {
		C.free(unsafe.Pointer(cArguments[i]))
	}

	return Comm(intercomm), arrayOfErrorcodes, int(err)

}

//TODO: Double/Triple Pointers make trouble... I think we have to
//do this in c!

//Comm_spawn_multiple
//While MPI_COMM_SPAWN is sufficient for most cases, it does not
//allow the spawning of multiple binaries, or of the same binary
//with multiple sets of arguments. The following routine spawns
//multiple binaries or the same binary with multiple sets of arguments,
//establishing communication with them and placing them in the
//same MPI_COMM_WORLD.
// func Comm_spawn_multiple(count int, commands []string, arguments [][]string,
// 	maxProcs []int, infos []Info, root int, comm Comm) (Comm, []int, int) {

// 	numberOfCommands := len(commands)

// 	cCommands := make([](*C.char), numberOfCommands)
// 	cMaxProcs := make([]C.int, numberOfCommands)
// 	cInfos := make([]C.MPI_Info, numberOfCommands)

// 	cArguments := make([]([]*C.char), numberOfCommands)

// 	var totalProcs C.int
// 	totalProcs = 0

// 	for i := 0; i < numberOfCommands; i++ {
// 		cCommands[i] = C.CString(commands[i])
// 		cMaxProcs[i] = C.int(maxProcs[i])
// 		totalProcs += C.int(maxProcs[i])
// 		cInfos[i] = C.MPI_Info(infos[i])

// 		cArguments[i] = make([](*C.char), len(arguments[i]))
// 		for j := 0; j < len(arguments[i]); j++ {
// 			cArguments[i][j] = C.CString(arguments[i][j])
// 		}
// 	}

// 	var intercomm C.MPI_Comm
// 	cArrayOfErrorcodes := make([]C.int, totalProcs)
// 	arrayOfErrorcodes := make([]int, totalProcs)

// 	err := C.MPI_Comm_spawn_multiple(C.int(count),
// 		&cCommands[0],
// 		&((**C.char)(cArguments[0])),
// 		&cMaxProcs[0],
// 		&cInfos[0],
// 		C.int(root),
// 		C.MPI_Comm(comm),
// 		&intercomm,
// 		(*C.int)(&cArrayOfErrorcodes[0]))

// 	for i := 0; i < int(totalProcs); i++ {
// 		arrayOfErrorcodes[i] = int(cArrayOfErrorcodes[i])
// 	}

// 	for i := 0; i < numberOfCommands; i++ {
// 		C.free(unsafe.Pointer(cCommands[i]))
// 		for j := 0; j < len(arguments[i]); j++ {
// 			C.free(unsafe.Pointer(cArguments[i][j]))
// 		}
// 	}

// 	return Comm(intercomm), arrayOfErrorcodes, int(err)
// }

//Comm_test_inter
//Tests to see if a comm is an inter-communicator
func Comm_test_inter(comm Comm) (bool, int) {

	var cFlag C.int
	flag := false
	err := C.MPI_Comm_test_inter(C.MPI_Comm(comm), &cFlag)

	if int(cFlag) != 0 {
		flag = true
	}

	return flag, int(err)
}
