package attribute

type Definition struct {
	definitionType string
	attribute Attribute
}

func NewDefinition(attr Attribute) Definition {
	return Definition{
		definitionType: "attribute",
		attribute: attr,
	}
}

func (defn Definition) Type() string {
	return defn.definitionType
}

func (defn Definition) Get() Attribute {
	return defn.attribute
}