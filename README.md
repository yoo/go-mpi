# go-mpi

go-mpi are GO bindings for the Message Passing Interface <a href=http://www.mpi-forum.org/>(MPI Forum)</a>.

## Installation

MPI is a standard but the different implementations differ in some details.
At the moment go-mpi support  <a href=http://www.open-mpi.de/>Open MPI</a> and <a href=http://www.mpich.org/>MPICH</a> version 2.

To tell go where to look for mpi.h and mpi library, use the CGO_CFLAGS and CGO_LDFALG environment variable to indicate respectively. 

Assume your mpi.h dir is "/usr/local/include", mpi library dir is "/usr/local/lib" and you use mpich, you could compile go-mpi/MPI like the following:
```sh
export CGO_CFLGAS='-I/usr/local/include'
export CGO_LDFLAGS='-L/usr/local/lib -lmpich'
go install go-mpi/MPI
```

## Syntax

Firstly you should import "go-mpi/MPI" in your go-mpi application.
```
import "go-mpi/MPI"
```

The syntax of MPI invokes in go-mpi is similar to it in C-binding MPI implementations.
```
<package_name>.Mpi_function(arguments)
```

If the bindings are imported as "MPI":
<pre>
  err = MPI.Init(&os.Args)
</pre>

Output parameter like request objects are returned by the function and not passed as pointers inside the arguments.

<pre>
  C:
  err = MPI_Irecv(recvBuffer, count, MPI_INT, 0, 1, MPI_COMM_WROLD, &request)

  GO:
  err = MPI.Irecv(recvBuffer, count, MPI.INT, 0, 1, MPI.COMM_WORLD, &request)
</pre>
