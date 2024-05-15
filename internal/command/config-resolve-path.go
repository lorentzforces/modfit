package command

import (
	"context"
	"fmt"
	"modfit/internal/platform"
	"os"
	"strings"
)

type ConfigResolvePathAction struct{}

func (cmd ConfigResolvePathAction) Name() string {
	return "resolve-path"
}

func (cmd ConfigResolvePathAction) ShortDescr() string {
	return "Get the path which modfit will attempt to resolve a config file at"
}

func (cmd ConfigResolvePathAction) UsageStr() string {
	return "PLACEHOLDER usage for [config resolve-path]"
}

type ConfigResolvePathArgs struct {
	BaseArgs
}

func (cmd ConfigResolvePathAction) Run(cts context.Context, args []string) {
	parsedArgs := new(ConfigResolvePathArgs)
	baseFlags := InitBaseFlags(&parsedArgs.BaseArgs)
	baseFlags.Parse(args)

	configPath, foundFile, err := platform.GetConfigPath(parsedArgs.ConfigPath)
	if foundFile && err != nil {
		platform.FailOut(err.Error())
	}

	foundString := "not_found"
	if foundFile {
		foundString = "found"
	}

	fmt.Printf("\"%s\" %s\n", configPath, foundString)

	if !strings.HasSuffix(configPath, ".toml") {
		fmt.Fprintln(
			os.Stderr,
			"Warning: resolved path does not end in .toml. This is not an error, but may " +
				"indicate a misconfigured configuration path.",
		)
	}

	if !foundFile {
		os.Exit(1)
	}
}
