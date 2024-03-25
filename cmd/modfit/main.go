package main

import (
	"context"
	"fmt"
	"os"

	"github.com/lorentzforces/modfit/internal/platform"
	"github.com/lorentzforces/modfit/internal/subcmd"
)

func main() {
	if len(os.Args) < 2 {
		platform.FailOut("Must specify a valid subcommand")
	}

	subCmd, cmdFound := getSubCmds()[os.Args[1]]

	if !cmdFound {
		platform.FailOut(fmt.Sprintf(
			"Expected a valid subcommand, but was given \"%s\"",
			os.Args[1]),
		)
	}

	subCmd.Run(context.TODO(), os.Args[2:])
}

type command interface {
	Name() string
	Run(cxt context.Context, args []string)
}

func getSubCmds() map[string]command {
	commands := make([]command, 0)
	commands = append(commands, new(subcmd.Ingest))

	cmdMap := make(map[string]command, len(commands))
	for _, command := range commands {
		cmdMap[command.Name()] = command
	}
	return cmdMap
}
