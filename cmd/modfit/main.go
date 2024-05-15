package main

import (
	"bytes"
	"context"
	"fmt"
	"modfit/internal/command"
	"modfit/internal/platform"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		platform.FailOut("Must specify a valid subcommand")
	}

	switch os.Args[1] {
		case "-h", "--help":
			printTopLevelHelp()
		case "help":
			if len(os.Args) == 2 {
				printTopLevelHelp()
				os.Exit(1)
			}

			subCmd, cmdFound := domainCmds()[os.Args[2]]
			if !cmdFound {
				failMsg :=
					"Expected a valid modfit object, but was given \"%s\"\n" +
					"Run \"modfit --help\" to see valid objects and usage information."
				platform.FailOut(fmt.Sprintf(
					failMsg,
					os.Args[2],
				))
			}

			fmt.Fprintln(os.Stderr, subCmd.UsageStr())
			os.Exit(1)
		default: {
			subCmd, cmdFound := domainCmds()[os.Args[1]]

			if !cmdFound {
				failMsg :=
					"Expected a valid modfit object, but was given \"%s\"\n" +
					"Run \"modfit --help\" to see valid objects and usage information."
				platform.FailOut(fmt.Sprintf(
					failMsg,
					os.Args[1],
				))
			}

			subCmd.Run(context.TODO(), os.Args[2:])
		}
	}
}

func domainCmds() map[string]command.Domain {
	return command.MapNames([]command.Domain{
		new(command.ModDomain),
		new(command.ConfigDomain),
	})
}

// TODO: clarify domain/object/command terminology
func printTopLevelHelp() {
	var outputBuf bytes.Buffer

	topHelpHeader :=
		"Modfit is a command-line mod manager.\n" +
		"Commands are collected under the domain objects they operate on, so a command takes " +
		"the form: \n" +
		"      modfit [OBJECT] [COMMAND]\n" +
		"\n" +
		"You can see this message by running \"modfit help\".\n" +
		"\n" +
		"OBJECTS IN MODFIT:\n"
	fmt.Fprint(&outputBuf, topHelpHeader)

	domains := domainCmds()

	nameSize := 0
	for _, domain := range domains {
		nameSize = len([]rune(domain.Name()))
	}
	nameSize++
	fmtString := fmt.Sprintf("  %%-%ds %%s\n", nameSize)

	for _, domain := range domains {
		fmt.Fprintf(&outputBuf, fmtString, domain.Name(), domain.ShortDescr())
	}

	fmt.Fprint(&outputBuf, "\nSee help about any object by running \"modfit help [object]\".")
	fmt.Fprint(&outputBuf, "\n\nCommon options:\n")

	fmt.Fprintln(os.Stderr, &outputBuf)
	flags := command.InitBaseFlags(new(command.BaseArgs))
	flags.PrintDefaults()
	os.Exit(1)
}
