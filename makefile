ifeq ($(OS),Windows_NT)
	CLEAN_CMD = del
	BINARY = oatstart.exe
else
	CLEAN_CMD = rm -f
	BINARY = oatstart
endif

GOPATH = $(CURDIR)/../../../../

.PHONY: all
all:
	go build oatstart.go

.PHONY: clean
clean:
	-$(CLEAN_CMD) $(BINARY)
