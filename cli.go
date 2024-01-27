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

	SSH struct {
		Generate *kingpin.CmdClause
	}
}

func NewCli() *Cli {
	cli := &Cli{}

	configCommand := kingpin.Command("config", "Manage configuration")

	configGitignoreCommand := configCommand.Command("ignore", "Manage gitignore configuration")

	cli.Config.Gitignore.Init = configGitignoreCommand.Command("init", "Initialize gitignore configuration").Action(configGitignoreInitAction)

	cli.Config.Init = configCommand.Command("init", "Initialize configuration").Action(configInitAction)
	configInitArgSshKey = cli.Config.Init.Arg("ssh-key", "SSH key to use").String()

	sshCommand := kingpin.Command("ssh", "Manage SSH keys")

	cli.SSH.Generate = sshCommand.Command("generate", "Generate SSH key").Action(sshGenerateAction)
	sshGenerateArgName = cli.SSH.Generate.Arg("name", "Name of the SSH key").Required().String()

	return cli
}

func (c *Cli) Parse() {
	kingpin.Parse()
}
