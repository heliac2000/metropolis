##
## Makefile for golang project
##

SHELL = /bin/bash

SRCS = main.go
TARGET = metropolis_V2
GCFLAGS = -gcflags='-B'
LDFLAGS = -ldflags='-s'
RELEASE_DIR = ../Release_V1

test:
	@go test -v ./...

tags:
	@find -name "*.go" | xargs etags

release rel:
	@test -d "$(RELEASE_DIR)" || mkdir "$(RELEASE_DIR)"
	@git archive --worktree-attributes --format=tar HEAD | tar -C "$(RELEASE_DIR)" -xf -

$(TARGET) b build: $(SRCS)
	@go build $(GCFLAGS) $(LDFLAGS) -o $(TARGET) .

testrun t:
	@./$(TARGET) -N 10000 > log 2>&1 &

##
## Execution tags
##
## Temparature sequence: 10K increments from 200K to 300K
##
N6e5 N3e5 N1e5 N1e4:
	@./$(TARGET) -N $(subst N,,$@) > $(TARGET)_$@.log 2>&1 &

NH2_test CH3_test:
	@./$(TARGET) -N 1e4 -DataDir data_$(subst _test,,$@) > $(TARGET)_$@.log 2>&1 &

## gccgo
##
GCCGO_FLAGS = -O3 -fno-go-check-divide-zero -fno-go-check-divide-overflow -static-libgo

$(TARGET).gccgo gccgo gg: $(SRCS)
	@go build -o $(TARGET).gccgo -gccgoflags="$(GCCGO_FLAGS)" -x -compiler gccgo .
