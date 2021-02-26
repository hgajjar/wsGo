package attribute

import (
	"github.com/hgajjar/wsGo/util"
)

type Attribute struct {
	key string
	value interface{}
}

const TYPE = "attribute"

func (attr Attribute) Key() string {
	return attr.key
}

func (attr Attribute) IsValid(key string) bool {
	return len(util.Tokenize(key)) == 1
}

func (attr *Attribute) Parse(key string, value string) {
	attr.key = key
	attr.value = value
}

func (attr Attribute) Type() string {
	return TYPE
}