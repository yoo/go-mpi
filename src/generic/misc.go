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
// Workaround for parameter list problems
// since Cgo can't handle ...
int MPI_Pcontrol_ (const int level) {
	return MPI_Pcontrol(level);
}
*/
import "C"

import (
	"unsafe"
)

//Wtick
//Returns the resolution of MPI_Wtime
func Wtick() float64 {
	return float64(C.MPI_Wtick())
}

//Wtime
//Returns an elapsed time on the calling processor
func Wtime() float64 {
	return float64(C.MPI_Wtime())
}

//Op_free
//Frees a user-defined combination function handle.
func Op_free(op Op) (Op, int) {

	cOp := C.MPI_Op(op)
	err := C.MPI_Op_free(&cOp)

	return Op(cOp), int(err)
}

//Query_thread
//Return the level of thread support provided by the MPI library.
func Query_thread() (int, int) {

	var provided C.int

	err := C.MPI_Query_thread(&provided)

	return int(provided), int(err)
}

//Request_free
//Frees a communication request object.
func Request_free(request Request) (Request, int) {

	cRequest := C.MPI_Request(request)
	err := C.MPI_Request_free(&cRequest)

	return Request(cRequest), int(err)
}

//Start
//Initiates a communication with a persistent request handle.
func Start(request Request) (Request, int) {

	cRequest := C.MPI_Request(request)

	err := C.MPI_Start(&cRequest)

	return Request(cRequest), int(err)
}

//Startall
//Starts a collection of persistent requests.
func Startall(requests []Request) int {

	count := len(requests)
	cArrayOfRequests := make([]C.MPI_Request, count)

	for i := 0; i < len(requests); i++ {
		cArrayOfRequests[i] = C.MPI_Request(requests[i])
	}

	err := C.MPI_Startall(C.int(count), &cArrayOfRequests[0])

	return int(err)
}

//Test
//Tests for the completion of a request.
func Test(request Request) (bool, Status, int) {

	cRequest := C.MPI_Request(request)
	var cFlag C.int
	var cStatus C.MPI_Status
	flag := false

	err := C.MPI_Test(&cRequest, &cFlag, &cStatus)

	if int(cFlag) != 0 {
		flag = true
	}

	return flag, Status(cStatus), int(err)
}

//Test_cancelled
//Tests whether a request was canceled.
func Test_cancelled(status Status) (bool, int) {

	cStatus := C.MPI_Status(status)
	var cFlag C.int
	flag := false

	err := C.MPI_Test_cancelled(&cStatus, &cFlag)

	if int(cFlag) != 0 {
		flag = true
	}

	return flag, int(err)
}

//MPI_Topo_test
//Determines the type of topology (if any) associated with a communicator.
func Topo_test(comm Comm) (int, int) {

	var topoType C.int

	err := C.MPI_Topo_test(C.MPI_Comm(comm), &topoType)

	return int(topoType), int(err)
}

//Iprobe
//Nonblocking test for a message.
func Iprobe(source int, tag int, comm Comm) (bool, Status, int) {

	flag := false
	cFlag := C.int(0)
	var cStatus C.MPI_Status

	err := C.MPI_Iprobe(C.int(source),
		C.int(tag),
		C.MPI_Comm(comm),
		&cFlag,
		&cStatus)

	if int(cFlag) != 0 {
		flag = true
	}

	return flag, Status(cStatus), int(err)
}

//Recv_init
//Builds a handle for a receive.
func Recv_init(buffer interface{}, count int, datatype Datatype,
	source int, tag int, comm Comm) (Request, int) {

	bufferVoidPtr := Get_void_ptr(buffer)
	var cRequest C.MPI_Request

	err := C.MPI_Recv_init(bufferVoidPtr,
		C.int(count),
		C.MPI_Datatype(datatype),
		C.int(source),
		C.int(tag),
		C.MPI_Comm(comm),
		&cRequest)

	return Request(cRequest), int(err)
}

//Reduce
//Reduces values on all processes within a group.
func Reduce(sendBuffer interface{}, recvBuffer interface{}, count int,
	datatype Datatype, op Op, root int, comm Comm) int {

	sendBufferVoidPtr := Get_void_ptr(sendBuffer)
	recvBufferVoidPtr := Get_void_ptr(recvBuffer)

	err := C.MPI_Reduce(sendBufferVoidPtr,
		recvBufferVoidPtr,
		C.int(count),
		C.MPI_Datatype(datatype),
		C.MPI_Op(op),
		C.int(root),
		C.MPI_Comm(comm))

	return int(err)
}

//Reduce_local
//Perform a local reduction
func Reduce_local(inbuff interface{}, inoutbuff interface{}, count int,
	datatype Datatype, op Op) int {

	inBufferVoidPtr := Get_void_ptr(inbuff)
	inoutBufferVoidPtr := Get_void_ptr(inoutbuff)

	err := C.MPI_Reduce_local(inBufferVoidPtr, inoutBufferVoidPtr,
		C.int(count),
		C.MPI_Datatype(datatype),
		C.MPI_Op(op))

	return int(err)
}

//Reduce_scatter
//Combines values and scatters the results.
func Reduce_scatter(sendBuffer interface{}, recvBuffer interface{}, recvCounts []int,
	datatype Datatype, op Op, comm Comm) int {

	sendBufferVoidPtr := Get_void_ptr(sendBuffer)
	recvBufferVoidPtr := Get_void_ptr(recvBuffer)

	cArrayOfRecvCounts := make([]C.int, len(recvCounts))
	for i := 0; i < len(recvCounts); i++ {
		cArrayOfRecvCounts[i] = C.int(recvCounts[i])
	}

	err := C.MPI_Reduce_scatter(sendBufferVoidPtr,
		recvBufferVoidPtr,
		&cArrayOfRecvCounts[0],
		C.MPI_Datatype(datatype),
		C.MPI_Op(op),
		C.MPI_Comm(comm))

	return int(err)
}

//Request_get_status
//Access information associated with a request without freeing the request.
func Requset_get_status(request Request) (bool, Status, int) {

	cFlag := C.int(0)
	flag := false
	var cStatus C.MPI_Status

	err := C.MPI_Request_get_status(C.MPI_Request(request), &cFlag, &cStatus)

	if int(cFlag) != 0 {
		flag = true
	}

	return flag, Status(cStatus), int(err)
}

//Rsend
//Ready send.
func Rsend(buf interface{}, count int, datatype Datatype, dst int, tag int, comm Comm) int {

	bufBufferVoidPtr := Get_void_ptr(buf)

	err := C.MPI_Rsend(bufBufferVoidPtr,
		C.int(count),
		C.MPI_Datatype(datatype),
		C.int(dst),
		C.int(tag),
		C.MPI_Comm(comm))

	return int(err)
}

//Rsend_init
//Builds a handle for a ready send.
func Rsend_init(buf interface{}, count int, datatype Datatype, dst int, tag int, comm Comm) (Request, int) {

	bufferVoidPtr := Get_void_ptr(buf)
	var cRequest C.MPI_Request

	err := C.MPI_Rsend_init(bufferVoidPtr,
		C.int(count),
		C.MPI_Datatype(datatype),
		C.int(dst),
		C.int(tag),
		C.MPI_Comm(comm),
		&cRequest)

	return Request(cRequest), int(err)

}

//Scan
//Computes an inclusive scan (partial reduction)
func Scan(sendBuffer interface{}, recvBuffer interface{}, count int, datatype Datatype,
	op Op, comm Comm) int {

	sendBufferVoidPtr := Get_void_ptr(sendBuffer)
	recvBufferVoidPtr := Get_void_ptr(recvBuffer)

	err := C.MPI_Scan(sendBufferVoidPtr,
		recvBufferVoidPtr,
		C.int(count),
		C.MPI_Datatype(datatype),
		C.MPI_Op(op),
		C.MPI_Comm(comm))

	return int(err)
}

//Scatterv
//Scatters a buffer in parts to all tasks in a group.
func Scatterv(sendBuffer interface{}, sendCounts []int, displacements []int, sendType Datatype,
	recvBuffer interface{}, recvCounts int, recvType Datatype, root int, comm Comm) int {

	sendBufferVoidPtr := Get_void_ptr(sendBuffer)
	recvBufferVoidPtr := Get_void_ptr(recvBuffer)
	cSendCounts := make([]C.int, len(sendCounts))
	cDisplacements := make([]C.int, len(displacements))

	for i := 0; i < len(sendCounts); i++ {
		cSendCounts[i] = C.int(sendCounts[i])
	}

	for i := 0; i < len(displacements); i++ {
		cDisplacements[i] = C.int(displacements[i])
	}

	err := C.MPI_Scatterv(sendBufferVoidPtr,
		&cSendCounts[0],
		&cDisplacements[0],
		C.MPI_Datatype(sendType),
		recvBufferVoidPtr,
		C.int(recvCounts),
		C.MPI_Datatype(recvType),
		C.int(root),
		C.MPI_Comm(comm))

	return int(err)
}

//Send_init
//Builds a handle for a standard send.
func Send_init(buffer interface{}, count int, datatype Datatype, dst int, tag int,
	comm Comm) (Request, int) {

	bufferVoidPtr := Get_void_ptr(buffer)
	var cRequest C.MPI_Request

	err := C.MPI_Send_init(bufferVoidPtr,
		C.int(count),
		C.MPI_Datatype(datatype),
		C.int(dst),
		C.int(tag),
		C.MPI_Comm(comm),
		&cRequest)

	return Request(cRequest), int(err)
}

//Sendrecv
//Sends and receives a message.
func Sendrecv(sendBuffer interface{}, sendCount int, sendType Datatype,
	dst int, sendTag int, recvBuffer interface{}, recvCount int, recvType Datatype,
	src int, recvTag int, comm Comm) (Status, int) {

	sendBufferVoidPtr := Get_void_ptr(sendBuffer)
	recvBufferVoidPtr := Get_void_ptr(recvBuffer)
	var cStatus C.MPI_Status

	err := C.MPI_Sendrecv(sendBufferVoidPtr,
		C.int(sendCount),
		C.MPI_Datatype(sendType),
		C.int(dst),
		C.int(sendTag),
		recvBufferVoidPtr,
		C.int(recvCount),
		C.MPI_Datatype(recvType),
		C.int(src),
		C.int(recvTag),
		C.MPI_Comm(comm),
		&cStatus)

	return Status(cStatus), int(err)
}

//Sendrecv_replace
//Sends and receives a message using a single buffer.
func Sendrecv_replace(buffer interface{}, count int, datatype Datatype, dst int,
	sendTag int, src int, recvTag int, comm Comm) (Status, int) {

	bufferVoidPtr := Get_void_ptr(buffer)
	var cStatus C.MPI_Status

	err := C.MPI_Sendrecv_replace(bufferVoidPtr,
		C.int(count),
		C.MPI_Datatype(datatype),
		C.int(dst),
		C.int(sendTag),
		C.int(src),
		C.int(recvTag),
		C.MPI_Comm(comm),
		&cStatus)

	return Status(cStatus), int(err)
}

//Ssend
//Standard synchronous send.
func Ssend(buffer interface{}, count int, datatype Datatype, dst int, tag int, comm Comm) int {

	bufferVoidPtr := Get_void_ptr(buffer)

	err := C.MPI_Ssend(bufferVoidPtr,
		C.int(count),
		C.MPI_Datatype(datatype),
		C.int(dst),
		C.int(tag),
		C.MPI_Comm(comm))

	return int(err)
}

//Ssend_init
//Builds a handle for a synchronous send.
func Ssend_init(buffer interface{}, count int, datatype Datatype, dst int,
	tag int, comm Comm) (Request, int) {

	bufferVoidPtr := Get_void_ptr(buffer)
	var cRequest C.MPI_Request

	err := C.MPI_Ssend_init(bufferVoidPtr,
		C.int(count),
		C.MPI_Datatype(datatype),
		C.int(dst),
		C.int(tag),
		C.MPI_Comm(comm),
		&cRequest)

	return Request(cRequest), int(err)
}

//Status_set_cancelled
//Sets status to indicate a request has been canceled.
func Status_set_canelled(status Status, flag bool) (Status, int) {

	cStatus := C.MPI_Status(status)
	cFlag := C.int(0)
	if flag {
		cFlag = C.int(1)
	}

	err := C.MPI_Status_set_cancelled(&cStatus, cFlag)

	return Status(cStatus), int(err)
}

//Status_set_elements
//Modifies opaque part of status to allow MPI_Get_elements to return count.
func Status_set_elements(status Status, datatype Datatype, count int) (Status, int) {

	cStatus := C.MPI_Status(status)

	err := C.MPI_Status_set_elements(&cStatus,
		C.MPI_Datatype(datatype),
		C.int(count))

	return Status(cStatus), int(err)
}

//Dims_create
//Creates a division of processors in a Cartesian grid.
func Dims_create(nnodes int, dims []int) ([]int, int) {

	ndims := len(dims)
	cDims := make([]C.int, ndims)
	for i := 0; i < ndims; i++ {
		cDims[i] = C.int(dims[i])
	}

	err := C.MPI_Dims_create(C.int(nnodes), C.int(ndims), &cDims[0])

	gDims := make([]int, ndims)
	for i := 0; i < ndims; i++ {
		gDims[i] = int(cDims[i])
	}

	return gDims, int(err)
}

//Errhandler_free
//Frees an MPI-style error handler.
func Errhandler_free(errhandler Errhandler) (Errhandler, int) {

	cErrhandler := C.MPI_Errhandler(errhandler)

	err := C.MPI_Errhandler_free(&cErrhandler)

	return Errhandler(cErrhandler), int(err)
}

//Error_class
//Converts an error code into an error class.
func Error_class(errorcode int) (int, int) {

	var cErrorClass C.int

	err := C.MPI_Error_class(C.int(errorcode), &cErrorClass)

	return int(cErrorClass), int(err)
}

//Exscan
//Computes an exclusive scan (partial reduction)
func Exscan(sendBuffer interface{}, recvBuffer interface{}, count int, datatype Datatype,
	op Op, comm Comm) int {

	sendBufferVoidPtr := Get_void_ptr(sendBuffer)
	recvBufferVoidPtr := Get_void_ptr(recvBuffer)

	err := C.MPI_Exscan(sendBufferVoidPtr,
		recvBufferVoidPtr,
		C.int(count),
		C.MPI_Datatype(datatype),
		C.MPI_Op(op),
		C.MPI_Comm(comm))

	return int(err)
}

//Get_address
//Gets the address of a location in memory.
func Get_address(location interface{}) (Aint, int) {

	locationVoidPtr := Get_void_ptr(location)
	var cAddress C.MPI_Aint

	err := C.MPI_Get_address(locationVoidPtr, &cAddress)

	return Aint(cAddress), int(err)
}

//Get_elements
//Returns the number of basic elements in a data type.
func Get_elements(status Status, datatype Datatype) (int, int) {

	cStatus := C.MPI_Status(status)
	var cCount C.int

	err := C.MPI_Get_elements(&cStatus, C.MPI_Datatype(datatype), &cCount)

	return int(cCount), int(err)
}

//Get
//Copies data from the target memory to the origin.
func Get(originAddr interface{}, originCount int, originDatatype Datatype,
	targetRank int, targetDisplacement Aint, targetCount int,
	targetDatatype Datatype, win Win) int {

	originAddrVoidPtr := Get_void_ptr(originAddr)

	err := C.MPI_Get(originAddrVoidPtr,
		C.int(originCount),
		C.MPI_Datatype(originDatatype),
		C.int(targetRank),
		C.MPI_Aint(targetDisplacement),
		C.int(targetCount),
		C.MPI_Datatype(targetDatatype),
		C.MPI_Win(win))

	return int(err)
}

//Get_processor_name
//Gets the name of the processor.
func Get_processor_name() (string, int) {

	var cstrLength C.int
	cName := make([]C.char, MAX_PROCESSOR_NAME)
	err := C.MPI_Get_processor_name(&cName[0], &cstrLength)

	name := C.GoStringN(&cName[0], cstrLength)

	return name, int(err)
}

//Get_version
//Returns the version of the standard corresponding to the current implementation.
func Get_() (int, int, int) {

	var cVersion C.int
	var cSubversion C.int

	err := C.MPI_Get_version(&cVersion, &cSubversion)

	return int(cVersion), int(cSubversion), int(err)
}

//Graph_create
//Makes a new communicator to which topology information has been attached.
func Graph_create(oldComm Comm, nnodes int, index []int, edges []int, reorder int) (Comm, int) {

	cSizeOfIndex := len(index)
	cSizeOfEdges := len(edges)

	cIndex := make([]C.int, cSizeOfIndex)
	cEdges := make([]C.int, cSizeOfEdges)

	for i := 0; i < cSizeOfIndex; i++ {
		cIndex[i] = C.int(index[i])
	}

	for i := 0; i < cSizeOfEdges; i++ {
		cEdges[i] = C.int(edges[i])
	}

	var cCommGraph C.MPI_Comm

	err := C.MPI_Graph_create(C.MPI_Comm(oldComm),
		C.int(nnodes),
		&cIndex[0],
		&cEdges[0],
		C.int(reorder),
		&cCommGraph)

	return Comm(cCommGraph), int(err)
}

//Graph_get
//Retrieves graph topology information associated with a communicator.
func Graph_get(comm Comm, maxIndex int, maxEdges int) ([]int, []int, int) {

	cIndex := make([]C.int, maxIndex)
	cEdges := make([]C.int, maxEdges)
	index := make([]int, maxIndex)
	edges := make([]int, maxEdges)

	err := C.MPI_Graph_get(C.MPI_Comm(comm),
		C.int(maxIndex),
		C.int(maxEdges),
		&cIndex[0],
		&cEdges[0])

	for i := 0; i < maxIndex; i++ {
		index[i] = int(cIndex[i])
	}

	for i := 0; i < maxEdges; i++ {
		edges[i] = int(cEdges[i])
	}

	return index, edges, int(err)
}

//Graph_map
//Maps process to graph topology information.
func Graph_map(comm Comm, nnodes int, index []int, edges []int) (int, int) {

	cSizeOfIndex := len(index)
	cSizeOfEdges := len(edges)

	cIndex := make([]C.int, cSizeOfIndex)
	cEdges := make([]C.int, cSizeOfEdges)

	for i := 0; i < cSizeOfIndex; i++ {
		cIndex[i] = C.int(index[i])
	}

	for i := 0; i < cSizeOfEdges; i++ {
		cEdges[i] = C.int(edges[i])
	}

	var cNewRank C.int

	err := C.MPI_Graph_map(C.MPI_Comm(comm),
		C.int(nnodes),
		&cIndex[0],
		&cEdges[0],
		&cNewRank)

	return int(cNewRank), int(err)
}

//Graph_neighbors_count
//Returns the number of neighbors of a node associated with a graph topology.
func Graph_neighbors_count(comm Comm, rank int) (int, int) {

	var cNNeigbors C.int

	err := C.MPI_Graph_neighbors_count(C.MPI_Comm(comm),
		C.int(rank),
		&cNNeigbors)

	return int(cNNeigbors), int(err)

}

//Graph_neighbors
//Returns the neighbors of a node associated with a graph topology.
func Graph_neighbors(comm Comm, rank int, maxNeighbors int) ([]int, int) {

	cNeighbors := make([]C.int, maxNeighbors)

	err := C.MPI_Graph_neighbors(C.MPI_Comm(comm),
		C.int(rank),
		C.int(maxNeighbors),
		&cNeighbors[0])

	neighbors := make([]int, maxNeighbors)
	for i := 0; i < maxNeighbors; i++ {
		neighbors[i] = int(cNeighbors[i])
	}

	return neighbors, int(err)
}

//Graphdims_get
//Retrieves graph topology information associated with a communicator.
func Graphdims_get(comm Comm) (int, int, int) {

	var cNNodes C.int
	var cNEdges C.int

	err := C.MPI_Graphdims_get(C.MPI_Comm(comm),
		&cNNodes,
		&cNEdges)

	return int(cNNodes), int(cNEdges), int(err)
}

//Grequest_complete
//Reports that a generalized request is complete.
func Grequest_completes(request Request) int {

	err := C.MPI_Grequest_complete(C.MPI_Request(request))

	return int(err)
}

//Intercomm_create
//Creates an intercommunicator from two intracommunicators.
func Intercomm_create(localComm Comm, localLeader int, peerComm Comm,
	remoteLeader int, tag int) (Comm, int) {

	var cNewInterComm C.MPI_Comm

	err := C.MPI_Intercomm_create(C.MPI_Comm(localComm),
		C.int(localLeader),
		C.MPI_Comm(peerComm),
		C.int(remoteLeader),
		C.int(tag),
		&cNewInterComm)

	return Comm(cNewInterComm), int(err)
}

//Intercomm_merge
//Creates an intracommunicator from an intercommunicator.
func Intercomm_merge(intercomm Comm, high int) (Comm, int) {

	var cNewInterComm C.MPI_Comm

	err := C.MPI_Intercomm_merge(C.MPI_Comm(intercomm),
		C.int(high),
		&cNewInterComm)

	return Comm(cNewInterComm), int(err)
}

//Irsend
//Starts a ready-mode nonblocking send.
func Irsend(buffer interface{}, count int, datatype Datatype, dst int, tag int, comm Comm) (Request, int) {

	bufferVoidPtr := Get_void_ptr(buffer)
	var cRequest C.MPI_Request

	err := C.MPI_Irsend(bufferVoidPtr,
		C.int(count),
		C.MPI_Datatype(datatype),
		C.int(dst),
		C.int(tag),
		C.MPI_Comm(comm),
		&cRequest)

	return Request(cRequest), int(err)
}

//Issend
//Starts a nonblocking synchronous send.
func Issend(buffer interface{}, count int, datatype Datatype, dst int, tag int,
	comm Comm) (Request, int) {

	bufferVoidPtr := Get_void_ptr(buffer)
	var cRequest C.MPI_Request

	err := C.MPI_Issend(bufferVoidPtr,
		C.int(count),
		C.MPI_Datatype(datatype),
		C.int(dst),
		C.int(tag),
		C.MPI_Comm(comm),
		&cRequest)

	return Request(cRequest), int(err)

}

//Is_thread_main
//Determines if thread called MPI_Init
func Is_thread_main() (bool, int) {

	var cFlag C.int
	err := C.MPI_Is_thread_main(&cFlag)

	flag := false
	if int(cFlag) != 0 {
		flag = true
	}

	return flag, int(err)
}

//Lookup_name
//Finds port associated with a service name
func Lookup_name(serviceName string, info Info) (string, int) {

	cstrServiceName := C.CString(serviceName)
	cPortname := (*C.char)(C.malloc(C.size_t(MAX_PORT_NAME)))

	err := C.MPI_Lookup_name(cstrServiceName,
		C.MPI_Info(info),
		cPortname)

	portName := C.GoString(cPortname)

	C.free(unsafe.Pointer(cstrServiceName))
	C.free(unsafe.Pointer(cPortname))

	return portName, int(err)
}

//Pack_external
//Writes data to a portable format
func Pack_external(datarep string, inBuffer interface{}, inCount int, datatype Datatype,
	outBuffer interface{}, outSize Aint, position Aint) (Aint, int) {

	inBufferVoidPtr := Get_void_ptr(inBuffer)
	outBufferVoidPtr := Get_void_ptr(outBuffer)

	cDatarep := C.CString(datarep)
	cPosition := C.MPI_Aint(position)

	err := C.MPI_Pack_external(cDatarep,
		inBufferVoidPtr,
		C.int(inCount),
		C.MPI_Datatype(datatype),
		outBufferVoidPtr,
		C.MPI_Aint(outSize),
		&cPosition)

	C.free(unsafe.Pointer(cDatarep))

	return Aint(cPosition), int(err)

}

//Pack_external_size
//Calculates upper bound on space needed to
func Pack_external_size(dataRep string, inCount int, datatype Datatype) (Aint, int) {

	cstrDataRep := C.CString(dataRep)
	var cSize C.MPI_Aint

	err := C.MPI_Pack_external_size(cstrDataRep,
		C.int(inCount),
		C.MPI_Datatype(datatype),
		&cSize)

	C.free(unsafe.Pointer(cstrDataRep))

	return Aint(cSize), int(err)
}

//Pack
//Packs data of a given datatype into contiguous memory.
func Pack(inBuffer interface{}, inCount int, datatype Datatype,
	outBuffer interface{}, outSize int, position int, comm Comm) (int, int) {

	inBufferVoidPtr := Get_void_ptr(inBuffer)
	outBufferVoidPtr := Get_void_ptr(outBuffer)
	cPosition := C.int(position)

	err := C.MPI_Pack(inBufferVoidPtr,
		C.int(inCount),
		C.MPI_Datatype(datatype),
		outBufferVoidPtr,
		C.int(outSize),
		&cPosition,
		C.MPI_Comm(comm))

	return int(cPosition), int(err)
}

//Pack_size
//Returns the upper bound on the amount of space needed to pack a message.
func Pack_size(inCount int, datatype Datatype, comm Comm) (int, int) {

	var cSize C.int

	err := C.MPI_Pack_size(C.int(inCount),
		C.MPI_Datatype(datatype),
		C.MPI_Comm(comm),
		&cSize)

	return int(cSize), int(err)
}

//Publish_name
//Publishes a service name associated with a port
func Publish_name(serviceName string, info Info, portName string) int {

	cServiceName := C.CString(serviceName)
	cPortName := C.CString(portName)
	err := C.MPI_Publish_name(cServiceName,
		C.MPI_Info(info),
		cPortName)

	C.free(unsafe.Pointer(cServiceName))
	C.free(unsafe.Pointer(cPortName))

	return int(err)
}

//Put
//Copies data from the origin memory to the target.
func Put(originAddr interface{}, originCount int, originType Datatype,
	targetRank int, targetDisplacement Aint, targetCount int,
	targetType Datatype, win Win) int {

	originAddrVoidPtr := Get_void_ptr(originAddr)

	err := C.MPI_Put(originAddrVoidPtr,
		C.int(originCount),
		C.MPI_Datatype(originType),
		C.int(targetRank),
		C.MPI_Aint(targetDisplacement),
		C.int(targetCount),
		C.MPI_Datatype(targetType),
		C.MPI_Win(win))

	return int(err)
}

//Testall
//Tests for the completion of all previously initiated communications in a list.
func Testall(requests []Request) ([]bool, []Status, int) {

	count := len(requests)
	cStatuses := make([]C.MPI_Status, count)
	cFlags := make([]C.int, count)
	cRequests := make([]C.MPI_Request, count)
	for i := 0; i < count; i++ {
		cRequests[i] = C.MPI_Request(requests[i])
	}

	err := C.MPI_Testall(C.int(count),
		&cRequests[0],
		&cFlags[0],
		&cStatuses[0])

	statuses := make([]Status, count)
	flags := make([]bool, count)
	for i := 0; i < count; i++ {
		statuses[i] = Status(cStatuses[i])
		if int(cFlags[i]) != 0 {
			flags[i] = true
		} else {
			flags[i] = false
		}
	}

	return flags, statuses, int(err)
}

//Testany
//Tests for completion of any one previously initiated communication in a list.
func Testany(requests []Request) ([]int, []bool, []Status, int) {

	count := len(requests)
	cStatuses := make([]C.MPI_Status, count)
	cFlags := make([]C.int, count)
	cIndex := make([]C.int, count)
	cRequests := make([]C.MPI_Request, count)
	for i := 0; i < count; i++ {
		cRequests[i] = C.MPI_Request(requests[i])
	}

	err := C.MPI_Testany(C.int(count),
		&cRequests[0],
		&cIndex[0],
		&cFlags[0],
		&cStatuses[0])

	statuses := make([]Status, count)
	flags := make([]bool, count)
	index := make([]int, count)
	for i := 0; i < count; i++ {
		statuses[i] = Status(cStatuses[i])
		index[i] = int(cIndex[i])
		if int(cFlags[i]) != 0 {
			flags[i] = true
		} else {
			flags[i] = false
		}
	}

	return index, flags, statuses, int(err)
}

//Testsome
//Tests for completion of one or more previously initiated communications in a list.
func Testsome(requests []Request) (int, []int, []Status, int) {

	count := len(requests)
	cStatuses := make([]C.MPI_Status, count)
	cIndex := make([]C.int, count)
	cRequests := make([]C.MPI_Request, count)
	for i := 0; i < count; i++ {
		cRequests[i] = C.MPI_Request(requests[i])
	}
	var cOutCount C.int

	err := C.MPI_Testsome(C.int(count),
		&cRequests[0],
		&cOutCount,
		&cIndex[0],
		&cStatuses[0])

	statuses := make([]Status, count)
	index := make([]int, count)
	for i := 0; i < count; i++ {
		statuses[i] = Status(cStatuses[i])
		index[i] = int(cIndex[i])
	}

	return int(cOutCount), index, statuses, int(err)
}

//Unpack
//Unpacks a datatype into contiguous memory.
func Unpack(inBuffer interface{}, inSize int, position int,
	outBuffer interface{}, outCount int, datatype Datatype, comm Comm) (int, int) {

	inBufferVoidPtr := Get_void_ptr(inBuffer)
	outBufferAddrVoidPtr := Get_void_ptr(outBuffer)
	cPoisiton := C.int(position)

	err := C.MPI_Unpack(inBufferVoidPtr,
		C.int(inSize),
		&cPoisiton,
		outBufferAddrVoidPtr,
		C.int(outCount),
		C.MPI_Datatype(datatype),
		C.MPI_Comm(comm))

	return int(cPoisiton), int(err)
}

//Unpublish_name
//Unpublishes a service name
func Unpublish_name(serviceName string, info Info, portName string) int {

	cServiceName := C.CString(serviceName)
	cPortName := C.CString(portName)

	err := C.MPI_Unpublish_name(cServiceName,
		C.MPI_Info(info),
		cPortName)

	C.free(unsafe.Pointer(cServiceName))
	C.free(unsafe.Pointer(cPortName))

	return int(err)
}

//Unpack_external
//Reads data from a portable format.
func Unpack_external(datarep string, inBuffer interface{}, inSize Aint, position Aint,
	outBuffer interface{}, outCount int, datatype Datatype) (Aint, int) {

	inBufferVoidPtr := Get_void_ptr(inBuffer)
	outBufferAddrVoidPtr := Get_void_ptr(outBuffer)
	cstrDatarep := C.CString(datarep)
	cPosition := C.MPI_Aint(position)

	err := C.MPI_Unpack_external(cstrDatarep,
		inBufferVoidPtr,
		C.MPI_Aint(inSize),
		&cPosition,
		outBufferAddrVoidPtr,
		C.int(outCount),
		C.MPI_Datatype(datatype))

	C.free(unsafe.Pointer(cstrDatarep))

	return Aint(cPosition), int(err)

}

//Ibsend
//Starts a nonblocking buffered send.
func Ibsend(buffer interface{}, count int, datatype Datatype, dst int, tag int,
	comm Comm) (Request, int) {

	bufferVoidPtr := Get_void_ptr(buffer)
	var cRequest C.MPI_Request

	err := C.MPI_Ibsend(bufferVoidPtr,
		C.int(count),
		C.MPI_Datatype(datatype),
		C.int(dst),
		C.int(tag),
		C.MPI_Comm(comm),
		&cRequest)

	return Request(cRequest), int(err)
}

//Pcontrol
//Controls profiling.
func Pcontrol(level int) int {

	err := C.MPI_Pcontrol_(C.int(1))

	return int(err)
}
