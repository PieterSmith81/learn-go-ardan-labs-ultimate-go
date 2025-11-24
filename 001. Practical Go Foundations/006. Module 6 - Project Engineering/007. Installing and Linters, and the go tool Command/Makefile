# A Makefile is a great way to run multiple terminal commands (especially handy for running multiple go test/go tool commands).
# Run "make test" or "make fuzz" in the terminal to run these commands/tests.
test:
	go tool staticcheck .
	go test -v

fuzz:
	go test -run NONE -fuzz . -fuzztime 10s -v