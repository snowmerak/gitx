package main

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
}
