package oat

import (
	"fmt"

	. "github.com/tusklang/tusk/lang/types"

	"github.com/tusklang/tusk/lang/compiler"

	oatenc "github.com/tusklang/oat/format/encoding"
)

//export Compile
func Compile(params CliParams) {
	vars, ce := compiler.Compile(params)

	if ce != nil {
		fmt.Println(ce)
	}

	oatenc.TuskstEncode(params.Output, vars)
}
