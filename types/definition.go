package types

import "github.com/pkg/errors"

type Definition struct {
	attribute Attribute
	command Command
}

func (defn Definition) isAttribute() bool {
	return defn.attribute.key != ""
}

func (defn Definition) isCommand() bool {
	return defn.command.name != ""
}

func (defn Definition) Attribute() (Attribute, error) {
	if ! defn.isAttribute() {
		return Attribute{}, errors.Errorf("%s is not valid attribute key", defn.attribute.key)
	}

	return defn.attribute, nil
}

func (defn Definition) Command() (Command, error) {
	if ! defn.isCommand() {
		return Command{}, errors.Errorf("%s is not valid command name", defn.command.name)
	}

	return defn.command, nil
}