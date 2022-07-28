package main

import (
	"flag"
	"fmt"
	"github.com/godzilathakur/bitfieldgen/builder"
	"github.com/godzilathakur/bitfieldgen/parser"
	"github.com/godzilathakur/bitfieldgen/printer"
	"io/ioutil"
	"log"
	"path/filepath"
)

var asciiWelcome []string = []string{
	` ___  _  _    ___  _       _     _    ___`,
	`| _ )(_)| |_ | __|(_) ___ | | __| |  / __| ___  _ _`,
	`| _ \| ||  _|| _| | |/ -_)| |/ _^ | | (_ |/ -_)| | \`,
	`|___/|_| \__||_|  |_|\___||_|\__,_|  \___|\___||_||_|`,
}

var registerDefsFileNamePtr = flag.String("file", "definitions.json", "register definitions file name")
var verbosePtr = flag.Bool("v", false, "print parsed definition")
var genCppHeaderPtr = flag.Bool("gencpp", false, "generate C++ header from register defs")
var genCHeaderPtr = flag.Bool("genc", false, "generate C header from register defs")
var genAll = flag.Bool("gen", false, "generate C and C++ headers from register defs")

// @TODO
// var genRustHeaderPtr = flag.Bool("genrust", false, "generate rust header from register parser")

func main() {
	for _, line := range asciiWelcome {
		fmt.Println(line)
	}
	for i := 0; i < 8; i++ {
		fmt.Println()
	}

	flag.Parse()

	fmt.Println("Generating for ", *registerDefsFileNamePtr)

	var parserHandle parser.RegisterDefinitionsParser
	if filepath.Ext(*registerDefsFileNamePtr) == ".json" {
		parserHandle = &parser.JsonRegDefParser{}
	} else {
		log.Println("Only JSON definition file supported currently... aborting")
		return
	}

	if text, err := ioutil.ReadFile(*registerDefsFileNamePtr); err != nil {
		fmt.Println(err)
	} else {
		if registerDefs, err := parserHandle.ParseRegisterDefinitions(text); err != nil {
			log.Fatal(err)
		} else {
			if *verbosePtr == true {
				printer.PrintRegisterDefs(registerDefs)
			}
			if *genCHeaderPtr == true || *genAll == true {
				hBuilder := builder.HBuilder{"_c_register_defs.h"}
				hBuilder.BuildHeader(registerDefs)
			}
			if *genCppHeaderPtr == true || *genAll == true {
				hppBuilder := builder.HppBuilder{"_cpp_register_defs.h"}
				hppBuilder.BuildHeader(registerDefs)
			}
		}
	}
}
