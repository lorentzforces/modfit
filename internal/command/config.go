package command

import (
	"context"
)

type ConfigDomain struct{}

func (cmd ConfigDomain) Name() string {
	return "config"
}

func (cmd ConfigDomain) ShortDescr() string {
	return "persistent settings to customize modfit behavior"
}

func (cmd ConfigDomain) ActionCmds() map[string]Action {
	return MapNames([]Action{
		new(ConfigGenerateAction),
		new(ConfigResolvePathAction),
	})
}

func (cmd ConfigDomain) Run(cxt context.Context, args []string) {
	CallDomainAction(cxt, cmd, args)
}
