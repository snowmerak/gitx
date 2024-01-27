package main

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

type Git struct {
	AuthKey    *ssh.PublicKeys
	Path       string
	Repository *git.Repository
}

func Clone(url string, authKey *ssh.PublicKeys) (*Git, error) {
	option := &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	}

	if authKey != nil {
		option.Auth = authKey
	}

	name := strings.TrimSuffix(filepath.Base(url), filepath.Ext(url))
	r, err := git.PlainClone(filepath.Join(".", name), false, option)

	if err != nil {
		return nil, fmt.Errorf("failed to clone %s to %s: %w", url, name, err)
	}

	return &Git{
		AuthKey:    authKey,
		Path:       filepath.Join(".", name),
		Repository: r,
	}, nil
}

func NewGit(path string, authKey *ssh.PublicKeys) (*Git, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open repository: %w", err)
	}

	return &Git{
		AuthKey:    authKey,
		Path:       path,
		Repository: r,
	}, nil
}

func (g *Git) Pull() error {
	r := g.Repository

	w, err := r.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get worktree: %w", err)
	}

	option := &git.PullOptions{
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	}

	if g.AuthKey != nil {
		option.Auth = g.AuthKey
	}

	err = w.Pull(option)
	if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		return fmt.Errorf("failed to pull: %w", err)
	}

	return nil
}

func (g *Git) Push() error {
	r := g.Repository

	option := &git.PushOptions{}

	if g.AuthKey != nil {
		option.Auth = g.AuthKey
	}

	if err := r.Push(option); err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		return fmt.Errorf("failed to push: %w", err)
	}

	return nil
}

func (g *Git) Remotes() ([]string, error) {
	r := g.Repository

	remotes, err := r.Remotes()
	if err != nil {
		return nil, fmt.Errorf("failed to get remotes: %w", err)
	}

	var urls []string
	for _, remote := range remotes {
		urls = append(urls, remote.Config().URLs...)
	}

	return urls, nil
}

func (g *Git) CreateBranch(branch string) error {
	r := g.Repository

	w, err := r.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get worktree: %w", err)
	}

	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(branch),
		Create: true,
	})
	if err != nil {
		return fmt.Errorf("failed to checkout: %w", err)
	}

	return nil
}

func (g *Git) GetCurrentBranch() (string, error) {
	r := g.Repository

	ref, err := r.Head()
	if err != nil {
		return "", fmt.Errorf("failed to get head: %w", err)
	}

	return ref.Name().Short(), nil
}

func (g *Git) CheckoutBranch(branch string) error {
	r := g.Repository

	w, err := r.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get worktree: %w", err)
	}

	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(branch),
	})
	if err != nil {
		return fmt.Errorf("failed to checkout: %w", err)
	}

	return nil
}

func (g *Git) DeleteBranch(branch string) error {
	r := g.Repository

	ref := plumbing.NewBranchReferenceName(branch)

	if err := r.Storer.RemoveReference(ref); err != nil {
		return fmt.Errorf("failed to remove branch: %w", err)
	}

	return nil
}

func (g *Git) Add(files []string) error {
	r := g.Repository

	w, err := r.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get worktree: %w", err)
	}

	for _, file := range files {
		_, err = w.Add(file)
		if err != nil {
			return fmt.Errorf("failed to add %s: %w", file, err)
		}
	}

	return nil
}

func (g *Git) Commit(message string) error {
	r := g.Repository

	w, err := r.Worktree()
	if err != nil {
		return fmt.Errorf("failed to get worktree: %w", err)
	}

	_, err = w.Commit(message, &git.CommitOptions{
		All: true,
	})
	if err != nil {
		return fmt.Errorf("failed to commit: %w", err)
	}

	return nil
}

func (g *Git) Status() (git.Status, error) {
	r := g.Repository

	w, err := r.Worktree()
	if err != nil {
		return nil, fmt.Errorf("failed to get worktree: %w", err)
	}

	status, err := w.Status()
	if err != nil {
		return nil, fmt.Errorf("failed to get status: %w", err)
	}

	return status, nil
}
