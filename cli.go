package main

import (
	"github.com/alecthomas/kingpin/v2"
)

type Cli struct {
	Config struct {
		Gitignore struct {
			Init *kingpin.CmdClause
			// Add    *kingpin.CmdClause
			// Remove *kingpin.CmdClause
		}

		Init *kingpin.CmdClause
	}
}

func NewCli() *Cli {
	configCommand := kingpin.Command("config", "Manage configuration")

	configGitignoreCommand := configCommand.Command("ignore", "Manage gitignore configuration")

	cli := &Cli{}

	cli.Config.Gitignore.Init = configGitignoreCommand.Command("init", "Initialize gitignore configuration").Action(configGitignoreInitAction)

	cli.Config.Init = configCommand.Command("init", "Initialize configuration").Action(configInitAction)
	configInitArgSshKey = cli.Config.Init.Arg("ssh-key", "SSH key to use").String()

	return cli
}

func (c *Cli) Parse() {
	kingpin.Parse()
}
