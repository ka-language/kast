package oat

import (
	"fmt"
	"os"

	oatenc "github.com/tusklang/oat/format/encoding"
	"github.com/tusklang/tusk/lang/interpreter"
	"github.com/tusklang/tusk/lang/types"
)

func Run(params types.CliParams) {
	d, e := oatenc.TuskstDecode(params.Name)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

	interpreter.RunInterpreter(d, params)
}
