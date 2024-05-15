package command

import (
	"context"
)

type ConfigObject struct{}

func (cmd ConfigObject) Name() string {
	return "config"
}

func (cmd ConfigObject) ShortDescr() string {
	return "persistent settings to customize modfit behavior"
}

func (cmd ConfigObject) ActionCmds() map[string]ActionCmd {
	return MapNames([]ActionCmd{
		new(ConfigGenerateAction),
		new(ConfigResolvePathAction),
	})
}

func (cmd ConfigObject) Run(cxt context.Context, args []string) {
	CallObjectAction(cxt, cmd, args)
}
