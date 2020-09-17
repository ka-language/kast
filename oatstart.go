package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tusklang/kore"
	oat "github.com/tusklang/oat/format"
	"github.com/tusklang/tusk/lang/types"
)

var output = flag.String("out", "", "Set output file")

func init() {
	flag.Usage = kore.Usagef("Tuskst")
}

func main() {
	flag.Parse()

	var cli_params types.CliParams

	if len(flag.Arg(1)) != 0 && flag.Arg(1)[0] == '-' {
		fmt.Println("Error, must list the filename as the second parameter")
		os.Exit(1)
	}

	var opt = flag.Arg(0)
	var filename = flag.Arg(1)
	cli_params.Name = filename

	if *output == "" {
		*output = strings.TrimSuffix(filename, filepath.Ext(filename)) + ".oat" //remove the .tusk and replace with .oat
	}
	cli_params.Output = *output

	if opt == "build" {
		//if they want to build an oat
		oat.Compile(cli_params)
	} else if opt == "run" {

		if *output != "" {
			fmt.Println("Error, cannot use -out while running an oat file")
			os.Exit(1)
		}

		os.Args = os.Args[2:] //remove the oat run

		//if they want to run an oat
		oat.Run(cli_params)
	} else {
		fmt.Println("Unrecognized option:", opt)
		os.Exit(1)
	}

	os.Exit(0)
}
