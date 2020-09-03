package oat

import (
	"fmt"

	. "github.com/omm-lang/omm/lang/types"

	"github.com/omm-lang/omm/lang/compiler"

	oatenc "github.com/omm-lang/oat/format/encoding"
)

//export Compile
func Compile(params CliParams) {
	vars, ce := compiler.Compile(params)

	if ce != nil {
		fmt.Println(ce)
	}

	oatenc.OatEncode(params.Output, vars)
}
