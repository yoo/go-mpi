# go-mpi

go-mpi are GO bindings for the Message Passing Interface <a href=http://www.mpi-forum.org/>(MPI Forum)</a>.

## Installation

MPI is a standard but the different implementations differ in some details.
At the moment go-mpi support  <a href=http://www.open-mpi.de/>Open MPI</a> and <a href=http://www.mpich.org/>MPICH</a> version 2.

To tell go where to look for the MPI library use the CGO_LDFALG environment variable. The following instructions uses the default path for Open MPI and MPICH.

For Open MPI:
```sh
export CGO_LDFLAGS='-L/usr/lib/openmpi -lmpi'
go get -tags openmpi github.com/JohannWeging/go-mpi
```

For MPICH:
```sh
export CGO_LDFLAGS='-L/usr/lib/ -lmpich'
go get -tags mpich github.com/JohannWeging/go-mpi
```
## Syntax

The syntax is similar to the C syntax of MPI.
```
<package_name>.Mpi_function(arguments)
```

If the bindings are imported as "MPI":
<pre>
  err = MPI.Init(os.Args)
</pre>

Output parameter like request objects are returned by the function and not passed as pointers inside the arguments.

<pre>
  C:
  err = MPI_Irecv(recvBuffer, count, MPI_INT, 0, 1, MPI_COMM_WROLD, &request)

  GO:
  err = MPI.Irecv(recvBuffer, count, MPI.INT, 0, 1, MPI.COMM_WORLD, &request)
</pre>
