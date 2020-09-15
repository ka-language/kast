ifeq ($(OS),Windows_NT)
	CLEAN_CMD = del
	BINARY = kast_start.exe
else
	CLEAN_CMD = rm -f
	BINARY = kast_start
endif

GOPATH = $(CURDIR)/../../

.PHONY: all
all:
	go build kast_start.go

.PHONY: clean
clean:
	-$(CLEAN_CMD) $(BINARY)