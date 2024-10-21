ifeq ($(OS),Windows_NT)
    SOURCES := $(shell dir /S /B *.go)
else
    SOURCES := $(shell find . -name '*.go')
endif

ifeq ($(shell uname),Darwin)
    GOOS = darwin
    GOARCH = amd64
    EXEEXT =
else ifeq ($(shell uname),Linux)
    GOOS = linux
    GOARCH = $(shell arch)
    EXEEXT =
else ifeq ($(OS),Windows_NT)
    GOOS = windows
    GOARCH = amd64
    EXEEXT = .exe
endif

APP := fatehound$(EXEEXT)
TARGET := ./dist/fatehound_$(GOOS)_$(GOARCH)_v1/$(APP)

$(APP): $(TARGET)
	cp $< $@

$(TARGET): $(SOURCES)
	gofumpt -w $(SOURCES)
	goreleaser release --snapshot --skip-validate --clean --skip=publish
	go vet ./...

all:
	goreleaser build --snapshot --clean

.PHONY: clean
clean:
	rm -f fatehound
	rm -f $(TARGET)
	rm -rf dist
