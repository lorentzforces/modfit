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
		failMsg :=
			"Must specify a valid object.\n" +
			"Run \"modfit help\" to see valid objects and usage information."
		platform.FailOut(failMsg)
	}

	switch os.Args[1] {
		case "-h", "--help":
			printTopLevelHelp()
		case "help":
			if len(os.Args) == 2 {
				fmt.Fprintln(os.Stderr, printTopLevelHelp())
				os.Exit(1)
			}

			objectCmd, cmdFound := objectCmds()[os.Args[2]]
			if !cmdFound {
				failMsg :=
					"Expected a valid modfit object, but was given \"%s\"\n" +
					"Run \"modfit help\" to see valid objects and usage information."
				platform.FailOut(fmt.Sprintf(
					failMsg,
					os.Args[2],
				))
			}

			fmt.Fprintln(os.Stderr, command.FmtHelp(objectCmd, ""))
			os.Exit(1)
		default: {
			objectCmd, cmdFound := objectCmds()[os.Args[1]]

			if !cmdFound {
				failMsg :=
					"Expected a valid modfit object, but was given \"%s\"\n" +
					"Run \"modfit help\" to see valid objects and usage information."
				platform.FailOut(fmt.Sprintf(
					failMsg,
					os.Args[1],
				))
			}

			objectCmd.Run(context.TODO(), os.Args[2:])
		}
	}
}

func objectCmds() map[string]command.ObjectCmd {
	return command.MapNames([]command.ObjectCmd{
		new(command.ModObject),
		new(command.ConfigObject),
	})
}

func printTopLevelHelp() string {
	var outputBuf bytes.Buffer

	topHelpHeader :=
		"Modfit is a command-line mod manager.\n" +
		"Actions are collected under the domain objects they operate on, so a command takes " +
		"the form: \n" +
		"      modfit [OBJECT] [ACTION]\n" +
		"\n" +
		"You can see this message by running \"modfit help\".\n" +
		"\n" +
		"OBJECTS IN MODFIT:\n"
	fmt.Fprint(&outputBuf, topHelpHeader)

	objects := objectCmds()

	nameSize := 0
	for _, object := range objects {
		curSize := len([]rune(object.Name()))
		if curSize > nameSize {
			nameSize = curSize
		}
	}
	nameSize++
	fmtString := fmt.Sprintf("  %%-%ds %%s\n", nameSize)

	for _, object := range objects {
		fmt.Fprintf(&outputBuf, fmtString, object.Name(), object.ShortDescr())
	}

	fmt.Fprint(&outputBuf, "\nSee help about any object by running \"modfit help [object]\".")
	fmt.Fprint(&outputBuf, "\n\nCOMMON OPTIONS:\n")
	flags := command.InitBaseFlags(new(command.BaseArgs))
	fmt.Fprint(&outputBuf, flags.FlagUsages())

	return outputBuf.String()
}
