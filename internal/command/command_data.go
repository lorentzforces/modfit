package command

import (
	"context"
	"fmt"
	"modfit/internal/platform"
	"os"

	"github.com/spf13/pflag"
)

// TODO: quick audit exported members, a lot can probably be unexported

type cmdData interface {
	Name() string
	ShortDescr() string
	Run(cxt context.Context, args []string)
	UsageStr() string
}

type Domain interface {
	cmdData
	ActionCmds() map[string]Action
}

type Action interface {
	cmdData
}

func MapNames[C cmdData](cmds []C) map[string]C {
	cmdMap := make(map[string]C, len(cmds))
	for _, command := range cmds {
		cmdMap[command.Name()] = command
	}
	return cmdMap
}

func CallDomainAction[D Domain](ctxt context.Context, domain D, args []string) {
	if len(args) < 1 {
		platform.FailOut("Must specify an action")
	}

	if args[0] == "help" {
		if len(args) == 1 {
			fmt.Fprintln(os.Stderr, domain.UsageStr())
			os.Exit(1)
		}
		cmd, cmdFound := domain.ActionCmds()[args[1]]
		if !cmdFound {
			platform.FailOut(fmt.Sprintf(
				"Expected a valid action, but was given \"%s\"",
				args[1],
			))
		}
		fmt.Fprintln(os.Stderr, cmd.UsageStr())
		os.Exit(1)
	}

	action, cmdFound := domain.ActionCmds()[args[0]]
	if !cmdFound {
		platform.FailOut(fmt.Sprintf(
			"Expected a valid action, but was given \"%s\"",
			args[0],
		))
	}

	action.Run(context.TODO(), args[1:])
}

func PrintDomainHelp[D Domain](ctxt context.Context, domain D, cmdName string) {
	cmd := domain.ActionCmds()[cmdName]
	fmt.Fprintln(os.Stderr, cmd.UsageStr())
}

type BaseArgs struct {
	ConfigPath string
	LogPath string
	Verbose bool
}

func InitBaseFlags(argData *BaseArgs) *pflag.FlagSet {
	flags := pflag.NewFlagSet("base flagset", pflag.PanicOnError)
	flags.StringVar(&argData.ConfigPath, "config", "", "The path to a modfit configuration file")
	flags.BoolVarP(&argData.Verbose, "verbose", "v", false, "Print progress information to the console")
	return flags
}

func (args *BaseArgs) ApplyToConfig(cfg *platform.Config) {
	if len(args.LogPath) > 0 {
		cfg.LogPath = args.LogPath
	}
	if args.Verbose {
		cfg.Verbosity = platform.Verbose
	}
}

