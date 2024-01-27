package main

import (
	"github.com/alecthomas/kingpin/v2"
	"github.com/rs/zerolog/log"
)

var commitPullAction = func(_ *kingpin.ParseContext) error {
	path := "."

	key, _ := GetAuthKeyFromConfig(path)

	g, err := NewGit(path, key)
	if err != nil {
		return err
	}

	if err := g.Pull(); err != nil {
		return err
	}

	return nil
}

var commitPushArgMessage *string

var commitPushAction = func(_ *kingpin.ParseContext) error {
	path := "."

	key, _ := GetAuthKeyFromConfig(path)

	g, err := NewGit(path, key)
	if err != nil {
		return err
	}

	if err := AddAllWorkedFiles(g); err != nil {
		return err
	}

	if err := g.Commit(*commitPushArgMessage); err != nil {
		return err
	}

	if err := g.Push(); err != nil {
		return err
	}

	return nil
}

var commitChangesAction = func(_ *kingpin.ParseContext) error {
	path := "."

	key, _ := GetAuthKeyFromConfig(path)

	g, err := NewGit(path, key)
	if err != nil {
		return err
	}

	status, err := g.Status()
	if err != nil {
		return err
	}

	changes := FormatChangesList(status)

	log.Info().Msg("changes\n" + changes)

	return nil
}
