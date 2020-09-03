package oat

import (
	"fmt"
	"os"

	oatenc "github.com/omm-lang/oat/format/encoding"
	"github.com/omm-lang/omm/lang/interpreter"
	"github.com/omm-lang/omm/lang/types"
)

func Run(params types.CliParams) {
	d, e := oatenc.OatDecode(params.Name)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

	interpreter.RunInterpreter(d, params)
}
