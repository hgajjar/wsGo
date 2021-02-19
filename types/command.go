package types

import (
	"fmt"
	"regexp"
	"strings"
)

type Command struct {
	name string
	usage string
	flags []*Flag // flag value needs to be mutable at runtime, that is why we need a pointer here
	script string
}

func NewCommand(name string, usage string, flags []*Flag, script string) Command {
	return Command{name: name, usage: usage, flags: flags, script: script}
}

type Flag struct {
	name string
	usage string
	defaultValue string
	value *string
}

func NewFlag(name string, usage string, defaultValue string, value *string) *Flag {
	return &Flag{name: name, usage: usage, defaultValue: defaultValue, value: value}
}

var cmdRegex = regexp.MustCompile(`['|"]([a-z]+)(.*)['|"]`)
var flagRegex = regexp.MustCompile(`-([a-z]+)=`)

func (c Command) IsValid(key string) bool {
	tokens := tokenize(key)

	if len(tokens) <= 1 || tokens[0] != "command" {
		return false
	}

	if !cmdRegex.MatchString(tokens[1]) {
		panic(fmt.Sprintf("Invalid command definition: %s", tokens[1]))
	}

	return true
}

func (c Command) Parse(key string, value string) Definition {
	tokens := tokenize(key)
	cmdParts := cmdRegex.FindStringSubmatch(tokens[1])

	name := cmdParts[1]
	usage := strings.Trim(strings.Trim(cmdParts[0], "'"), "\"")
	var flags []*Flag

	for _, cmdChunk := range strings.Split(strings.Trim(cmdParts[2], " "), " ") {
		if flagRegex.MatchString(cmdChunk) {
			// it's a flag
			flagName := flagRegex.FindStringSubmatch(cmdChunk)
			flags = append(flags, &Flag{name: flagName[1]})
		} else {
			// it's an argument
		}
	}

	return Definition{
		command: Command{
			name:   name,
			usage:  usage,
			flags:  flags,
			script: value,
		},
	}
}

func (c Command) Name() string {
	return c.name
}

func (c Command) Usage() string {
	return c.usage
}

func (c *Command) Flags() []*Flag {
	return c.flags
}

func (c Command) Script() string {
	return c.script
}

func (f Flag) Name() string {
	return f.name
}

func (f Flag) DefaultValue() string {
	return f.defaultValue
}

func (f Flag) Usage() string {
	return f.usage
}

func (f Flag) Value() *string {
	return f.value
}

func (f *Flag) SetValue(val *string) {
	f.value = val
}