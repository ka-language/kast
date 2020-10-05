ifeq ($(OS),Windows_NT)
	CLEAN_CMD = del
	BINARY = oatstart.exe
else
	CLEAN_CMD = rm -f
	BINARY = oatstart.out
endif

GOPATH = $(CURDIR)/../../../../

.PHONY: default
default: all

.PHONY: all
all:
	go build -o $(BINARY) oatstart.go

.PHONY: clean
clean:
	-$(CLEAN_CMD) $(BINARY)
