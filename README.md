# go-mpi

go-mpi are GO bindings for the Message Passing Interface <a href=http://www.mpi-forum.org/>(MPI Forum)</a>.

## Installation

MPI is a standard but the different implementations differ in some details.
At the moment go-mpi support  <a href=http://www.open-mpi.de/>Open MPI</a> and <a href=http://www.mpich.org/>MPICH</a> version 2.

To tell go where to look for mpi.h and mpi library, use the CGO_CFLAGS and CGO_LDFALG environment variable to indicate respectively. You can find paths of mpi.h and mpi library through command:

	$ mpichversion


Assume mpi.h dir is "/usr/local/include", mpi library dir is "/usr/local/lib", you could compile go-mpi/MPI like the following:

	export CGO_CFLAGS='-I/usr/local/include'
	export CGO_LDFLAGS='-L/usr/local/lib -lmpich'
	go install go-mpi/MPI


## Syntax

Firstly you should import "go-mpi/MPI" in your go-mpi application.

	import "go-mpi/MPI"


The syntax of MPI invokes in go-mpi is similar to it in C-binding MPI implementations.

	<package_name>.Mpi_function(arguments)

For example:

	err = MPI.Init(&os.Args)
	err = MPI.Irecv(recvBuffer, count, MPI.INT, 0, 1, MPI.COMM_WORLD, &request)
