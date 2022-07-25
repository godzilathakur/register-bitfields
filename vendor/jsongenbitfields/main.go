package main

/*
  ___  _  _    ___  _       _     _    ___
 | _ )(_)| |_ | __|(_) ___ | | __| |  / __| ___  _ _
 | _ \| ||  _|| _| | |/ -_)| |/ _` | | (_ |/ -_)| ' \
 |___/|_| \__||_|  |_|\___||_|\__,_|  \___|\___||_||_|
*/

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var asciiWelcome []string = []string{
	` ___  _  _    ___  _       _     _    ___`,
	`| _ )(_)| |_ | __|(_) ___ | | __| |  / __| ___  _ _`,
	`| _ \| ||  _|| _| | |/ -_)| |/ _^ | | (_ |/ -_)| | \`,
	`|___/|_| \__||_|  |_|\___||_|\__,_|  \___|\___||_||_|`,
}

type Config struct {
	Width int `json:"width"`
}

type RegisterField struct {
	Name   string                 `json:"name"`
	Msb    int                    `json:"msb"`
	Lsb    int                    `json:"lsb"`
	Values map[string]interface{} `json:"values"`
}

type Register struct {
	Name   string          `json:"name"`
	Fields []RegisterField `json:"fields"`
}

type Definitions struct {
	PeripheralName string     `json:"peripheral_name"`
	Config         Config     `json:"config"`
	Registers      []Register `json:"registers"`
}

func printRegisterDefs(registerDefs Definitions) {
	fmt.Printf("Peripheral Name: %s\n", registerDefs.PeripheralName)
	fmt.Printf("Width: %d\n", registerDefs.Config.Width)

	for _, register := range registerDefs.Registers {
		fmt.Printf("Register Name: %s\n", register.Name)
		for _, field := range register.Fields {
			fmt.Printf("Field Name: %s\n", field.Name)
			fmt.Printf("Field LSB: %d\n", field.Lsb)
			fmt.Printf("Field MSB: %d\n", field.Msb)

			if field.Values != nil {
				for name, value := range field.Values {
					valInt := (int)(value.(float64))
					fmt.Printf("%s = %d\n",
						strings.ToUpper(name),
						valInt)
				}
			}
		}
	}
}

func convertFieldToCppStructBitField(field RegisterField) []string {
	result := []string{}
	return result
}

func convertFieldValuesToCppEnum(values map[string]interface{}, name string) []string {
	result := []string{}
	result = append(result, fmt.Sprintf("enum %s_t {", name))
	for name, value := range values {
		valInt := (int)(value.(float64))
		result = append(result, fmt.Sprintf("  %s = %d,",
			strings.ToUpper(name),
			valInt))
	}
	result = append(result, fmt.Sprintf("};"))
	return result
}

func generateRegisterDefs(registerDefs Definitions) {
	file, err := os.Create(strings.ToLower(registerDefs.PeripheralName) + "_register_defs.h")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintln(file, `#pragma once`)
	fmt.Fprintln(file, `// auto-generated file`)
	fmt.Fprintln(file)

	for _, register := range registerDefs.Registers {
		fmt.Fprintf(file, "struct %s_t {\n", register.Name)
		for _, field := range register.Fields {
			fmt.Fprintf(file, "  unsigned int %s : %d;\n",
				field.Name,
				field.Msb-field.Lsb+1)
			if field.Values != nil {
				for _, fieldEnumLine := range convertFieldValuesToCppEnum(field.Values, field.Name) {
					fmt.Fprintf(file, "  %s\n", fieldEnumLine)
				}
			}
		}
		fmt.Fprintf(file, "};\n\n")
	}
}

var registerDefsFileNamePtr = flag.String("file", "definitions.json", "register definitions file name")
var verbosePtr = flag.Bool("v", false, "print parsed definition")
var genCppHeaderPtr = flag.Bool("gen", false, "generate C++ header from register definitions")

func main() {
	for _, line := range asciiWelcome {
		fmt.Println(line)
	}
	for i := 0; i < 8; i++ {
		fmt.Println()
	}

	flag.Parse()

	fmt.Println(*registerDefsFileNamePtr)
	if file, err := ioutil.ReadFile(*registerDefsFileNamePtr); err != nil {
		fmt.Println(err)
	} else {
		registerDefs := Definitions{}
		if err := json.Unmarshal([]byte(file), &registerDefs); err != nil {
			fmt.Println(err)
		} else {
			if *verbosePtr == true {
				printRegisterDefs(registerDefs)
			}
			if *genCppHeaderPtr == true {
				generateRegisterDefs(registerDefs)
			}
		}
	}
}
