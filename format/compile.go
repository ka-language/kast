package oat

import (
	"fmt"

	. "tusk/lang/types"

	"tusk/lang/compiler"

	oatenc "oat/format/encoding"
)

//export Compile
func Compile(params CliParams) {
	vars, ce := compiler.Compile(params)

	if ce != nil {
		fmt.Println(ce)
	}

	oatenc.TuskstEncode(params.Output, vars)
}
