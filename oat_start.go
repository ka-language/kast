package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	oat "github.com/omm-lang/oat/format"
	suite "github.com/omm-lang/omm-suite"
	"github.com/omm-lang/omm/lang/types"
)

var cwd = flag.String("cwd", "", "set the current working directory (automatically placed by the shell/pwsh script)")
var output = flag.String("o", "", "set output file")
var prec = flag.Uint64("prec", 20, "set the precision of an Omm instance")

func init() {
	flag.Usage = suite.Usagef("Oat")
}

func main() {
	flag.Parse()

	var cli_params types.CliParams

	var filenamei = 1
	for flag.Arg(filenamei) != "" && flag.Arg(filenamei)[0] == '-' {
		filenamei++ //only inside the block for formatting
	}

	if flag.Arg(filenamei-1) == "" {
		fmt.Println("Error, no input file was given")
		os.Exit(1)
	}

	var opt = flag.Arg(0)
	var filename = flag.Arg(filenamei)
	cli_params.Name = filename

	if *output == "" {
		*output = strings.TrimSuffix(filename, filepath.Ext(filename)) + ".oat" //remove the .omm and replace with .oat
	}
	cli_params.Output = *output
	cli_params.Prec = *prec

	if opt == "build" {
		//if they want to build an oat
		oat.Compile(cli_params)
	} else if opt == "run" {
		//if they want to run an oat
		oat.Run(cli_params)
	} else {
		fmt.Println("Unrecognized option:", opt)
		os.Exit(1)
	}

	os.Exit(0)
}
