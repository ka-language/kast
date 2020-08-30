package oat

import (
	"fmt"
	"io/ioutil"
	"os"

	. "github.com/omm-lang/omm/lang/types"

	"github.com/omm-lang/omm/lang/compiler"

	oatenc "github.com/omm-lang/oat/format/encoding"
)

//export Compile
func Compile(params CliParams) {
	fileName := params.Name

	file, e := ioutil.ReadFile(fileName)

	if e != nil {
		fmt.Println("Could not find file:", fileName)
		os.Exit(1)
	}

	vars, ce := compiler.Compile(string(file), fileName, params)

	if ce != nil {
		ce.Print()
	}

	oatenc.OatEncode(params.Output, vars)
}
