# go-mpi

go-mpi are GO bindings for the Message Passing Interface <a href=http://www.mpi-forum.org/>(MPI Forum)</a>.

## Installation

MPI is a standard but the different implementations differ in some details.
At the moment go-mpi support  <a href=http://www.open-mpi.de/>Open MPI</a> and <a href=http://www.mpich.org/>MPICH</a> version 2.

The install script uses pkg-config to determine the the include- and library
path for the MPI implementation.
```
git clone git://github.com/JohannWeging/go-mpi.git
cd go-mpi
# For Open MPI
./install openmpi
# Or for MPICH2
./install mpich2
```

If the package differs from the default package name:
```
./install <implementation> --pkg-config <package_name>
```

If the path can not be determined by pkg-config it can be set manually.
The library needs to be in the format of:
```
lib<name>.so
```
```
./install <implementation> --lib mympi --lib-path /usr/lib/mympilibrary
```

Once the bindings support more than one implementation you may want to have more than one version of the bindings installed. By default the package name is MPI.
```
./install openmpi --install-as openmpi
./install mpich2 --install-as mpich2
```

Now you can use booth implementations in your program.
```
package main

import MPI "openmpi"

[...]
```
```
package main

import MPI "mpich2"

[...]
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
  request, err = MPI.Irecv(recvBuffer, count, MPI.INT, 0, 1, MPI.COMM_WORLD)
</pre>
