package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	oat "github.com/omm-lang/oat/format"
	"github.com/omm-lang/omm/lang/types"
)

var cwd = flag.String("cwd", "", "set cwd")

func main() {
	flag.Parse()

	var cli_params types.CliParams

	fmt.Println(flag.Args())

	var opt = flag.Arg(0)
	var filename = flag.Arg(1)
	cli_params.Name = filename
	var output = strings.TrimSuffix(filename, filepath.Ext(filename)) + ".oat" //remove the .omm and replace with .oat

	//get the current working directory
	//and change to it
	os.Chdir(*cwd)

	for i := 2; i < len(os.Args); i++ {
		if flag.Arg(i) == "-o" || flag.Arg(i) == "--output" {
			if i+1 < len(flag.Args()) {
				output = flag.Arg(i + 1)
			} else {
				fmt.Println("Expected a value after", flag.Arg(i))
				os.Exit(1)
			}
		}
	}

	if opt == "build" {
		//if they want to build an oat
		cli_params.Output = output
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
