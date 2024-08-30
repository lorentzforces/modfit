package command

type ConfigObject struct{}

func (cmd ConfigObject) Name() string {
	return "config"
}

func (cmd ConfigObject) ShortDescr() string {
	return "persistent settings to customize modfit behavior"
}

func (cmd ConfigObject) actionCmds() map[string]actionCmd {
	return MapNames([]actionCmd{
		new(ConfigGenerateAction),
		new(ConfigResolvePathAction),
	})
}

func (cmd ConfigObject) Run( args []string) {
	callObjectAction(cmd, args)
}
