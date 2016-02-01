# go-lib
Utilities and libraries for golang

##Input
It's an input module. What do you expect?  Depending on the function,
it will take in input and return either a string, an int or a double.
Errors aren't returned because it is handled within the function.

##Other
Random things I personally need quite often. I compacted it into one
module just for ease of access. Ignore it if you don't need anything
from it.

##Solver
A calculator that can do assorted math functions that I am too lazy to
do each time for math homework. If you want to use it, copy and paste
this into a main.go and run it:

    package main
    import m "github.com/whitman-colm/go-lib/src/solver"

    func main() {
      m.Start()
    }

#End (for now)

