ifeq ($(OS),Windows_NT)
	CLEAN_CMD = del
	BINARY = oat_start.exe
else
	CLEAN_CMD = rm -f
	BINARY = oat_start
endif

GOPATH = $(CURDIR)/../../

.PHONY: all
all:
	go build oat_start.go

.PHONY: clean
clean:
	-$(CLEAN_CMD) $(BINARY)