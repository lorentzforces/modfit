package subcmd

import (
	"context"
	"fmt"

	"github.com/lorentzforces/modfit/internal/platform"
)

type Ingest struct{}

func (cmd *Ingest) Name() string {
	return "ingest"
}

func (cmd *Ingest) Run(ctx context.Context, args []string) {
	parsedArgs := new(ingestArgs)
	baseFlags := platform.InitBaseFlags(&parsedArgs.BaseArgs)
	baseFlags.Parse(args)

	parsedArgs.ConfigPath = platform.ExpandHomePath(parsedArgs.ConfigPath)
	fmt.Printf("--DEBUG-- config path: \"%s\"\n", parsedArgs.ConfigPath)


	// TODO: load config file
}

type ingestArgs struct {
	platform.BaseArgs
}
