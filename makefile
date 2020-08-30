ifeq ($(OS),Windows_NT)
	CLEAN_CMD = del
	GOATV_SHARED_DYNAMIC_LIB = goatv.dll
else
	CLEAN_CMD = rm -f
	GOATV_SHARED_DYNAMIC_LIB = goatv.so
endif

GOPATH = $(CURDIR)/../../../../

.PHONY: all
all:
	go build oat.go
	go build -buildmode=c-shared -o $(GOATV_SHARED_DYNAMIC_LIB) format/goatv/goatv.go

.PHONY: clean
clean:
	-$(CLEAN_CMD) $(BINARY)
	-$(CLEAN_CMD) $(GOATV_SHARED_DYNAMIC_LIB)