package command

import (
	"context"
	"fmt"
	"modfit/internal/platform"
)

type ModDomain struct{}

func (cmd ModDomain) Name() string {
	return "mod"
}

func (cmd ModDomain) ShortDescr() string {
	return "user modification to a game"
}

func (cmd ModDomain) UsageStr() string {
	return "PLACEHOLDER usage for [mod]"
}

func (cmd ModDomain) ActionCmds() map[string]Action {
	return MapNames([]Action{
		new(ModIngestAction),
	})
}

func (cmd ModDomain) Run(cxt context.Context, args []string) {
	CallDomainAction(cxt, cmd, args)
}

type ModIngestAction struct{}

func (cmd ModIngestAction) Name() string {
	return "ingest"
}

func (cmd ModIngestAction) ShortDescr() string {
	return "Add a game mod to modfit's database from existing files"
}

func (cmd ModIngestAction) UsageStr() string {
	return "PLACEHOLDER usage for [mod ingest]"
}

type modIngestArgs struct {
	BaseArgs
}

func (cmd ModIngestAction) Run(ctx context.Context, args []string) {
	parsedArgs := new(modIngestArgs)
	baseFlags := InitBaseFlags(&parsedArgs.BaseArgs)
	baseFlags.Parse(args)

	config := platform.ParseConfig(parsedArgs.ConfigPath)
	parsedArgs.ApplyToConfig(&config)

	fmt.Printf("--DEBUG-- parsed config values: %+v\n", config)

	// TODO: load file config
}
