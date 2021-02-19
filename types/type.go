package types

import "strings"

type ConfigType interface {
	IsValid(key string) bool
	Parse(key string, value string) Definition
}

func tokenize(str string) []string {
	return strings.Split(str, "(")
}