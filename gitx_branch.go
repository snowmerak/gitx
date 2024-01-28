package main

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

type BranchType uint8

const (
	BranchRelease BranchType = iota
	BranchDevelop
	BranchFeature
	BranchHotfix
	BranchBugfix
	BranchProposal
	BranchDaily
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
	git *Git
}

func NewBranch(path string, git *Git) (*Branch, error) {
	return &Branch{
		git: git,
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

	return nil
}

func (b *Branch) CheckoutToDaily() error {
	if err := b.checkBranchIsClean(); err != nil {
		return err
	}

	r := make([]byte, 16)
	if _, err := rand.Read(r); err != nil {
		return err
	}

	today := time.Now().Format("20060102")

	bn := "daily/" + today + "-" + hex.EncodeToString(r)

	if err := b.git.CreateBranch(bn); err != nil {
		return err
	}

	return nil
}

type BranchStackIsHeadError struct{}

func (e *BranchStackIsHeadError) Error() string {
	return "branch stack is head"
}
