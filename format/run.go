package oat

import (
	"fmt"
	"os"

	oatenc "oat/format/encoding"
	"omm/lang/interpreter"
	"omm/lang/types"
)

func Run(params types.CliParams) {
	d, e := oatenc.OatDecode(params.Name)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

	interpreter.RunInterpreter(d, params)
}
