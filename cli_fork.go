package main

import "github.com/alecthomas/kingpin/v2"

var ForkAction = func(path string, name string, kind BranchType) error {
	key, _ := GetAuthKeyFromConfig(path)

	g, err := NewGit(path, key)
	if err != nil {
		return err
	}

	b, err := NewBranch(path, g)
	if err != nil {
		return err
	}

	switch kind {
	case BranchFeature:
		return b.CheckoutToFeature(name)
	case BranchProposal:
		return b.CheckoutToProposal(name)
	case BranchHotfix:
		return b.CheckoutToHotfix(name)
	case BranchBugfix:
		return b.CheckoutToBugfix(name)
	case BranchDaily:
		return b.CheckoutToDaily()
	}

	return nil
}

var forkCommandArgBranchName *string

var forkFeatureAction = func(_ *kingpin.ParseContext) error {
	path := "."

	if err := ForkAction(path, *forkCommandArgBranchName, BranchFeature); err != nil {
		return err
	}

	return nil
}

var forkProposalAction = func(_ *kingpin.ParseContext) error {
	path := "."

	if err := ForkAction(path, *forkCommandArgBranchName, BranchProposal); err != nil {
		return err
	}

	return nil
}

var forkHotfixAction = func(_ *kingpin.ParseContext) error {
	path := "."

	if err := ForkAction(path, *forkCommandArgBranchName, BranchHotfix); err != nil {
		return err
	}

	return nil
}

var forkBugfixAction = func(_ *kingpin.ParseContext) error {
	path := "."

	if err := ForkAction(path, *forkCommandArgBranchName, BranchBugfix); err != nil {
		return err
	}

	return nil
}

var forkDailyAction = func(_ *kingpin.ParseContext) error {
	path := "."

	if err := ForkAction(path, *forkCommandArgBranchName, BranchDaily); err != nil {
		return err
	}

	return nil
}

var forkRevertAction = func(_ *kingpin.ParseContext) error {
	path := "."

	key, _ := GetAuthKeyFromConfig(path)

	g, err := NewGit(path, key)
	if err != nil {
		return err
	}

	b, err := NewBranch(path, g)
	if err != nil {
		return err
	}

	if err := b.ReturnToPrevious(); err != nil {
		return err
	}

	return nil
}
