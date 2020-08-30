package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	oat "github.com/omm-lang/oat/format"
	"github.com/omm-lang/omm/lang/types"
)

func main() {

	var cli_params types.CliParams

	if len(os.Args) <= 2 {
		fmt.Println("Error, no option was given")
		os.Exit(1)
	}

	var opt = os.Args[1]
	var filename = os.Args[2]
	cli_params.Name = filename
	var output = strings.TrimSuffix(filename, filepath.Ext(filename)) + ".oat" //remove the .omm and replace with .oat

	//get the current working directory
	//and change to it
	dirname, _ := os.Getwd()
	os.Chdir(dirname)

	for i := 3; i < len(os.Args); i++ {
		if os.Args[i] == "-o" || os.Args[i] == "--output" {
			if i+1 < len(os.Args) {
				output = os.Args[i+1]
			} else {
				fmt.Println("Expected a value after", os.Args[i])
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
