package builder

import (
	"errors"
	"fmt"
	"github.com/godzilathakur/bitfieldgen/parser"
	"log"
	"os"
	"strings"
)

func convertFieldValuesToCppEnum(values map[string]int, name string) []string {
	result := []string{}
	result = append(result, fmt.Sprintf("enum class %s_t {", name))
	for name, value := range values {
		result = append(result, fmt.Sprintf("  %s = %d,",
			strings.ToUpper(name),
			value))
	}
	result = append(result, fmt.Sprintf("};"))
	return result
}

func generateCppRegisterDefs(file *os.File, registerDefs parser.RegisterDefinitionsType) {
	fmt.Fprintln(file, `#pragma once`)
	fmt.Fprintln(file, `// auto-generated file using bitfield-gen`)
	fmt.Fprintln(file)

	for _, register := range registerDefs.RegisterDefinitions() {
		fmt.Fprintf(file, "struct %s_t {\n", register.Name())
		for _, field := range register.Fields() {
			if field.Msb() >= registerDefs.PeripheralConfig().Width() {
				fmt.Fprintf(file, "  unsigned int: 0;\n")
			}
			fmt.Fprintf(file, "  unsigned int %s : %d;\n",
				field.Name(),
				field.Msb()-field.Lsb()+1)
			if field.Values != nil {
				for _, fieldEnumLine := range convertFieldValuesToCppEnum(field.Values(), field.Name()) {
					fmt.Fprintf(file, "  %s\n", fieldEnumLine)
				}
			}
		}
		fmt.Fprintf(file, "};\n\n")
		log.Println("Generated C++ header")
	}
}

type HppBuilder struct {
	Filename string
}

func (hpp HppBuilder) BuildHeader(defs interface{}) error {
	registerDefs, ok := defs.(parser.RegisterDefinitionsType)
	if !ok {
		return errors.New("Unable to convert to register parser")
	}
	file, err := os.Create(strings.ToLower(registerDefs.PeripheralName()) + hpp.Filename)
	defer file.Close()
	if err != nil {
		return err
	} else {
		generateCppRegisterDefs(file, registerDefs)
	}
	return nil
}
