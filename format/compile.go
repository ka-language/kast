package oat

import (
	"fmt"

	. "omm/lang/types"

	"omm/lang/compiler"

	oatenc "oat/format/encoding"
)

//export Compile
func Compile(params CliParams) {
	vars, ce := compiler.Compile(params)

	if ce != nil {
		fmt.Println(ce)
	}

	oatenc.OatEncode(params.Output, vars)
}
