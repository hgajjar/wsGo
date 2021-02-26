package cmd

import (
	"github.com/hgajjar/wsGo/factory"
	"github.com/hgajjar/wsGo/types"
	"github.com/hgajjar/wsGo/types/command"
	"github.com/spf13/viper"
)

func init() {
	InitConfig()

	loader := configLoader{
		types: factory.GetAvailableConfigTypes(),
	}

	loader.loadConfig()
}

type configLoader struct {
	types []types.ConfigType
}

func (loader configLoader) loadConfig() {
	loader.parseTypes()
	loader.loadCommands()
}

func (loader configLoader) parseTypes() {
	for _, key := range viper.AllKeys() {
		for _, configType := range loader.types {
			if !configType.IsValid(key) {
				continue
			}
			configType.Parse(key, viper.GetString(key))
		}
	}
}

func (loader configLoader) loadCommands() {
	for _, configType := range loader.types {
		if configType.Type() == command.TYPE {
			rootCmd.AddCommand(buildCommand(configType))
		}
	}
}