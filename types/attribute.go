package types

type Attribute struct {
	key string
	value interface{}
}

func (attr Attribute) IsValid(key string) bool {
	return len(tokenize(key)) == 1
}

func (attr Attribute) Parse(key string, value string) Definition {
	return Definition{
		attribute: Attribute{
			key: key,
			value: value,
		},
	}
}