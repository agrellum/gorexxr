package main

import (
	. "github.com/agrellum/gorexxr/pkg/lang"
)

func main() {

	// Create a function shortcut, so we do not have to type long names.
	str := RxFromString

	bark := str("woof") // our base stem

	node, _ := bark.GetNode(str("pup"))
	node.Leaf = str("yap")
	node, _ = bark.GetNode(str("bulldog"))
	node.Leaf = str("grrrrr")
	bark.GetNode(str("terrier"))

	node, _ = bark.GetNode(str("pup"))
	pup := node.Leaf
	node, _ = bark.GetNode(str("terrier"))
	terrier := node.Leaf
	node, _ = bark.GetNode(str("bulldog"))
	bulldog := node.Leaf

	SayRexx(str("Sounds: ").OpCcBlank(nil, pup).OpCcBlank(nil, terrier).OpCcBlank(nil, bulldog))

}
