package factory

import (
	"github.com/hgajjar/wsGo/types"
	"github.com/hgajjar/wsGo/types/attribute"
	"github.com/hgajjar/wsGo/types/command"
	"github.com/hgajjar/wsGo/types/confd"
)

func GetAvailableConfigTypes() []types.ConfigType {
	return []types.ConfigType{
		&attribute.Attribute{},
		&command.Command{},
		&confd.Confd{},
	}
}
