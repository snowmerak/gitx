package main

import (
	"github.com/alecthomas/kingpin/v2"
	"github.com/rs/zerolog/log"
)

var configInitArgSshKey *string

var configInitAction = func(_ *kingpin.ParseContext) error {
	log.Info().Msg("init configuration")

	cfg := Config{}

	if configInitArgSshKey != nil && *configInitArgSshKey != "" {
		cfg.SSHKeyFile = *configInitArgSshKey
	}

	if err := WriteGitxConfig(".", cfg); err != nil {
		return err
	}

	return nil
}
