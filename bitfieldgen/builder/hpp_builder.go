package builder

import (
	"errors"
	"fmt"
	"github.com/godzilathakur/bitfieldgen/parser"
	"log"
	"os"
	"strings"
)

const (
	tabSize = 2
)

func buildIoMethods(file *os.File, tab int) {
	fmt.Fprintf(file, `%svoid io_write(const unsigned int data) {

	};`, WhiteSpace(tabSize, tab))
	fmt.Fprintln(file)

	fmt.Fprintf(file, `%svoid io_read(unsigned int& data) {

	};`, WhiteSpace(tabSize, tab))
	fmt.Fprintln(file)
}

func buildStructBitfield(file *os.File, fields []parser.RegisterFieldType, registerWidthBits int, tab int) {
	fmt.Fprintf(file, "%sstruct fields_t {\n",
		WhiteSpace(tabSize, tab))

	tab++
	var lastMsb int
	var reservedFieldCount int
	text := []string{}
	for i := len(fields) - 1; i >= 0; i-- {
		field := fields[i]
		if field.Msb() >= registerWidthBits {
			text = append(text, fmt.Sprintf("%sunsigned int: 0;",
				WhiteSpace(tabSize, tab)))
		}
		if lastMsb != 0 && field.Lsb()-lastMsb != 1 {
			text = append(text, fmt.Sprintf("%sunsigned int %s_%d : %d; // %s",
				WhiteSpace(tabSize, tab),
				"reserved",
				reservedFieldCount,
				field.Lsb()-lastMsb-1,
				"UNSUPPORTED"))
			reservedFieldCount++
		}
		text = append(text, fmt.Sprintf("%sunsigned int %s : %d; // %s",
			WhiteSpace(tabSize, tab),
			field.Name(),
			field.Msb()-field.Lsb()+1,
			field.Attribute()))
		lastMsb = field.Msb()
	}
	for i := len(text) - 1; i >= 0; i-- {
		fmt.Fprintln(file, text[i])
	}
	tab--
	fmt.Fprintf(file, "%s} m_fields;\n",
		WhiteSpace(tabSize, tab))
	fmt.Fprintf(file, "%sunsigned int m_data;\n",
		WhiteSpace(tabSize, tab))
	fmt.Fprintln(file)

	fmt.Fprintf(file, `%sregister_defs_t() {
		io_read(m_data);
	};`, WhiteSpace(tabSize, tab))
	fmt.Fprintln(file)

	fmt.Fprintf(file, `%sregister_defs_t& operator=(const register_defs_t& other) {
		m_data = other.m_data;
		return *this;
	};`, WhiteSpace(tabSize, tab))
	fmt.Fprintln(file)

	fmt.Fprintf(file, `%svoid operator=(const unsigned int val) {
		m_data = val;
		io_write(m_data);
	};`, WhiteSpace(tabSize, tab))
	fmt.Fprintln(file)

	tab--
	fmt.Fprintf(file, "%s};\n",
		WhiteSpace(tabSize, tab))
	fmt.Fprintln(file)
}

func buildFieldValueEnum(file *os.File, field parser.RegisterFieldType, tab int) {
	fmt.Fprintf(file, "%senum class %s_values_t {\n",
		WhiteSpace(tabSize, tab),
		field.Name())
	tab++
	for name, value := range field.Values() {
		fmt.Fprintf(file, "%s%s = %d,\n",
			WhiteSpace(tabSize, tab),
			strings.ToUpper(name),
			value)
	}
	tab--
	fmt.Fprintf(file, "%s};\n", WhiteSpace(tabSize, tab))
	fmt.Fprintln(file)
}

func buildReadAccessorMethod(file *os.File, field parser.RegisterFieldType, tab int) {
	if len(field.Values()) != 0 {
		fmt.Fprintf(file, `%s%s_t read_%s() const {`,
			WhiteSpace(tabSize, tab),
			field.Name(),
			field.Name())
		fmt.Fprintln(file)
		tab++
		fmt.Fprintf(file, `%sauto defs = register_defs_t{};`, WhiteSpace(tabSize, tab))
		fmt.Fprintln(file)
		fmt.Fprintf(file, `%sreturn static_cast<%s_t>(defs.m_fields.%s);`,
			WhiteSpace(tabSize, tab),
			field.Name(),
			field.Name())
		fmt.Fprintln(file)
	} else {
		fmt.Fprintf(file, `%sunsigned int read_%s() const {`,
			WhiteSpace(tabSize, tab),
			field.Name())
		fmt.Fprintln(file)
		tab++
		fmt.Fprintf(file, `%sauto defs = register_defs_t{};`, WhiteSpace(tabSize, tab))
		fmt.Fprintln(file)
		fmt.Fprintf(file, `%sreturn defs.m_fields.%s;`,
			WhiteSpace(tabSize, tab),
			field.Name())
		fmt.Fprintln(file)
	}

	tab--
	fmt.Fprintf(file, "%s};\n", WhiteSpace(tabSize, tab))
	fmt.Fprintln(file)
}

func buildWriteAccessorMethod(file *os.File, field parser.RegisterFieldType, tab int) {
	if len(field.Values()) != 0 {
		fmt.Fprintf(file, `%svoid write_%s(%s_t val) {`,
			WhiteSpace(tabSize, tab),
			field.Name(),
			field.Name())
	} else {
		fmt.Fprintf(file, `%svoid write_%s(unsigned int val) {`,
			WhiteSpace(tabSize, tab),
			field.Name())
	}
	fmt.Fprintln(file)

	tab++
	fmt.Fprintf(file, `%sauto defs = register_defs_t{};`, WhiteSpace(tabSize, tab))
	fmt.Fprintln(file)
	fmt.Fprintf(file, `%sdefs.m_fields.%s = static_cast<unsigned int>(val);`,
		WhiteSpace(tabSize, tab),
		field.Name())
	fmt.Fprintln(file)
	fmt.Fprintf(file, `%sdefs = defs.m_data;`,
		WhiteSpace(tabSize, tab))
	fmt.Fprintln(file)

	tab--
	fmt.Fprintf(file, "%s};\n", WhiteSpace(tabSize, tab))
	fmt.Fprintln(file)
}

func buildAccessorMethods(file *os.File, field parser.RegisterFieldType, tab int) {
	switch field.Attribute() {
	case parser.READ_ONLY:
		buildReadAccessorMethod(file, field, tab)
	case parser.WRITE_ONLY:
		buildWriteAccessorMethod(file, field, tab)
	case parser.READ_WRITE:
		buildReadAccessorMethod(file, field, tab)
		buildWriteAccessorMethod(file, field, tab)
	}
}

func buildRegisterClass(file *os.File, register parser.RegisterType, registerWidthBits int, tab int) {
	fmt.Fprintf(file, "%sclass register_%s_t {\n",
		WhiteSpace(tabSize, tab),
		register.Name())
	fmt.Fprintf(file, "%sprivate:\n",
		WhiteSpace(tabSize, tab))

	tab++
	fmt.Fprintf(file, "%sunion register_defs_t {\n",
		WhiteSpace(tabSize, tab))

	tab++
	buildStructBitfield(file, register.Fields(), registerWidthBits, tab)

	tab--
	tab--
	fmt.Fprintf(file, "%spublic:\n",
		WhiteSpace(tabSize, tab))

	tab++
	for _, field := range register.Fields() {
		if field.Values() != nil {
			buildFieldValueEnum(file, field, tab)
		}
	}

	fmt.Fprintf(file, `%sregister_%s_t() = default;
	~register_%s_t() = default;`,
		WhiteSpace(tabSize, tab),
		register.Name(),
		register.Name())
	fmt.Fprintln(file)

	for _, field := range register.Fields() {
		buildAccessorMethods(file, field, tab)
	}

	tab--
	fmt.Fprintf(file, "%s};\n", WhiteSpace(tabSize, tab))
	fmt.Fprintln(file)
}

func generateCppRegisterDefs(file *os.File, registerDefs parser.RegisterDefinitionsType) {
	fmt.Fprintln(file, `#pragma once`)
	fmt.Fprintln(file, `/* auto-generated file using bitfieldgen`)
	fmt.Fprintln(file)
	fmt.Fprintf(file, ` Peripheral Name %s`, registerDefs.PeripheralName())
	fmt.Fprintln(file)
	fmt.Fprintf(file, ` Description %s`, registerDefs.PeripheralDescription())
	fmt.Fprintln(file)
	fmt.Fprintf(file, ` Specifications %s`, registerDefs.PeripheralSpecUrl())
	fmt.Fprintln(file)
	fmt.Fprintln(file, `*/`)
	fmt.Fprintln(file)

	tab := 0
	buildIoMethods(file, tab)

	for _, register := range registerDefs.RegisterDefinitions() {
		buildRegisterClass(file, register, registerDefs.PeripheralConfig().Width(), tab)
		fmt.Fprintln(file)
	}
	log.Println("Generated C++ header")
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
