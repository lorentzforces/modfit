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

func (cmd ModObject) actionCmds() map[string]actionCmd {
	return MapNames([]actionCmd{
		new(ModIngestAction),
	})
}

func (cmd ModObject) Run(cxt context.Context, args []string) {
	callObjectAction(cxt, cmd, args)
}

