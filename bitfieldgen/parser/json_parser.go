package parser

import (
	"encoding/json"
	"errors"
	"strings"
)

type Config struct {
	ConfigWidth int `json:"width"`
}

func (config Config) Width() int {
	return config.ConfigWidth
}

type Field struct {
	FieldName      string                 `json:"name"`
	FieldAttribute string                 `json:"attribute"`
	FieldMsb       int                    `json:"msb"`
	FieldLsb       int                    `json:"lsb"`
	FieldValues    map[string]interface{} `json:"values"`
}

func (field Field) Name() string {
	return field.FieldName
}

func (field Field) Attribute() RegisterFieldAttribute {
	if strings.Compare("r", field.FieldAttribute) == 0 {
		return READ_ONLY
	} else if strings.Compare("w", field.FieldAttribute) == 0 {
		return WRITE_ONLY
	} else if strings.Compare("rw", field.FieldAttribute) == 0 {
		return READ_WRITE
	}
	return UNSUPPORTED
}

func (field Field) Msb() int {
	return field.FieldMsb
}

func (field Field) Lsb() int {
	return field.FieldLsb
}

func (field Field) Values() map[string]int {
	if len(field.FieldValues) == 0 {
		return nil
	}
	values := map[string]int{}
	for name, val := range field.FieldValues {
		values[name] = (int)(val.(float64))
	}
	return values
}

type Register struct {
	RegisterName   string  `json:"name"`
	RegisterFields []Field `json:"fields"`
}

func (reg Register) Name() string {
	return reg.RegisterName
}

func (reg Register) Fields() []RegisterFieldType {
	fields := []RegisterFieldType{}
	for _, f := range reg.RegisterFields {
		fields = append(fields, f)
	}
	return fields
}

type RegisterDefinitions struct {
	Name      string     `json:"peripheral_name"`
	Config    Config     `json:"config"`
	Registers []Register `json:"registers"`
}

func (def RegisterDefinitions) PeripheralName() string {
	return def.Name
}

func (def RegisterDefinitions) PeripheralConfig() PeripheralConfigType {
	return def.Config
}

func (def RegisterDefinitions) RegisterDefinitions() []RegisterType {
	regs := []RegisterType{}
	for _, r := range def.Registers {
		regs = append(regs, r)
	}
	return regs
}

type JsonRegDefParser struct{}

func (parser *JsonRegDefParser) ParseRegisterDefinitions(text []byte) (RegisterDefinitionsType, error) {
	regDefs := RegisterDefinitions{}
	if err := json.Unmarshal(text, &regDefs); err != nil {
		return regDefs, errors.Unwrap(err)
	}
	return regDefs, nil
}
