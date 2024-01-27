package main

import (
	"github.com/alecthomas/kingpin/v2"
	"github.com/rs/zerolog/log"
)

var sshGenerateArgName *string

var sshGenerateAction = func(_ *kingpin.ParseContext) error {
	log.Info().Msg("generate ssh key")

	if err := GenerateSSHKey(".", *sshGenerateArgName); err != nil {
		return err
	}

	return nil
}
