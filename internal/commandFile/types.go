package commandFile

type CommandConfig struct {
	Command     string `yaml:"command"`
	Description string `yaml:"description"`
}

type CommandsFile struct {
	Commands map[string]CommandConfig `yaml:"commands"`
}