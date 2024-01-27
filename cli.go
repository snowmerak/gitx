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

	Fork struct {
		Feature  *kingpin.CmdClause
		Proposal *kingpin.CmdClause
		Hotfix   *kingpin.CmdClause
		Bugfix   *kingpin.CmdClause
		Daily    *kingpin.CmdClause
		Revert   *kingpin.CmdClause
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

	forkCommand := kingpin.Command("fork", "Manage branches")

	cli.Fork.Feature = forkCommand.Command("feature", "Fork a feature branch").Action(forkFeatureAction)
	forkCommandArgBranchName = cli.Fork.Feature.Arg("name", "Name of the feature branch").Required().String()

	cli.Fork.Proposal = forkCommand.Command("proposal", "Fork a proposal branch").Action(forkProposalAction)
	forkCommandArgBranchName = cli.Fork.Proposal.Arg("name", "Name of the proposal branch").Required().String()

	cli.Fork.Hotfix = forkCommand.Command("hotfix", "Fork a hotfix branch").Action(forkHotfixAction)
	forkCommandArgBranchName = cli.Fork.Hotfix.Arg("name", "Name of the hotfix branch").Required().String()

	cli.Fork.Bugfix = forkCommand.Command("bugfix", "Fork a bugfix branch").Action(forkBugfixAction)
	forkCommandArgBranchName = cli.Fork.Bugfix.Arg("name", "Name of the bugfix branch").Required().String()

	cli.Fork.Daily = forkCommand.Command("daily", "Fork a daily branch").Action(forkDailyAction)
	forkCommandArgBranchName = cli.Fork.Daily.Arg("name", "Name of the daily branch").Required().String()

	cli.Fork.Revert = forkCommand.Command("revert", "Fork a revert branch").Action(forkRevertAction)

	return cli
}

func (c *Cli) Parse() {
	kingpin.Parse()
}
