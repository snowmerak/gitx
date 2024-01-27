package main

import (
	"github.com/alecthomas/kingpin/v2"
	"github.com/rs/zerolog/log"
)

var configGitignoreInitAction = func(_ *kingpin.ParseContext) error {
	log.Info().Msg("init gitignore configuration")
	if err := AppendGitignoreFields(".", defaultGitIgnoreFields...); err != nil {
		return err
	}
	return nil
}
