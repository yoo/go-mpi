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
	"fmt"
)

// Jo
func TestErrorcode(functionName string, functionMap map[string]string, errorCode int) {

	fmt.Println("-----")
	if errorCode == SUCCESS {
		fmt.Println(functionName, ": returns MPI.SUCCESS -> PASS")
		functionMap[functionName] = "PASS"
	} else {
		fmt.Println(functionName, ": returns MPI.SUCCESS -> ERROR")
		functionMap[functionName] = "ERROR"
	}
}

// Jo
func TestAssert(functionName string, errorCode int, functionMap map[string]string, assert bool) {

	TestErrorcode(functionName, functionMap, errorCode)

	if assert {
		fmt.Println(functionName, ": assert -> PASS")
		functionMap[functionName] = "PASS"
	} else {
		fmt.Println(functionName, ": assert -> ERROR")
		functionMap[functionName] = "ERROR"
	}
}

// Jo
func TestBuffer(functionName string, errorCode int, functionMap map[string]string, buffer [][]int, validateBuffer [][]int) {

	TestErrorcode(functionName, functionMap, errorCode)

	strSlice1 := fmt.Sprintf("%v", buffer)
	strSlice2 := fmt.Sprintf("%v", validateBuffer)
	if strSlice1 == strSlice2 {
		fmt.Println(functionName, ": compare buffers -> PASS")
	} else {
		fmt.Println(functionName, ": compare buffers -> ERROR")
		functionMap[functionName] = "ERROR"
	}

}
