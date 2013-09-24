
all: test


test: test.go
	GOPATH=`readlink -f deps` ./go/bin/go build $^


clean:
	GOPATH=`readlink -f deps` ./go/bin/go clean

distclean: clean
	rm -rf deps/pkg/*

go/bin/go:
	wget "http://go.googlecode.com/files/go1.1.2.src.tar.gz"
	tar -xzf go1.1.2.src.tar.gz
	rm go1.1.2.src.tar.gz
	cd go/src && ./all.bash

deps: go/bin/go
	GOPATH=`readlink -f deps` ./go/bin/go get github.com/ErikDubbelboer/gspt


.PHONY: all clean distclean deps

