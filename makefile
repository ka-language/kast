ifeq ($(OS),Windows_NT)
	CLEAN_CMD = del
	BINARY = oatstart.exe
else
	CLEAN_CMD = rm -f
	BINARY = oatstart
endif

GOPATH = $(CURDIR)/../../../../

.PHONY: all
all: go.mod
	go get -u
	go build oatstart.go

go.mod:
	go mod init

.PHONY: clean
clean:
	-go mod tidy
	-$(CLEAN_CMD) $(BINARY)
