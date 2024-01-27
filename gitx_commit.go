package main

import (
	"slices"
	"strings"

	"github.com/go-git/go-git/v5"
)

type Status uint8

const (
	StatusCreated Status = iota
	StatusChanged
	StatusDeleted
	StatusUnchanged
)

func FormatChangesList(status git.Status) string {
	changes := make(map[Status][]string)

	for k, v := range status {
		switch v.Worktree {
		case git.Modified:
			changes[StatusChanged] = append(changes[StatusChanged], k)
		case git.Unmodified:
			changes[StatusUnchanged] = append(changes[StatusUnchanged], k)
		case git.Deleted:
			changes[StatusDeleted] = append(changes[StatusDeleted], k)
		case git.Renamed:
			changes[StatusChanged] = append(changes[StatusChanged], k)
		case git.Copied:
			changes[StatusChanged] = append(changes[StatusChanged], k)
		case git.Untracked:
			changes[StatusCreated] = append(changes[StatusCreated], k)
		case git.UpdatedButUnmerged:
			changes[StatusChanged] = append(changes[StatusChanged], k)
		}
	}

	keys := make([]Status, 0, len(changes))
	for k := range changes {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	sb := strings.Builder{}

	for _, k := range keys {
		v := changes[k]

		if len(v) == 0 {
			continue
		}

		switch k {
		case StatusChanged:
			sb.WriteString("Changed files:\n")
		case StatusUnchanged:
			continue
		case StatusDeleted:
			sb.WriteString("Deleted files:\n")
		case StatusCreated:
			sb.WriteString("Created files:\n")
		default:
			sb.WriteString("Unknown:\n")
		}

		for _, s := range v {
			sb.WriteString("  ")
			sb.WriteString(s)
			sb.WriteString("\n")
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func AddAllWorkedFiles(g *Git) error {
	w, err := g.Repository.Worktree()
	if err != nil {
		return err
	}

	_, err = w.Add(".")
	if err != nil {
		return err
	}

	return nil
}
