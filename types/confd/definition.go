package confd

type Definition struct {
	definitionType string
	confd Confd
}

func NewDefinition(confd Confd) Definition {
	return Definition{
		definitionType: "confd",
		confd: confd,
	}
}

func (defn Definition) Type() string {
	return defn.definitionType
}