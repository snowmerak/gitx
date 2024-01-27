package main

import (
	"path/filepath"
)

type BranchType uint8

const (
	BranchRelease BranchType = iota
	BranchDevelop
	BranchFeature
	BranchHotfix
	BranchBugfix
	BranchProposal
)

func (b BranchType) String() string {
	switch b {
	case BranchRelease:
		return "release"
	case BranchDevelop:
		return "develop"
	case BranchFeature:
		return "feature"
	case BranchHotfix:
		return "hotfix"
	case BranchBugfix:
		return "bugfix"
	case BranchProposal:
		return "proposal"
	default:
		return "unknown"
	}
}

func (b BranchType) Short() string {
	switch b {
	case BranchRelease:
		return "r"
	case BranchDevelop:
		return "d"
	case BranchFeature:
		return "f"
	case BranchHotfix:
		return "h"
	case BranchBugfix:
		return "b"
	case BranchProposal:
		return "p"
	default:
		return "u"
	}
}

type Branch struct {
	git   *Git
	stack *BranchStack
}

func NewBranch(path string, git *Git) (*Branch, error) {
	bd := filepath.Join(path, gitxFolder, gitxBranchDir)
	stack, err := NewBranchStack(bd)
	if err != nil {
		return nil, err
	}

	if stack.Len() == 0 {
		currentBranchName, err := git.GetCurrentBranch()
		if err != nil {
			return nil, err
		}

		if err := stack.Push(currentBranchName); err != nil {
			return nil, err
		}
	}

	return &Branch{
		stack: stack,
		git:   git,
	}, nil
}

type BranchIsNotCleanError struct{}

func (e *BranchIsNotCleanError) Error() string {
	return "branch is not clean"
}

func (b *Branch) checkBranchIsClean() error {
	st, err := b.git.Status()
	if err != nil {
		return err
	}

	if !st.IsClean() {
		return &BranchIsNotCleanError{}
	}

	return nil
}

func (b *Branch) CheckoutToFeature(name string) error {
	if err := b.checkBranchIsClean(); err != nil {
		return err
	}

	bn := "feature/" + name

	if err := b.git.CreateBranch(bn); err != nil {
		return err
	}

	if err := b.stack.Push(bn); err != nil {
		return err
	}

	return nil
}

func (b *Branch) CheckoutToProposal(name string) error {
	if err := b.checkBranchIsClean(); err != nil {
		return err
	}

	bn := "proposal/" + name

	if err := b.git.CreateBranch(bn); err != nil {
		return err
	}

	if err := b.stack.Push(bn); err != nil {
		return err
	}

	return nil
}

func (b *Branch) CheckoutToHotfix(name string) error {
	if err := b.checkBranchIsClean(); err != nil {
		return err
	}

	bn := "hotfix/" + name

	if err := b.git.CreateBranch(bn); err != nil {
		return err
	}

	if err := b.stack.Push(bn); err != nil {
		return err
	}

	return nil
}

func (b *Branch) CheckoutToBugfix(name string) error {
	if err := b.checkBranchIsClean(); err != nil {
		return err
	}

	bn := "bugfix/" + name

	if err := b.git.CreateBranch(bn); err != nil {
		return err
	}

	if err := b.stack.Push(bn); err != nil {
		return err
	}

	return nil
}

type BranchStackIsHeadError struct{}

func (e *BranchStackIsHeadError) Error() string {
	return "branch stack is head"
}

func (b *Branch) ReturnToPrevious() error {
	current, ok := b.stack.Top()
	if !ok {
		return &BranchStackIsHeadError{}
	}

	if _, err := b.stack.Pop(); err != nil {
		return err
	}

	previous, ok := b.stack.Top()
	if !ok {
		return &BranchStackIsHeadError{}
	}

	if err := b.git.CheckoutBranch(previous); err != nil {
		return err
	}

	if err := b.git.DeleteBranch(current); err != nil {
		return err
	}

	return nil
}
