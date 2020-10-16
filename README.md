# genscratch

A C++ scratch environment generator (along with Makefile) implemented in Go.

### Install

$ go install github.com/ellemouton/genscratch

### Run

If no environment variable is set then a scratch playground will be created in the current directory.
Otherwise, set the 'GENSCRATCHDIR' environment variable to your chosen scratch directory and then this will be used.
Each scratch environment will be in its own sub directory.

$ genscratch
