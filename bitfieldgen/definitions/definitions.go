package definitions

import (
	"encoding/json"
	"errors"
)

type Config struct {
	Width int `json:"width"`
}

type Field struct {
	Name      string                 `json:"name"`
	Attribute string                 `json:"attribute"`
	Msb       int                    `json:"msb"`
	Lsb       int                    `json:"lsb"`
	Values    map[string]interface{} `json:"values"`
}

type Register struct {
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

type RegisterDefinitions struct {
	PeripheralName string     `json:"peripheral_name"`
	Config         Config     `json:"config"`
	Registers      []Register `json:"registers"`
}

type RegisterDefinitionsParser interface {
	ParseRegisterDefinitions([]byte, *RegisterDefinitions)
}

type JsonRegDefParser struct{}

func (parser *JsonRegDefParser) ParseRegisterDefinitions(text []byte) (RegisterDefinitions, error) {
	regDefs := RegisterDefinitions{}
	if err := json.Unmarshal(text, &regDefs); err != nil {
		return regDefs, errors.Unwrap(err)
	}
	return regDefs, nil
}
