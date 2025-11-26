# A Makefile is a great way to run multiple terminal commands (especially handy for running multiple go test/go tool commands).
# Run "make test" or "make fuzz" in the terminal to run these commands/tests.
#
# Note the use of "./..." in some of the commands below.
# This is used to run the Go tests in the current directory, as well as in all sub-directories.
test:
	go tool staticcheck ./...
	go test -v ./...

fuzz:
	go test -run NONE -fuzz . -fuzztime 10s -v