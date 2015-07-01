##
## Makefile for golang project
##

SRCS = hop_data.go main.go
TARGET = metropolis_V2
GCFLAGS = -gcflags='-B -largemodel'
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

## gccgo
##
GCCGO_FLAGS = -O3 -fno-go-check-divide-zero -fno-go-check-divide-overflow -static-libgo

$(TARGET).gccgo gccgo gg: $(SRCS)
	@go build -o $(TARGET).gccgo -gccgoflags="$(GCCGO_FLAGS)" -x -compiler gccgo .
