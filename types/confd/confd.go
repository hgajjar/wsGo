package confd

import (
	"fmt"
	"github.com/hgajjar/wsGo/util"
	"regexp"
)

const TYPE = "confd"

type Confd struct {
	path string
	templates []Template
}

type Template struct {
	src string
	dst string
}

func NewConfd(path string, templates []Template) *Confd {
	return &Confd{path: path, templates: templates}
}

func (c Confd) Type() string {
	return TYPE
}

var confdRegex = regexp.MustCompile(`['|"](/.*)['|"]`)

func (c Confd) IsValid(key string) bool {
	tokens := util.Tokenize(key)

	if len(tokens) <= 1 || tokens[0] != "confd" {
		return false
	}

	if !confdRegex.MatchString(tokens[1]) {
		panic(fmt.Sprintf("Invalid confd definition: %s", tokens[1]))
	}

	return true
}

func (c *Confd) Parse(key string, value string) {
	tokens := util.Tokenize(key)
	parts := confdRegex.FindStringSubmatch(tokens[1])

	c.path = parts[1]
}

func (c Confd) Path() string {
	return c.path
}

func (c Confd) Apply() {

}