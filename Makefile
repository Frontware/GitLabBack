# GoLang parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
STRIPFLAGS=-ldflags="-s -w"
BINARY_FILE=GitLabBack

ifeq ($(GOARCH),)
	GOARCH=amd64
endif

all: linux windows

linux:
	GOARCH=$(GOARCH) GOOS=$@ $(GOBUILD) $(STRIPFLAGS) -o $(BINARY_FILE)_$@_$(GOARCH)

windows:
	GOARCH=$(GOARCH) GOOS=$@ $(GOBUILD) $(STRIPFLAGS) -o $(BINARY_FILE)_$@_$(GOARCH)

clean:
	$(GOCLEAN)
	rm -f $(BINARY_FILE)_*
