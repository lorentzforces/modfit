package command

import (
	"context"
)

type ModDomain struct{}

func (cmd ModDomain) Name() string {
	return "mod"
}

func (cmd ModDomain) ShortDescr() string {
	return "user modification to a game"
}

func (cmd ModDomain) ActionCmds() map[string]Action {
	return MapNames([]Action{
		new(ModIngestAction),
	})
}

func (cmd ModDomain) Run(cxt context.Context, args []string) {
	CallDomainAction(cxt, cmd, args)
}

