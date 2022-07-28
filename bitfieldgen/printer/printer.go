package printer

import (
	"fmt"
	"github.com/godzilathakur/bitfieldgen/parser"
	"strings"
)

func PrintRegisterDefs(registerDefs parser.RegisterDefinitionsType) {
	fmt.Printf("RegisterDefinitions Name: %s\n", registerDefs.PeripheralName())
	fmt.Printf("ConfigWidth: %d\n", registerDefs.PeripheralConfig().Width())

	for _, register := range registerDefs.RegisterDefinitions() {
		fmt.Printf("Register Name: %s\n", register.Name)
		for _, field := range register.Fields() {
			fmt.Printf("Field Name: %s\n", field.Name())
			fmt.Printf("Field LSB: %d\n", field.Lsb())
			fmt.Printf("Field MSB: %d\n", field.Msb())

			if field.Values != nil {
				for name, value := range field.Values() {
					fmt.Printf("%s = %d\n",
						strings.ToUpper(name),
						value)
				}
			}
		}
	}
}
