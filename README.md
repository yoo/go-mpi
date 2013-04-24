# go-mpi

go-mpi are GO bindings for the Message Passing Interface <a href=http://www.mpi-forum.org/>(MPI Forum)</a>.

## Installation

MPI is a standard but the different implementations differ in some details.
The only supported implementation at the moment is <a href=http://www.open-mpi.de/>Open MPI</a>.

The install script uses pkg-config to determine the the include- and library
path for the MPI implementation.
<pre>
  git clone git://github.com/JohannWeging/go-mpi.git
  cd go-mpi
  ./install openmpi
</pre>

If the package differs from the default package name:
<pre>
  ./install openmpi --pkg-config &ltpackage_name&gt
</pre>

If the path can not be determined by pkg-config it can be set manually.
The library needs to be in the format of lib&ltname&gt.so
<pre>
  ./install openmpi --lib mympi --lib-path /usr/lib/mympilibrary
</pre>

Once the bindings support more than one implementation you may want to have more than one version of the bindings installed. By default the package name is MPI.
<pre>
  ./install openmpi --install-as openmpi
</pre>

In your program you can do:
<pre>
  package main

  import MPI "openmpi"

  [...]
</pre>
## Syntax

The syntax is similar to the C syntax of MPI.
<pre>
  &ltpackage_name&gt.Mpi_function(arguments)
</pre>

If the bindings are imported as "MPI":
<pre>
  err = MPI.Init(os.Args)
</pre>

Output parameter like request objects are returned by the function and not passed as pointers inside the arguments.

<pre>
  C:
  err = MPI_Irecv(recvBuffer, count, MPI_INT, 0, 1, MPI_COMM_WROLD, &amp

  GO:
  request, err = MPI.Irecv(recvBuffer, count, MPI.INT, 0, 1, MPI.COMM_WORLD)
</pre>
