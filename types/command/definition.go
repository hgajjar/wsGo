package command

type Definition struct {
	definitionType string
	command Command
}

func NewDefinition(cmd Command) Definition {
	return Definition{
		definitionType: "command",
		command: cmd,
	}
}

func (defn Definition) Type() string {
	return defn.definitionType
}

func (defn Definition) Get() Command {
	return defn.command
}