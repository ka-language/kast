package oat

import (
	"fmt"
	"os"

	oatenc "oat/format/encoding"
	"ka/lang/interpreter"
	"ka/lang/types"
)

func Run(params types.CliParams) {
	d, e := oatenc.KastDecode(params.Name)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

	interpreter.RunInterpreter(d, params)
}
