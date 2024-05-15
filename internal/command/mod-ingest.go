package command

import (
	"context"
	"fmt"
	"modfit/internal/platform"
)

type ModIngestAction struct{}

func (cmd ModIngestAction) Name() string {
	return "ingest"
}

func (cmd ModIngestAction) ShortDescr() string {
	return "Add a game mod to modfit's database from existing files"
}

func (cmd ModIngestAction) usageStr() string {
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
	parsedArgs.applyToConfig(&config)

	fmt.Printf("--DEBUG-- parsed config values: %+v\n", config)

	// TODO: load file config
}
