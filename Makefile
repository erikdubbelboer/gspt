
all:


go/bin/go:
	wget "http://go.googlecode.com/files/go1.1.2.src.tar.gz"
	tar -xzf go1.1.2.src.tar.gz
	rm go1.1.2.src.tar.gz
	cd go/src && ./all.bash

deps: go/bin/go


test: deps
	GOPATH=`readlink -f deps` ./go/bin/go test


.PHONY: all test

