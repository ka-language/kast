package main

import (
	"fmt"
	"unsafe"

	"github.com/omm-lang/goat"
	oatenc "github.com/omm-lang/oat/format/encoding"
	"github.com/omm-lang/omm/lang/interpreter"
	"github.com/omm-lang/omm/lang/types"
)

//#include "helper.h"
import "C"

//export CallFunc
func CallFunc(oatf *C.char, fname *C.char, args *unsafe.Pointer, argc C.int) (unsafe.Pointer, *C.char) {

	//decompile the oat
	data, e := oatenc.OatDecode(C.GoString(oatf), 0)

	if e != nil {
		return nil, C.CString(e.Error())
	}

	//create the instance and get the variable
	var instance = goat.NewInstance(data)
	var fn = instance.Fetch(C.GoString(fname))

	//if it does not exist, error
	if fn == nil {
		return nil, C.CString(fmt.Sprintf("Variable %s does not exist", fname))
	}

	//if it is not a function, error
	if (*fn.Value).Type() != "function" {
		return nil, C.CString(fmt.Sprintf("Variable %s is not a function", fname))
	}

	//create the argv as an omm array
	var argv types.OmmArray

	for i := 0; i < int(argc); i++ {
		argv.PushBack(*(*types.OmmType)(C.getidx(args, C.int(i))))
	}

	//call the func
	return unsafe.Pointer(interpreter.Operations["function <- array"](*fn.Value, argv, instance, []string{"at goat caller"}, 0, "goat caller", 0)), C.CString("")
}

func main() {}
