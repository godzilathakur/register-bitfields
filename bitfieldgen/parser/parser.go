package parser

type PeripheralConfigType interface {
	Width() int
}

type RegisterFieldAttribute int

const (
	READ_ONLY   RegisterFieldAttribute = 0
	WRITE_ONLY                         = 1
	READ_WRITE                         = 2
	UNSUPPORTED                        = -1
)

func (attr RegisterFieldAttribute) String() string {
	switch attr {
	case READ_ONLY:
		return "READ_ONLY"
	case WRITE_ONLY:
		return "WRITE_ONLY"
	case READ_WRITE:
		return "READ_WRITE"
	}
	return "UNSUPPORTED"
}

type RegisterFieldType interface {
	Name() string
	Attribute() RegisterFieldAttribute
	Msb() int
	Lsb() int
	Values() map[string]int
}

type RegisterType interface {
	Name() string
	Fields() []RegisterFieldType
}

type RegisterDefinitionsType interface {
	PeripheralName() string
	PeripheralConfig() PeripheralConfigType
	RegisterDefinitions() []RegisterType
}

type RegisterDefinitionsParser interface {
	ParseRegisterDefinitions([]byte) (RegisterDefinitionsType, error)
}
