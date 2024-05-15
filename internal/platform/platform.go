package platform

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/BurntSushi/toml"
)

const configEnvVar = "MODFIT_CONFIG"
const defaultConfigPath =  "~/.config/modfit/config.toml"

func FailOut(msg string) {
	fmt.Fprintln(os.Stderr, "ERROR: "+msg)
	os.Exit(1)
}

func ErrMsg(msg string) string {
	return "ERROR: "+msg
}

type Config struct {
	LogPath string
	Verbosity VerbosityLevel
}

// TODO: figure out how to make this something we can actually get as input from user
type VerbosityLevel string
const (
	NotVerbose VerbosityLevel = "NOT_VERBOSE"
	Verbose = "VERBOSE"
)

func DefaultConfig() Config {
	return Config{
		LogPath: replaceTilde("~/.local/state/modfit/modfit.log"),
		Verbosity: NotVerbose,
	}
}

func ParseConfig(configPath string) Config {
	path, foundFile, err := GetConfigPath(configPath)
	if err != nil {
		panic(ErrMsg(err.Error()))
	}

	configData := DefaultConfig()

	if !foundFile {
		return configData
	}

	toml.DecodeFile(path, &configData)

	return configData
}

// Resolves the config from the specified path, if any. If the passed path is empty, will use:
// - value of the MODFIT_CONFIG environment variable
// - the default config path on the platform
// For now, modfit only supports the default XDG_CONFIG directory: ~/.config/modfit/config.toml
func GetConfigPath(configPath string) (path string, foundFile bool, err error) {
	fileSpecified := true
	envPath := os.Getenv(configEnvVar)

	if len(configPath) == 0 && len(envPath) > 0 {
		configPath = envPath
	}

	if len(configPath) == 0 {
		fileSpecified = false
		configPath = defaultConfigPath
	}

	configFile, err := os.Open(configPath)
	defer configFile.Close()

	// no config file exists, so just use the default
	if err != nil && !fileSpecified {
		// TODO: perhaps log that default configuration was used?
		return configPath, false, nil
	}
	// config file was specified, but we couldn't read it
	if err != nil {
		return configPath, false, err
	}
	 // in theory if above succeeded this can't fail
	configStat, err := configFile.Stat()
	Assert(err == nil, err)

	if configStat.IsDir() {
		return configPath,
			true,
			errors.New(fmt.Sprintf("Resolved config file is a directory: \"%s\"", configPath))
	}

	return configPath, true, nil
}

func replaceTilde(s string) string {
	if strings.HasPrefix(s, "~") {
		usr, err := user.Current()
		if err != nil {
			FailOut(err.Error())
		}
		s = strings.Replace(s, "~", usr.HomeDir, 1)
	}
	return s
}

func Assert(condition bool, more any) {
	if condition { return }
	panic(fmt.Sprintf("Assertion violated: %s", more))
}
