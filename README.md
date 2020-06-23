# Sketch

Sample usage of gRPC.

## Fetch dependencies

There are two dependencies (and a truckload of transitive
dependencies).  *`protolint` imports a lot of stuff, which is why the
`go.sum` file looks like someone put all the universe's lost socks in
that file* 

To fetch them do a:

    make deps
	
This should fetch the dependencies.

## Build

The default build target is the most useful since it runs linting,
vetting and testing.  For more details look at the `Makefile`.
