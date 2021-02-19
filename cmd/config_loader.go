package cmd

import (
	"github.com/hgajjar/wsGo/types"
	"github.com/spf13/viper"
)

func init() {
	InitConfig()

	loader := configLoader{
		types: []types.ConfigType{
			types.Attribute{},
			types.Command{},
		},
	}

	loader.loadConfig()
}

type configLoader struct {
	types []types.ConfigType
}

func (loader configLoader) loadConfig() {
	definitions := loader.getDefinitions()
	loader.loadCommands(definitions)
}

func (loader configLoader) getDefinitions() []types.Definition {
	var definitions []types.Definition

	for _, key := range viper.AllKeys() {
		for _, configType := range loader.types {
			if !configType.IsValid(key) {
				continue
			}
			definitions = append(definitions, configType.Parse(key, viper.GetString(key)))
		}
	}

	return definitions
}

func (loader configLoader) loadCommands(defns []types.Definition) {
	for _, defn := range defns {
		if cmd, err := defn.Command(); err == nil {
			rootCmd.AddCommand(buildCommand(cmd))
		}
	}
}