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
#include <stdio.h>
#include <mpi.h>
void dbg( char* c ){
	printf("%s\n", c);
}

*/
import "C"

import (
	"fmt"
	"unsafe"
)

// this is a macro - probs
// type MPI_File_errhandler_function unsafe.Pointer

const MAX_BUFFER_SIZE = 255

//File_call_errhandler
//Passes the supplied error code to the error handler assigned to a file
func File_call_errhandler(fh File, errorcode int) int {

	err := C.MPI_File_call_errhandler(C.MPI_File(fh), C.int(errorcode))

	return int(err)
}

// func File_create_errhandler(function MPI_File_errhandler_function) (MPI_Errhandler, int) {
// 	var errhandler C.MPI_Errhandler
// 	err := C.MPI_File_create_errhandler(C.MPI_File_errhandler_function(function), &errhandler)
// 	return MPI_Errhandler(errhandler), int(err)
// }

//File_set_errhandler
//Set the error handler for an MPI file.
func File_set_errhandler(fh File, errhandler Errhandler) int {

	err := C.MPI_File_set_errhandler(C.MPI_File(fh), C.MPI_Errhandler(errhandler))

	return int(err)
}

//File_get_errhandler
//Get the error handler attached to a file.
func File_get_errhandler(fh File) (Errhandler, int) {

	var errhandler C.MPI_Errhandler
	err := C.MPI_File_get_errhandler(C.MPI_File(fh), &errhandler)

	return Errhandler(errhandler), int(err)
}

//File_open
//Opens a file (collective).
func File_open(comm Comm, filename string, amode int, info Info) (*File, int) {

	var fh C.MPI_File
	// fh := (*C.MPI_File)(C.malloc(C.sizeof(C.MPI_File)))
	name := C.CString(filename)
	// C.dbg(name)

	err := C.MPI_File_open(
		C.MPI_Comm(comm),
		name,
		C.int(amode),
		C.MPI_Info(info),
		&fh)

	fmt.Println(err)

	C.free((unsafe.Pointer)(name))

	return (*File)(unsafe.Pointer(&fh)), int(err)
}

//File_close
//Closes a file (collective).
//TODO: refactoring - in parameter needs to be a pointer
func File_close(fh *File) int {

	err := C.MPI_File_close((*C.MPI_File)(unsafe.Pointer(fh)))

	return int(err)
}

//File_delete
//Deletes a file.
func File_delete(filename string, info Info) int {

	err := C.MPI_File_delete(C.CString(filename), C.MPI_Info(info))

	return int(err)
}

//File_set_size
//Resizes a file (collective).
func File_set_size(fh File, size Offset) int {

	err := C.MPI_File_set_size(C.MPI_File(fh), C.MPI_Offset(size))

	return int(err)
}

//File_preallocate
//Preallocates a specified amount of storage space at the beginning of a file (collective).
func File_preallocate(fh File, size Offset) int {

	err := C.MPI_File_preallocate(C.MPI_File(fh), C.MPI_Offset(size))

	return int(err)
}

//File_get_size
//Returns the current size of the file.
func File_get_size(fh File) (Offset, int) {

	var size C.MPI_Offset
	err := C.MPI_File_get_size(C.MPI_File(fh), &size)

	return Offset(size), int(err)
}

//File_get_group
//Returns a duplicate of the process group of a file.
func File_get_group(fh File) (Group, int) {

	var group C.MPI_Group
	err := C.MPI_File_get_group(C.MPI_File(fh), &group)

	return Group(group), int(err)
}

//File_get_amode
//Returns access mode associated with an open file.
func File_get_amode(fh File) (int, int) {

	var amode C.int
	err := C.MPI_File_get_amode(C.MPI_File(fh), &amode)

	return int(amode), int(err)
}

//File_set_info
//Sets new values for hints (collective).
func File_set_info(fh File, info Info) (File, int) {

	new_fh := C.MPI_File(fh)
	err := C.MPI_File_set_info(new_fh, C.MPI_Info(info))

	return File(new_fh), int(err)
}

//File_get_info
//Returns a new info object containing values for
//current hints associated with a file.
func File_get_info(fh File) (Info, int) {

	var info_used C.MPI_Info
	err := C.MPI_File_get_info(C.MPI_File(fh), &info_used)

	return Info(info_used), int(err)

}

//File_set_view
//Changes process’s view of data in file (collective).
func File_set_view(fh File, disp Offset, etype Datatype, filetype Datatype, datarep string, info Info) (File, int) {

	new_fh := C.MPI_File(fh)
	err := C.MPI_File_set_view(fh,
		C.MPI_Offset(disp),
		C.MPI_Datatype(etype),
		C.MPI_Datatype(filetype),
		C.CString(datarep),
		C.MPI_Info(info))

	return File(new_fh), int(err)
}

//File_get_view
//Returns the process’s view of data in the file.
func File_get_view(fh File) (Offset, Datatype, Datatype, string, int) {

	var disp C.MPI_Offset
	var etype C.MPI_Datatype
	var filetype C.MPI_Datatype
	datarep := (*C.char)(C.malloc(MAX_BUFFER_SIZE))

	err := C.MPI_File_get_view(
		C.MPI_File(fh),
		&disp,
		&etype,
		&filetype,
		datarep)

	return Offset(disp), Datatype(etype), Datatype(filetype), C.GoString(datarep), int(err)

}

//File_iread
//Reads a file starting at the location specified by the individual
//file pointer (nonblocking, noncollective).
func File_iread(fh File, count int, datatype Datatype) (File, unsafe.Pointer, Request, int) {

	new_fh := C.MPI_File(fh)
	var buffer unsafe.Pointer
	var request C.MPI_Request

	err := C.MPI_File_iread(new_fh,
		buffer,
		C.int(count),
		C.MPI_Datatype(datatype),
		&request)

	return File(new_fh), buffer, Request(request), int(err)
}

//File_iwrite
//Writes a file starting at the location specified by the individual file pointer (nonblocking, noncollective).
func File_iwrite(fh File, buffer unsafe.Pointer, count int, datatype Datatype) (File, Request, int) {

	new_fh := C.MPI_File(fh)
	var request C.MPI_Request

	err := C.MPI_File_iwrite(new_fh,
		buffer,
		C.int(count),
		C.MPI_Datatype(datatype),
		&request)

	return File(new_fh), Request(request), int(err)
}

//File_seek
//Updates individual file pointers (noncollective).
func File_seek(fh File, offset Offset, whence int) int {

	err := C.MPI_File_seek(C.MPI_File(fh),
		C.MPI_Offset(offset),
		C.int(whence))

	return int(err)
}

//File_seek_shared
//Updates the global shared file pointer (collective).
func File_seek_shared(fh File, offset Offset, whence int) int {

	err := C.MPI_File_seek_shared(C.MPI_File(fh),
		C.MPI_Offset(offset),
		C.int(whence))

	return int(err)
}

//File_get_position
//Returns the current position of the individual file pointer.
func File_get_position(fh File) (Offset, int) {

	var offset C.MPI_Offset

	err := C.MPI_File_get_position(C.MPI_File(fh), &offset)

	return Offset(offset), int(err)
}

//File_get_position_shared
//Returns the current position of the shared file pointer.
func File_get_position_shared(fh File) (Offset, int) {

	var offset C.MPI_Offset

	err := C.MPI_File_get_position_shared(C.MPI_File(fh), &offset)

	return Offset(offset), int(err)
}

//File_get_byte_offset
//Converts a view-relative offset into an absolute byte position.
func File_get_byte_offset(fh File, offset Offset) (Offset, int) {

	var disp C.MPI_Offset

	err := C.MPI_File_get_byte_offset(C.MPI_File(fh),
		C.MPI_Offset(offset),
		&disp)

	return Offset(disp), int(err)
}

//File_read_at
//Reads a file at an explicitly specified offset (blocking, noncollective).
func File_read_at(fh File, offset Offset, count int, datatype Datatype) (unsafe.Pointer, Status, int) {

	var buffer unsafe.Pointer
	var status C.MPI_Status

	err := C.MPI_File_read_at(C.MPI_File(fh),
		C.MPI_Offset(offset),
		buffer,
		C.int(count),
		C.MPI_Datatype(datatype),
		&status)

	return buffer, Status(status), int(err)
}

//File_read_at_all
//Reads a file at explicitly specified offsets (blocking, collective).
func File_read_at_all(fh File, offset Offset, count int, datatype Datatype) (unsafe.Pointer, Status, int) {

	var buffer unsafe.Pointer
	var status C.MPI_Status

	err := C.MPI_File_read_at_all(C.MPI_File(fh),
		C.MPI_Offset(offset),
		buffer,
		C.int(count),
		C.MPI_Datatype(datatype),
		&status)

	return buffer, Status(status), int(err)
}

//File_iread_at
//Reads a file at an explicitly specified offset (nonblocking, noncollective).
func File_iread_at(fh File, offset Offset, count int, datatype Datatype) (unsafe.Pointer, Request, int) {

	var buffer unsafe.Pointer
	var request C.MPI_Request

	err := C.MPI_File_iread_at(C.MPI_File(fh),
		C.MPI_Offset(offset),
		buffer,
		C.int(count),
		C.MPI_Datatype(datatype),
		&request)

	return buffer, Request(request), int(err)
}

//File_read
//Reads a file starting at the location specified by the individual file
//pointer (blocking, noncollective).
func File_read(fh File, count int, datatype Datatype) (unsafe.Pointer, Status, int) {

	var buffer unsafe.Pointer
	var status C.MPI_Status

	err := C.MPI_File_read(C.MPI_File(fh),
		buffer,
		C.int(count),
		C.MPI_Datatype(datatype),
		&status)

	return buffer, Status(status), int(err)
}

//File_read_all
//Reads a file starting at the locations specified by individual file pointers (blocking, collective).
func File_read_all(fh File, count int, datatype Datatype) (unsafe.Pointer, Status, int) {

	var buffer unsafe.Pointer
	var status C.MPI_Status

	err := C.MPI_File_read_all(C.MPI_File(fh),
		buffer,
		C.int(count),
		C.MPI_Datatype(datatype),
		&status)

	return buffer, Status(status), int(err)
}

//File_read_shared
//Reads a file using the shared file pointer (blocking, noncollective).
func File_read_shared(fh File, count int, datatype Datatype) (File, unsafe.Pointer, Status, int) {

	new_fh := C.MPI_File(fh)
	var buffer unsafe.Pointer
	var status C.MPI_Status

	err := C.MPI_File_read_shared(new_fh,
		buffer,
		C.int(count),
		C.MPI_Datatype(datatype),
		&status)

	return File(new_fh), buffer, Status(status), int(err)
}

//File_iread_shared
//Reads a file using the shared file pointer (nonblocking, noncollective).
func File_iread_shared(fh File, count int, datatype Datatype) (File, unsafe.Pointer, Request, int) {

	new_fh := C.MPI_File(fh)
	var buffer unsafe.Pointer
	var request C.MPI_Request

	err := C.MPI_File_iread_shared(new_fh,
		buffer,
		C.int(count),
		C.MPI_Datatype(datatype),
		&request)

	return File(new_fh), buffer, Request(request), int(err)
}

//File_read_ordered
//Reads a file at a location specified by a shared file pointer (blocking, collective).
func File_read_ordered(fh File, count int, datatype Datatype) (unsafe.Pointer, Status, int) {

	var buffer unsafe.Pointer
	var status C.MPI_Status

	err := C.MPI_File_read_ordered(C.MPI_File(fh),
		buffer,
		C.int(count),
		C.MPI_Datatype(datatype),
		&status)

	return buffer, Status(status), int(err)
}

//File_read_at_all_begin
//Reads a file at explicitly specified offsets;
//beginning part of a split collective routine (nonblocking).
func File_read_at_all_begin(fh File, offset Offset, count int, datatype Datatype) (unsafe.Pointer, int) {

	var buffer unsafe.Pointer

	err := C.MPI_File_read_at_all_begin(C.MPI_File(fh),
		C.MPI_Offset(offset),
		buffer,
		C.int(count),
		C.MPI_Datatype(datatype))

	return buffer, int(err)
}

//File_read_at_all_end
//Reads a file at explicitly specified offsets;
//ending part of a split collective routine (blocking).
func File_read_at_all_end(fh File) (File, unsafe.Pointer, Status, int) {

	new_fh := C.MPI_File(fh)
	var buf unsafe.Pointer
	var status C.MPI_Status

	err := C.MPI_File_read_at_all_end(new_fh,
		buf,
		&status)

	return File(new_fh), buf, Status(status), int(err)
}

//File_read_all_begin
//Reads a file starting at the locations specified by individual
//file pointers; beginning part of a split collective routine (nonblocking).
func File_read_all_begin(fh File, count int, datatype Datatype) (File, unsafe.Pointer, int) {

	new_fh := C.MPI_File(fh)
	var buf unsafe.Pointer

	err := C.MPI_File_read_all_begin(new_fh,
		buf,
		C.int(count),
		C.MPI_Datatype(datatype))

	return File(new_fh), buf, int(err)
}

//File_read_all_end
//Reads a file starting at the locations specified by individual
//file pointers; ending part of a split collective routine (blocking).
func File_read_all_end(fh File) (File, unsafe.Pointer, Status, int) {

	new_fh := C.MPI_File(fh)
	var buf unsafe.Pointer
	var status C.MPI_Status

	err := C.MPI_File_read_all_end(new_fh,
		buf,
		&status)

	return File(new_fh), buf, Status(status), int(err)
}

//File_read_ordered_begin
//Reads a file at a location specified by a shared file pointer;
//beginning part of a split collective routine (nonblocking).
func File_read_ordered_begin(fh File, count int, datatype Datatype) (File, unsafe.Pointer, int) {

	new_fh := C.MPI_File(fh)
	var buf unsafe.Pointer

	err := C.MPI_File_read_ordered_begin(new_fh,
		buf,
		C.int(count),
		C.MPI_Datatype(datatype))

	return File(new_fh), buf, int(err)
}

//File_read_ordered_end
//Reads a file at a location specified by a shared file pointer;
//ending part of a split collective routine (blocking).
func File_read_ordered_end(fh File) (File, unsafe.Pointer, Status, int) {

	new_fh := C.MPI_File(fh)
	var buf unsafe.Pointer
	var status C.MPI_Status

	err := C.MPI_File_read_ordered_end(new_fh,
		buf,
		&status)

	return File(new_fh), buf, Status(status), int(err)
}

//File_write
//Writes a file starting at the location specified by the individual file pointer (blocking, noncollective).
func File_write(fh File, buffer unsafe.Pointer, count int, datatype Datatype) (File, Status, int) {

	new_fh := C.MPI_File(fh)
	var status C.MPI_Status

	err := C.MPI_File_write(new_fh,
		buffer,
		C.int(count),
		C.MPI_Datatype(datatype),
		&status)

	return File(new_fh), Status(status), int(err)
}

//File_write_all
//Writes a file starting at the locations specified by individual file pointers (blocking, collective).
func File_write_all(fh File, buffer unsafe.Pointer, count int, datatype Datatype) (File, Status, int) {

	new_fh := C.MPI_File(fh)
	var status C.MPI_Status

	err := C.MPI_File_write_all(new_fh,
		buffer,
		C.int(count),
		C.MPI_Datatype(datatype),
		&status)

	return File(new_fh), Status(status), int(err)
}

//File_write_shared
//Writes a file using the shared file pointer (blocking, noncollective).
func File_write_shared(fh File, buffer unsafe.Pointer, count int, datatype Datatype) (File, Status, int) {

	new_fh := C.MPI_File(fh)
	var status C.MPI_Status

	err := C.MPI_File_write_shared(new_fh,
		buffer,
		C.int(count),
		C.MPI_Datatype(datatype),
		&status)

	return File(new_fh), Status(status), int(err)
}

//File_write_at
//Writes a file at an explicitly specified offset (blocking, noncollective).
func File_write_at(fh File, offset Offset, buffer unsafe.Pointer, count int, datatype Datatype) (Status, int) {

	var status C.MPI_Status

	err := C.MPI_File_write_at(C.MPI_File(fh),
		C.MPI_Offset(offset),
		buffer,
		C.int(count),
		C.MPI_Datatype(datatype),
		&status)

	return Status(status), int(err)
}

//File_write_at_all
//Writes a file at explicitly specified offsets (blocking, collective).
func File_write_at_all(fh File, offset Offset, buffer unsafe.Pointer, count int, datatype Datatype) (Status, int) {

	var status C.MPI_Status

	err := C.MPI_File_write_at_all(C.MPI_File(fh),
		C.MPI_Offset(offset),
		buffer,
		C.int(count),
		C.MPI_Datatype(datatype),
		&status)

	return Status(status), int(err)
}

//File_iwrite_at
//Writes a file at an explicitly specified offset (nonblocking, noncollective).
func File_iwrite_at(fh File, offset Offset, buffer unsafe.Pointer, count int, datatype Datatype) (Request, int) {

	var request C.MPI_Request

	err := C.MPI_File_iwrite_at(C.MPI_File(fh),
		C.MPI_Offset(offset),
		buffer,
		C.int(count),
		C.MPI_Datatype(datatype),
		&request)

	return Request(request), int(err)
}

//File_iwrite_shared
//Writes a file using the shared file pointer (nonblocking, noncollective).
func File_iwrite_shared(fh File, buffer unsafe.Pointer, count int, datatype Datatype) (File, Request, int) {

	new_fh := C.MPI_File(fh)
	var request C.MPI_Request

	err := C.MPI_File_iwrite_shared(new_fh,
		buffer,
		C.int(count),
		C.MPI_Datatype(datatype),
		&request)

	return File(new_fh), Request(request), int(err)
}

//File_write_ordered
//Writes a file at a location specified by a shared file pointer (blocking, collective).
func File_write_ordered(fh File, buffer unsafe.Pointer, count int, datatype Datatype) (Status, int) {

	var status C.MPI_Status

	err := C.MPI_File_write_ordered(C.MPI_File(fh),
		buffer,
		C.int(count),
		C.MPI_Datatype(datatype),
		&status)

	return Status(status), int(err)
}

//File_write_at_all_begin
//Writes a file at explicitly specified offsets
//beginning part of a split collective routine (nonblocking).
func File_write_at_all_begin(fh File, offset Offset, buf unsafe.Pointer, count int, datatype Datatype) (File, int) {

	new_fh := C.MPI_File(fh)

	err := C.MPI_File_write_at_all_begin(new_fh,
		C.MPI_Offset(offset),
		buf,
		C.int(count),
		C.MPI_Datatype(datatype))

	return File(new_fh), int(err)
}

//File_write_at_all_end
//Writes a file at explicitly specified offsets; ending part of a split collective routine (blocking).
func File_write_at_all_end(fh File, buf unsafe.Pointer) (File, Status, int) {

	new_fh := C.MPI_File(fh)
	var status C.MPI_Status

	err := C.MPI_File_write_at_all_end(new_fh,
		buf,
		&status)

	return File(new_fh), Status(status), int(err)
}

//File_write_all_begin
//Writes a file starting at the locations specified by individual file pointers;
//beginning part of a split collective routine (nonblocking).
func File_write_all_begin(fh File, buf unsafe.Pointer, count int, datatype Datatype) (File, int) {

	new_fh := C.MPI_File(fh)

	err := C.MPI_File_write_all_begin(new_fh,
		buf,
		C.int(count),
		C.MPI_Datatype(datatype))

	return File(new_fh), int(err)
}

//File_write_all_end
//Writes a file starting at the locations specified by individual file pointers;
//ending part of a split collective routine (blocking).
func File_write_all_end(fh File, buf unsafe.Pointer) (File, Status, int) {

	new_fh := C.MPI_File(fh)
	var status C.MPI_Status

	err := C.MPI_File_write_all_end(new_fh,
		buf,
		&status)

	return File(new_fh), Status(status), int(err)
}

//File_write_ordered_begin
//Writes a file at a location specified by a shared file pointer;
//beginning part of a split collective routine (nonblocking).
func File_write_ordered_begin(fh File, buf unsafe.Pointer, count int, datatype Datatype) (File, int) {

	new_fh := C.MPI_File(fh)

	err := C.MPI_File_write_ordered_begin(new_fh,
		buf,
		C.int(count),
		C.MPI_Datatype(datatype))

	return File(new_fh), int(err)
}

//File_write_ordered_end
//Writes a file at a location specified by a shared file pointer;
//ending part of a split collective routine (blocking).
func File_write_ordered_end(fh File) (File, unsafe.Pointer, Status, int) {

	new_fh := C.MPI_File(fh)
	var buf unsafe.Pointer
	var status C.MPI_Status

	err := C.MPI_File_write_ordered_end(new_fh,
		buf,
		&status)

	return File(new_fh), buf, Status(status), int(err)
}

//File_get_type_extent
//Returns the extent of the data type in a file.
func File_get_type_extent(fh File, datatype Datatype) (Aint, int) {

	var extend C.MPI_Aint

	err := C.MPI_File_get_type_extent(C.MPI_File(fh),
		C.MPI_Datatype(datatype),
		&extend)

	return Aint(extend), int(err)
}

//File_set_atomicity
//Sets consistency semantics for data-access operations (collective).
func File_set_atomicity(fh File, flag int) int {

	return int(C.MPI_File_set_atomicity(C.MPI_File(fh),
		C.int(flag)))
}

//File_get_atomicity
//Returns current consistency semantics for data-access operations.
func File_get_atomicity(fh File) (int, int) {

	var flag C.int

	err := C.MPI_File_get_atomicity(C.MPI_File(fh),
		&flag)

	return int(flag), int(err)
}

//File_sync
//Makes semantics consistent for data-access operations (collective).
func File_sync(fh File) int {
	return int(C.MPI_File_sync(C.MPI_File(fh)))
}
