package oat

import (
	"fmt"

	. "ka/lang/types"

	"ka/lang/compiler"

	oatenc "oat/format/encoding"
)

//export Compile
func Compile(params CliParams) {
	vars, ce := compiler.Compile(params)

	if ce != nil {
		fmt.Println(ce)
	}

	oatenc.KastEncode(params.Output, vars)
}
