package command

import (
	"bytes"
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
}

type ObjectCmd interface {
	cmdData
	ActionCmds() map[string]ActionCmd
}

type ActionCmd interface {
	cmdData
	UsageStr() string
}

func MapNames[C cmdData](cmds []C) map[string]C {
	cmdMap := make(map[string]C, len(cmds))
	for _, command := range cmds {
		cmdMap[command.Name()] = command
	}
	return cmdMap
}

func CallObjectAction[O ObjectCmd](ctxt context.Context, object O, args []string) {
	if len(args) < 1 {
		platform.FailOut("Must specify an action")
	}

	if args[0] == "help" {
		if len(args) == 1 {
			fmt.Fprintln(os.Stderr, FmtHelp(object, ""))
			os.Exit(1)
		}
		cmd, cmdFound := object.ActionCmds()[args[1]]
		if !cmdFound {
			platform.FailOut(fmt.Sprintf(
				"Expected a valid action, but was given \"%s\"",
				args[1],
			))
		}
		fmt.Fprintln(os.Stderr, cmd.UsageStr())
		os.Exit(1)
	}

	action, cmdFound := object.ActionCmds()[args[0]]
	if !cmdFound {
		platform.FailOut(fmt.Sprintf(
			"Expected a valid action, but was given \"%s\"",
			args[0],
		))
	}

	action.Run(context.TODO(), args[1:])
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

func FmtHelp(o ObjectCmd, actionName string) string {
	if len(actionName) == 0 {
		return fmtObjectHelp(o)
	}

	action, cmdFound := o.ActionCmds()[actionName]
	if !cmdFound {
		platform.FailOut(fmt.Sprintf(
			"Expected a valid action, but was given \"%s\"",
			actionName,
		))
	}

	return action.UsageStr()
}

func fmtObjectHelp(o ObjectCmd) string {
	var outputBuf bytes.Buffer

	fmt.Fprintf(&outputBuf, "OBJECT: %s\n", o.Name())
	fmt.Fprintf(&outputBuf, "%s\n\n", o.ShortDescr())
	fmt.Fprintf(&outputBuf, "Usage: modfit %s [COMMAND]\n\n", o.Name())
	fmt.Fprintf(&outputBuf, "COMMANDS:\n")

	cmds := o.ActionCmds()

	nameSize := 0
	for _, cmd := range cmds {
		curSize := len([]rune(cmd.Name()))
		if curSize > nameSize {
			nameSize = curSize
		}
	}
	nameSize++

	fmtString := fmt.Sprintf("  %%-%ds %%s\n", nameSize)
	for _, cmd := range cmds {
		fmt.Fprintf(&outputBuf, fmtString, cmd.Name(), cmd.ShortDescr())
	}
	fmt.Fprintf(&outputBuf, "\nSee help about any command by running \"modfit %s help [command]\".", o.Name())

	return outputBuf.String()
}

