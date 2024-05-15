package command

import (
	"context"
	"fmt"
	"modfit/internal/platform"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

type ConfigDomain struct{}

func (cmd ConfigDomain) Name() string {
	return "config"
}

func (cmd ConfigDomain) ShortDescr() string {
	return "persistent settings to customize modfit behavior"
}

func (cmd ConfigDomain) ActionCmds() map[string]Action {
	return MapNames([]Action{
		new(ConfigGenerateAction),
		new(ConfigResolvePathAction),
	})
}

func (cmd ConfigDomain) Run(cxt context.Context, args []string) {
	CallDomainAction(cxt, cmd, args)
}

func (cmd ConfigDomain) UsageStr() string {
	return "PLACEHOLDER usage for [config]"
}

type ConfigGenerateAction struct{}

func (cmd ConfigGenerateAction) Name() string {
	return "generate"
}

func (cmd ConfigGenerateAction) ShortDescr() string {
	return "Generate a modfit config file"
}

func (cmd ConfigGenerateAction) UsageStr() string {
	return "PLACEHOLDER usage for [config generate]"
}

type ConfigGenerateArgs struct {
	BaseArgs
}

func (cmd ConfigGenerateAction) Run(ctx context.Context, args []string) {
	parsedArgs := new(ConfigGenerateArgs)
	baseFlags := InitBaseFlags(&parsedArgs.BaseArgs)
	baseFlags.Parse(args)

	// Error usually means that we didn't find the file. If something will prevent us writing the
	// file, we'll encounter that later anyway.
	configPath, foundFile, _  := platform.GetConfigPath(parsedArgs.ConfigPath)
	if foundFile {
		platform.FailOut(fmt.Sprintf("Resolved config file path already exists: %s", configPath))
	}

	configFile, err := os.Create(configPath)
	defer configFile.Close()
	if err != nil {
		platform.FailOut(fmt.Sprintf(
			"Could not open file for writing: %s\n%s",
			configPath,
			err.Error(),
		))
	}

	_, err = configFile.WriteString(configHeader)
	if err != nil {
		platform.FailOut(err.Error())
	}

	tomlEncoder := toml.NewEncoder(configFile)
	err = tomlEncoder.Encode(platform.DefaultConfig())
	if err != nil {
		platform.FailOut(fmt.Sprintf(
			"Failed to write default config to %s, but the file was created successfully. You " +
				"may wish to delete this file.\n%s",
			configPath,
			err.Error(),
		))
	}
}

const configHeader string = "# this file was generated by modfit\n"

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
	os.Exit(0)
}
