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

//MPI_Win_call_errhandler
//Passes the supplied error code to the error handler assigned to a window.
func Win_call_errhandler(win Win, errorcode int) int {

	err := C.MPI_Win_call_errhandler(C.MPI_Win(win), C.int(errorcode))

	return int(err)
}

//MPI_Win_complete
//Completes an RMA access epoch on win started by a call to MPI_Win_start
func Win_complete(win Win) int {

	err := C.MPI_Win_complete(C.MPI_Win(win))

	return int(err)
}

//Win_create
//One-sided MPI call that returns a window object for RMA operations.
func Win_create(base interface{}, size Aint, displacementUnit int, info Info, comm Comm) (Win, int) {

	baseVoidPtr := Get_void_ptr(base)

	var cWin C.MPI_Win

	err := C.MPI_Win_create(baseVoidPtr,
		C.MPI_Aint(size),
		C.int(displacementUnit),
		C.MPI_Info(info),
		C.MPI_Comm(comm),
		&cWin)

	return Win(cWin), int(err)
}

//MPI_Win_delete_attr
//Deletes an attribute from a window.
func Win_delete_attr(win Win, winKeyval int) (Win, int) {

	cWin := C.MPI_Win(win)
	err := C.MPI_Win_delete_attr(cWin, C.int(winKeyval))

	return Win(cWin), int(err)
}

//Win_fence
//Synchronizes RMA calls on a window.
func Win_fence(assert bool, win Win) int {

	cAssert := C.int(0)

	if assert {
		cAssert = 1
	}

	err := C.MPI_Win_fence(cAssert, C.MPI_Win(win))

	return int(err)
}

//Win_free
//Frees the window object and returns a null handle.
func Win_free(win Win) (Win, int) {

	cWin := C.MPI_Win(win)
	err := C.MPI_Win_free(&cWin)

	return Win(cWin), int(err)
}

//Win_free_keyval
//Frees a window keyval.
func Win_free_keyval(keyval int) (int, int) {

	cKeyval := C.int(keyval)
	err := C.MPI_Win_free_keyval(&cKeyval)

	return int(cKeyval), int(err)
}

//Win_get_attr
//Obtains the value of a window attribute.
func Win_get_attr(win Win, winKeyval int) (uintptr, bool, int) {

	var attributeVal uintptr
	var cFlag C.int
	flag := false
	err := C.MPI_Win_get_attr(C.MPI_Win(win),
		C.int(winKeyval),
		unsafe.Pointer(attributeVal),
		&cFlag)

	if int(cFlag) != 0 {
		flag = true
	}

	return attributeVal, flag, int(err)
}

//Win_get_errhandler
//Retrieves the error handler currently associated with a window.
func Win_get_errhandler(win Win) (Errhandler, int) {

	var cErrhandler C.MPI_Errhandler
	err := C.MPI_Win_get_errhandler(C.MPI_Win(win), &cErrhandler)

	return Errhandler(cErrhandler), int(err)
}

//Win_get_group
//Returns a duplicate of the group of the communicator used to create the window.
func Win_get_group(win Win) (Group, int) {

	var cGroup C.MPI_Group
	err := C.MPI_Win_get_group(C.MPI_Win(win), &cGroup)

	return Group(cGroup), int(err)
}

//MPI_Win_get_name
//Obtains the name of a window.
func Win_get_name(win Win) (string, int) {

	cstrWinName := (*C.char)(C.malloc(C.size_t(MAX_OBJECT_NAME)))
	cResultLength := C.int(0)

	err := C.MPI_Win_get_name(C.MPI_Win(win),
		cstrWinName,
		&cResultLength)

	winName := C.GoStringN(cstrWinName, cResultLength)
	C.free(unsafe.Pointer(cstrWinName))

	return winName, int(err)
}

//Win_lock
//Starts an RMA access epoch locking access to a particular rank.
func Win_lock(lockTpe int, rank int, assert bool, win Win) int {

	cAssert := C.int(0)
	if assert {
		cAssert = C.int(1)
	}

	err := C.MPI_Win_lock(C.int(lockTpe),
		C.int(rank),
		cAssert,
		C.MPI_Win(win))

	return int(err)
}

//Win_post
//Starts an RMA exposure epoch for the local window associated with win.
func Win_post(group Group, assert bool, win Win) int {

	cAssert := C.int(0)
	if assert {
		cAssert = C.int(1)
	}

	err := C.MPI_Win_post(C.MPI_Group(group),
		cAssert,
		C.MPI_Win(win))

	return int(err)
}

//Win_set_attr
//Sets the value of a window attribute.
func Win_set_attr(win Win, winKeyval int, attributeVal interface{}) int {

	attributeValVoidPtr := Get_void_ptr(attributeVal)

	err := C.MPI_Win_set_attr(C.MPI_Win(win),
		C.int(winKeyval),
		attributeValVoidPtr)

	return int(err)
}

//Win_set_errhandler
//Attaches a new error handler to a window.
func Win_set_errhandler(win Win, errhandler Errhandler) int {

	err := C.MPI_Win_set_errhandler(C.MPI_Win(win), C.MPI_Errhandler(errhandler))

	return int(err)
}

//Win_set_name
//Sets the name of a window.
func Win_set_name(win Win, winName string) int {

	cstrWinName := C.CString(winName)
	err := C.MPI_Win_set_name(C.MPI_Win(win), cstrWinName)

	C.free(unsafe.Pointer(cstrWinName))

	return int(err)
}

//Win_start
//Starts an RMA access epoch for win.
func Win_start(group Group, assert bool, win Win) int {

	cAssert := C.int(0)
	if assert {
		cAssert = C.int(1)
	}

	err := C.MPI_Win_start(C.MPI_Group(group), cAssert, C.MPI_Win(win))

	return int(err)
}

//MPI_Win_test
//Attempts to complete an RMA exposure epoch; a nonblocking version of MPI_Win_wait.
func Win_test(win Win) (bool, int) {

	var cFlag C.int
	flag := false
	err := C.MPI_Win_test(C.MPI_Win(win), &cFlag)

	if int(cFlag) != 0 {
		flag = true
	}

	return flag, int(err)
}

//Win_unlock
//Completes an RMA access epoch started by a call to MPI_Win_lock.
func Win_unlock(rank int, win Win) int {

	err := C.MPI_Win_unlock(C.int(rank), C.MPI_Win(win))

	return int(err)
}

//Win_wait
//Completes an RMA exposure epoch started by a call to MPI_Win_post on win.
func Win_wait(win Win) int {

	err := C.MPI_Win_wait(C.MPI_Win(win))

	return int(err)
}
