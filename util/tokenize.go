package util

import "strings"

func Tokenize(str string) []string {
	return strings.Split(str, "(")
}
