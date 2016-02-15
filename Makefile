##
## Makefile for golang project
##

SHELL = /bin/bash

SRCS = main.go
TARGET = metropolis_V2
GCFLAGS = -gcflags='-B'
RELEASE_DIR = ../Release_V1

test:
	@go test -v ./...

tags:
	@find -name "*.go" | xargs etags

release rel:
	@test -d "$(RELEASE_DIR)" || mkdir "$(RELEASE_DIR)"
	@git archive --worktree-attributes --format=tar HEAD | tar -C "$(RELEASE_DIR)" -xf -

$(TARGET) b build: $(SRCS)
	@go build $(GCFLAGS) -o $(TARGET) .

testrun t:
	@./$(TARGET) -N 10000 > log 2>&1 &

N1e5:
	@./$(TARGET) -N 100000 > log 2>&1 &

N3e5:
	@./$(TARGET) -N 300000 -Temp 200,300,10 > log 2>&1 &

## gccgo
##
GCCGO_FLAGS = -O3 -fno-go-check-divide-zero -fno-go-check-divide-overflow -static-libgo

$(TARGET).gccgo gccgo gg: $(SRCS)
	@go build -o $(TARGET).gccgo -gccgoflags="$(GCCGO_FLAGS)" -x -compiler gccgo .
