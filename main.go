package main

import "time"

func main() {
	g, err := NewGit(".", nil)
	if err != nil {
		panic(err)
	}

	b, err := NewBranch(".", g)
	if err != nil {
		panic(err)
	}

	if err := b.CheckoutToFeature("test"); err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)

	if err := b.CheckoutToProposal("test"); err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)

	if err := b.ReturnToPrevious(); err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)

	if err := b.ReturnToPrevious(); err != nil {
		panic(err)
	}
}
