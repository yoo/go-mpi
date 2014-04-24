package MPI

import (
	"fmt"
	"os"
	"testing"
)

func TestType_create_gostruct(T *testing.T) {
	err := Init(&os.Args)

	if err != SUCCESS {
		T.Error("MPI.Init failed.\n")
	}

	rank, err := Comm_rank(COMM_WORLD)
	if err != SUCCESS {
		T.Error("MPI.Comm_rank failed.\n")
		Abort(COMM_WORLD, err)
	}

	type embed struct {
		f1 int
	}

	type testStruct struct {
		f1 int
		f2 float64
		f3 [10]byte
		f4 embed
	}

	var ts testStruct

	if rank == 0 {
		ts.f1 = 1
		ts.f2 = 2.0
		ts.f3 = [10]byte{'f', 'o', 'o'}
		ts.f4.f1 = 1
	}

	MPI_testStruct, err := Type_create_gostruct(ts)
	Type_commit(&MPI_testStruct)
	if err != SUCCESS {
		T.Error("Failed to call Type_create_gostruct.")
		Abort(COMM_WORLD, err)
	}

	if rank == 0 {
		Send(&ts, 1, MPI_testStruct, 1, 1, COMM_WORLD)
	} else {
		Recv(&ts, 1, MPI_testStruct, 0, 1, COMM_WORLD)
	}

	if rank != 0 {
		var tsValid testStruct

		tsValid.f1 = 1
		tsValid.f2 = 2.0
		tsValid.f3 = [10]byte{'f', 'o', 'o'}
		tsValid.f4.f1 = 1

		tsString := fmt.Sprintf("%#v", ts)
		tsValidString := fmt.Sprintf("%#v", tsValid)

		if tsString != tsValidString {
			T.Errorf("Expected: %s\n,was: %s\n", tsValidString, tsString)
		}
	}
}
