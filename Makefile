# GoLang parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
STRIPFLAGS=-ldflags="-s -w"
BINARY_FILE=GitLabBack

all: linux windows

linux:
	GOOS=$@ $(GOBUILD) $(STRIPFLAGS) -o $(BINARY_FILE)_$@

windows:
	GOOS=$@ $(GOBUILD) $(STRIPFLAGS) -o $(BINARY_FILE)_$@

clean:
	$(GOCLEAN)
	rm -f $(BINARY_FILE)_linux
	rm -f $(BINARY_FILE)_windows
