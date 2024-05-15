package command

import (
	"context"
)

type ModObject struct{}

func (cmd ModObject) Name() string {
	return "mod"
}

func (cmd ModObject) ShortDescr() string {
	return "user modification to a game"
}

func (cmd ModObject) ActionCmds() map[string]ActionCmd {
	return MapNames([]ActionCmd{
		new(ModIngestAction),
	})
}

func (cmd ModObject) Run(cxt context.Context, args []string) {
	CallObjectAction(cxt, cmd, args)
}

