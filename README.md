# genscratch

A C++ scratch environment generator (along with Makefile) implemented in Go.

### Install

$ go install github.com/ellemouton/genscratch

### Run

If no environment variable is set then a scratch playground will be created in the current directory.
Otherwise, set the 'GENSCRATCHPATH' environment variable to your chosen scratch directory and then this will be used.
Each scratch environment will be in its own sub directory.

$ genscratch

You can overide the name of the scratch directory as well as the path to where it will be created as follows:

$ genscratch --path=. --name=testIO

### Updating the Makefile

To get the most out of this tool you should also install 'genmake' ($ go get github.com/ellemouton/genmake). That way, if you add any extra source or header files in your scratch directory, you can automatically update the Makefile appropriately by running:

$ genmake

### Recommended aliases:

$ alias gm="genmake"

$ alias gs="genscratch"
