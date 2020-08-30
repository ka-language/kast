ifeq ($(OS),Windows_NT)
	CLEAN_CMD = del
else
	CLEAN_CMD = rm -f
endif

GOPATH = $(CURDIR)/../../../../

.PHONY: all
all:
	go build oat.go

.PHONY: clean
clean:
	-$(CLEAN_CMD) $(BINARY)