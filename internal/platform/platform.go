package platform

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/spf13/pflag"
)

func FailOut(msg string) {
	fmt.Fprintln(os.Stderr, "ERROR: "+msg)
	os.Exit(1)
}

func ExpandHomePath(path string) string {
	usr, err := user.Current()
	if err != nil {
		FailOut(err.Error())
	}

	if path == "~" {
		return usr.HomeDir
	} else if strings.HasPrefix(path, "~/") {
		return filepath.Join(usr.HomeDir, path[2:])
	}

	return path
}

type BaseArgs struct {
	ConfigPath string
}

func InitBaseFlags(argData *BaseArgs) *pflag.FlagSet {
	flags := pflag.NewFlagSet("base flagset", pflag.PanicOnError)
	flags.StringVar(&argData.ConfigPath, "config", "", "The path to a modfit configuration file")
	return flags
}

